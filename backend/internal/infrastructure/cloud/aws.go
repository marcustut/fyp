package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// NewAWSConfig loads the credential values from environment
// variables and returns aws-sdk's default configuration.
func NewAWSConfig() (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithDefaultRegion("ap-southeast-1"))
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// NewS3Client initialize a new aws-s3 API client with the
// provided config.
func NewS3Client(cfg aws.Config) *s3.Client {
	client := s3.NewFromConfig(cfg)
	return client
}
