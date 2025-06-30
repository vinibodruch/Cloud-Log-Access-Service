package services

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// S3Service é uma interface para interagir com o S3.
type S3Service interface {
	ListObjectsInBucket(bucketName string) ([]types.Object, error)
	GetObjectFromBucket(bucketName, objectKey string) ([]byte, error)
}

// s3ServiceImpl implementa a interface S3Service.
type s3ServiceImpl struct {
	s3Client *s3.Client
}

// NewS3Service cria e retorna uma nova instância de S3Service.
func NewS3Service(cfg aws.Config) S3Service {
	return &s3ServiceImpl{
		s3Client: s3.NewFromConfig(cfg),
	}
}

// ListObjectsInBucket lista todos os objetos em um bucket S3.
func (s *s3ServiceImpl) ListObjectsInBucket(bucketName string) ([]types.Object, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}

	result, err := s.s3Client.ListObjectsV2(context.TODO(), input)
	if err != nil {
		log.Printf("Erro ao listar objetos no bucket %s: %v", bucketName, err)
		return nil, err
	}
	return result.Contents, nil
}

// GetObjectFromBucket obtém o conteúdo de um objeto específico de um bucket S3.
func (s *s3ServiceImpl) GetObjectFromBucket(bucketName, objectKey string) ([]byte, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	resp, err := s.s3Client.GetObject(context.TODO(), input)
	if err != nil {
		log.Printf("Erro ao obter objeto %s do bucket %s: %v", objectKey, bucketName, err)
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler o corpo do objeto %s: %v", objectKey, err)
		return nil, err
	}
	return buf.Bytes(), nil
}
