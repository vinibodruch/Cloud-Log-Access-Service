package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
)

// AWSConfig contém a configuração AWS carregada.
type AWSConfig struct {
	Config aws.Config
}

// LoadAWSConfig carrega a configuração AWS a partir das variáveis de ambiente ou do arquivo de credenciais padrão.
func LoadAWSConfig() AWSConfig {
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		log.Fatalf("Erro ao carregar a configuração AWS: %v", err)
	}
	return AWSConfig{Config: cfg}
}
