package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/joho/godotenv"
)

// LoadAWSConfig Load configurations from env variables or .env file for AWS SDK
func LoadAWSConfig() aws.Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Could not load the .env file. AWS environment variables will be used if set.")
	}

	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		log.Fatalf("Error loading AWS configuration: %v", err)
	}
	return cfg
}
