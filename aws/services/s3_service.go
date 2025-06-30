package services

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// S3Service is an interface for interacting with S3.
type S3Service interface {
	ListObjectsInBucket(bucketName string) ([]types.Object, error)
	GetObjectFromBucket(bucketName, objectKey string) ([]byte, error)
	// Add other methods as needed (upload, delete, etc.)
}

// s3ServiceImpl implements the S3Service interface.
type s3ServiceImpl struct {
	s3Client *s3.Client
}

// NewS3Service creates and returns a new instance of S3Service.
func NewS3Service(cfg aws.Config) S3Service {
	return &s3ServiceImpl{
		s3Client: s3.NewFromConfig(cfg),
	}
}

// ListObjectsInBucket lists all objects in an S3 bucket.
func (s *s3ServiceImpl) ListObjectsInBucket(bucketName string) ([]types.Object, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}

	result, err := s.s3Client.ListObjectsV2(context.TODO(), input)
	if err != nil {
		log.Printf("Error listing objects in bucket %s: %v", bucketName, err)
		return nil, err
	}
	return result.Contents, nil
}

// GetObjectFromBucket retrieves the content of a specific object from an S3 bucket.
func (s *s3ServiceImpl) GetObjectFromBucket(bucketName, objectKey string) ([]byte, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	resp, err := s.s3Client.GetObject(context.TODO(), input)
	if err != nil {
		log.Printf("Error getting object %s from bucket %s: %v", objectKey, bucketName, err)
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		log.Printf("Error reading body of object %s: %v", objectKey, err)
		return nil, err
	}
	return buf.Bytes(), nil
}
