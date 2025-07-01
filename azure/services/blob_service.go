package services

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	//"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
)

// BlobService is the interface for interacting with Azure Blob Storage.
type BlobService interface {
	ListContainers() ([]string, error)
	ListBlobsInContainer(containerName string) ([]string, error)
}

// blobServiceImpl implements the BlobService interface.
type blobServiceImpl struct {
	client *azblob.Client
}

// NewBlobService creates and returns a new instance of BlobService.
func NewBlobService(client *azblob.Client) BlobService {
	return &blobServiceImpl{
		client: client,
	}
}

// ListContainers lists all blob containers in the storage account.
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

// ListBlobsInContainer lists all blobs in a specific container.
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
