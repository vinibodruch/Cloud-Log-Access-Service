package services

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	//"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
)

// BlobService é a interface para interagir com o Azure Blob Storage.
type BlobService interface {
	ListContainers() ([]string, error)
	ListBlobsInContainer(containerName string) ([]string, error)
}

// blobServiceImpl implementa a interface BlobService.
type blobServiceImpl struct {
	client *azblob.Client
}

// NewBlobService cria e retorna uma nova instância de BlobService.
func NewBlobService(client *azblob.Client) BlobService {
	return &blobServiceImpl{
		client: client,
	}
}

// ListContainers lista todos os containers de blob na conta de armazenamento.
func (s *blobServiceImpl) ListContainers() ([]string, error) {
	var containers []string
	pager := s.client.NewListContainersPager(&azblob.ListContainersOptions{})
	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Printf("Error listing Azure containers: %v", err)
			return nil, err
		}
		for _, container := range resp.ContainerItems {
			containers = append(containers, *container.Name)
		}
	}
	return containers, nil
}

// ListBlobsInContainer lista todos os blobs em um container específico.
func (s *blobServiceImpl) ListBlobsInContainer(containerName string) ([]string, error) {
	var blobs []string
	containerClient := s.client.ServiceClient().NewContainerClient(containerName)
	pager := containerClient.NewListBlobsFlatPager(&container.ListBlobsFlatOptions{})
	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Printf("Error listing Azure blobs in container %s: %v", containerName, err)
			return nil, err
		}
		for _, blob := range resp.Segment.BlobItems {
			blobs = append(blobs, *blob.Name)
		}
	}
	return blobs, nil
}
