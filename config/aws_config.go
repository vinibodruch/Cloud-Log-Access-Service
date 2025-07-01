package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
)

// AWSConfig contains the loaded AWS configuration.
type AWSConfig struct {
	Config aws.Config
}

// LoadAWSConfig loads the AWS configuration from environment variables or the default credentials file.
func LoadAWSConfig() AWSConfig {
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		log.Fatalf("Error loading AWS configuration: %v", err)
	}
	return AWSConfig{Config: cfg}
}
