package dto

import (

	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type SecretAttachReq struct {
	Name         string `json:"Name" validate:"required" description:"Name ID of Azureclient."`
	TenantID     string `json:"TenantId" validate:"required" description:"Tenant ID of Azure."`
	ClientID     string `json:"ClientId" validate:"required" description:"Client ID of Azure."`
	ClientSecret string `json:"ClientSecret" validate:"required" description:"Client secret of Azure."`
}

type SecretAttachResp struct {
	ID   int
}

func (s *SecretAttachReq) ConvertToCluster() (*entity.AzureSecret) {
	return &entity.AzureSecret{
		Name: s.Name,
		TenantID: s.TenantID,
		ClientID: s.ClientID,
		ClientSecret: s.ClientSecret,
	}
}
