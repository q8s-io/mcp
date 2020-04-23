package service

import (
	"github.com/q8s-io/mcp/pkg/dto"
	"github.com/q8s-io/mcp/pkg/persistence"
)

type Secret_Azure struct{
	secretPersistence *persistence.Secret_Azure
}

func NewSecretService() *Secret_Azure {
	return &Secret_Azure{
		secretPersistence: persistence.NewSecretPersistence(),
	}
}

func (s *Secret_Azure) Attach(attachSecret *dto.SecretAttachReq) (*dto.SecretAttachReq, error) {
	secret, err := attachSecret.ConvertToCluster()
	if err != nil {
		return nil, err
	}
	err = s.secretPersistence.Add(secret)
	if err != nil {
		return nil, err
	}
	//response := &dto.SecretAttachResp{}
	//response.convertFromCluster(attachSecret)
	return attachSecret, nil
}
