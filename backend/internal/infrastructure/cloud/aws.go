package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
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
