package service

import (
	"github.com/q8s-io/mcp/pkg/dto"
	"github.com/q8s-io/mcp/pkg/persistence"
	"github.com/q8s-io/mcp/pkg/domain/repository"
)

type AzureSecret struct
{
	AzureSecretRepo repository.AzureSecretRepository
}

func NewSecretService() *AzureSecret {
	return &AzureSecret{
		AzureSecretRepo: persistence.GetRepositories().AzureSecretRepo,
	}
}
func (s *AzureSecret) SecretCreate(attachSecret *dto.SecretAttachReq) (*dto.SecretAttachResp, error) {
	secret := attachSecret.ConvertToCluster()
	err := s.AzureSecretRepo.Add(secret)
	if err != nil {
		return nil, err
	}

	//response := &dto.SecretAttachResp{}
	//response.convertFromCluster(attachSecret)
	return &dto.SecretAttachResp{ID:1}, nil
}
