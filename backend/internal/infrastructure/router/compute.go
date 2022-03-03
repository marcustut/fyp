package router

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/gorilla/mux"
	"github.com/marcustut/fyp/backend/ent"
	"github.com/marcustut/fyp/backend/ent/schema/ulid"
	"github.com/marcustut/fyp/backend/internal/adapter/controller"
	awsConstant "github.com/marcustut/fyp/backend/internal/const/aws"
	"github.com/marcustut/fyp/backend/internal/entity/model"
)

// RunRequestInput is the input for RunRequest
type RunRequestInput struct {
	SlideID string `json:"slide_id"`
}

// TerminateRequestInput is the input for TerminateRequest
type TerminateRequestInput struct {
	SlideID string `json:"slide_id"`
}

// NewCompute ...
func NewCompute(ctrl controller.Controller, cfg aws.Config) *mux.Router {
	r := mux.NewRouter()

	client := ec2.NewFromConfig(cfg)

	r.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		// parse request body
		var rri RunRequestInput
		err := json.NewDecoder(r.Body).Decode(&rri)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", model.NewInvalidRequestBodyError(err))) != nil {
				panic("error encoding json")
			}
			return
		}

		// fetch instance (return if already exists)
		ic, err := ctrl.Instance.List(r.Context(), nil, nil, nil, nil, &ent.InstanceWhereInput{
			HasSlideWith: []*ent.SlideWhereInput{
				{ID: (*ulid.ID)(&rri.SlideID)},
			},
		}, nil)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", err)) != nil {
				panic("error encoding json")
			}
			return
		}
		if ic != nil {
			for _, edge := range ic.Edges {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				if json.NewEncoder(w).Encode(edge.Node) != nil {
					panic("error encoding json")
				} else {
					return
				}
			}
		}

		// fetch slide
		s, err := ctrl.Slide.Get(r.Context(), ulid.ID(rri.SlideID))
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", err)) != nil {
				panic("error encoding json")
			}
			return
		}
		// fetch owner of the slide
		userID, err := s.QueryUser().OnlyID(r.Context())
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", err)) != nil {
				panic("error encoding json")
			}
			return
		}

		// update userData for ec2 startup
		file, err := ioutil.ReadFile("scripts/ec2-startup.sh")
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", err)) != nil {
				panic("error encoding json")
			}
			return
		}
		userData := strings.ReplaceAll(string(file), "${BUCKET_NAME}", awsConstant.BucketName)
		userData = strings.ReplaceAll(userData, "${USER_ID}", string(userID))
		userData = strings.ReplaceAll(userData, "${SLIDE_ID}", string(s.ID))
		b64UserData := base64.StdEncoding.EncodeToString([]byte(userData))

		// start creating ec2 instance
		output, err := client.RunInstances(r.Context(), &ec2.RunInstancesInput{
			MaxCount:     aws.Int32(1),
			MinCount:     aws.Int32(1),
			ImageId:      aws.String(awsConstant.AmiAmazonLinux2x86_64),
			InstanceType: types.InstanceTypeT2Micro,
			IamInstanceProfile: &types.IamInstanceProfileSpecification{
				Arn: aws.String(awsConstant.EC2InstanceProfileArn),
			},
			KeyName: aws.String(awsConstant.EC2KeyPairName),
			NetworkInterfaces: []types.InstanceNetworkInterfaceSpecification{
				{
					DeviceIndex:              aws.Int32(0),
					AssociatePublicIpAddress: aws.Bool(true),
					DeleteOnTermination:      aws.Bool(true),
				},
			},
			UserData: aws.String(b64UserData),
		})
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", model.NewInternalServerError(err))) != nil {
				panic("error encoding json")
			}
			return
		}

		// wait for ec2 creation to be complete
		waiter := ec2.NewInstanceRunningWaiter(client)
		instancesOutput, err := waiter.WaitForOutput(r.Context(), &ec2.DescribeInstancesInput{
			InstanceIds: aws.ToStringSlice([]*string{output.Instances[0].InstanceId}),
		}, time.Minute*5)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", model.NewInternalServerError(err))) != nil {
				panic("error encoding json")
			}
			return
		}

		// create an instance record in db
		createdInstance := instancesOutput.Reservations[0].Instances[0]
		i, err := ctrl.Instance.Create(r.Context(), model.CreateInstanceInput{
			InstanceID:       *createdInstance.InstanceId,
			InstanceType:     string(createdInstance.InstanceType),
			PrivateDNSName:   *createdInstance.PrivateDnsName,
			PrivateIPAddress: *createdInstance.PrivateIpAddress,
			PublicDNSName:    *createdInstance.PublicDnsName,
			PublicIPAddress:  *createdInstance.PublicIpAddress,
			ImageID:          *createdInstance.ImageId,
			Architecture:     string(createdInstance.Architecture),
			AvailabilityZone: *createdInstance.Placement.AvailabilityZone,
			UserID:           userID,
			SlideID:          s.ID,
		})
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", model.NewInternalServerError(err))) != nil {
				panic("error encoding json")
			}
			return
		}

		// keep pinging ec2 instance until get html response (verify slidev has started)
		var res *http.Response
		reqURL := fmt.Sprintf("http://%s", *createdInstance.PublicDnsName)
		req, err := http.NewRequest(http.MethodGet, reqURL, nil)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", model.NewInternalServerError(err))) != nil {
				panic("error encoding json")
			}
			return
		}
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		for res == nil || res.StatusCode != http.StatusOK {
			httpClient := http.Client{
				Timeout: 5 * time.Minute,
			}
			res, _ = httpClient.Do(req)
			time.Sleep(1 * time.Second)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if json.NewEncoder(w).Encode(i) != nil {
			panic("error encoding json")
		}
	}).Methods("POST")

	r.HandleFunc("/terminate", func(w http.ResponseWriter, r *http.Request) {
		// parse request body
		var tri TerminateRequestInput
		err := json.NewDecoder(r.Body).Decode(&tri)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", model.NewInvalidRequestBodyError(err))) != nil {
				panic("error encoding json")
			}
			return
		}

		// query instance with slide_id
		ic, err := ctrl.Instance.List(r.Context(), nil, nil, nil, nil, &model.InstanceWhereInput{
			HasSlideWith: []*ent.SlideWhereInput{
				{ID: (*ulid.ID)(&tri.SlideID)},
			},
		}, nil)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", model.NewDBError(err))) != nil {
				panic("error encoding json")
			}
			return
		}
		// unable to find instance
		if ic.TotalCount < 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", model.NewNotFoundError(err, fmt.Sprintf("slide of `%s` has no running ec2 instance", tri.SlideID)))) != nil {
				panic("error encoding json")
			}
			return
		}

		// get the instance
		instance := ic.Edges[0].Node

		// terminate ec2 instance
		_, err = client.TerminateInstances(r.Context(), &ec2.TerminateInstancesInput{
			InstanceIds: []string{string(instance.ID)},
		})
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", model.NewInternalServerError(err))) != nil {
				panic("error encoding json")
			}
			return
		}

		// wait for terminatation of ec2 instance
		waiter := ec2.NewInstanceTerminatedWaiter(client)
		_, err = waiter.WaitForOutput(r.Context(), &ec2.DescribeInstancesInput{
			InstanceIds: []string{string(instance.ID)},
		}, time.Minute*5)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", model.NewInternalServerError(err))) != nil {
				panic("error encoding json")
			}
			return
		}

		// delete instances from db
		_, err = ctrl.Instance.Delete(r.Context(), instance.ID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			if json.NewEncoder(w).Encode(fmt.Sprintf("%v", model.NewInternalServerError(err))) != nil {
				panic("error encoding json")
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if json.NewEncoder(w).Encode(instance) != nil {
			panic("error encoding json")
		}
	}).Methods("POST")

	return r
}
