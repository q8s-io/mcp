package dto

import (

	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type SecretAttachReq struct {
	NAME         string `json:"name" validate:"required" description:"Name ID of Azureclient."`
	TenantID     string `json:"tenant_id" validate:"required" description:"Tenant ID of Azure."`
	ClientID     string `json:"client_id" validate:"required" description:"Client ID of Azure."`
	ClientSecret string `json:"client_secret" validate:"required" description:"Client secret of Azure."`
}

type SecretAttachResp struct {
	NAME         string `json:"name" validate:"required" description:"Name ID of Azureclient."`
	TenantID     string `json:"tenant_id" validate:"required" description:"Tenant ID of Azure."`
	ClientID     string `json:"client_id" validate:"required" description:"Client ID of Azure."`
	ClientSecret string `json:"client_secret" validate:"required" description:"Client secret of Azure."`
}

func (s *SecretAttachReq) ConvertToCluster() (*entity.Secret_Azure,error) {
	return &entity.Secret_Azure{
		NAME: s.NAME,
		TenantID: s.TenantID,
		ClientID: s.ClientID,
		ClientSecret: s.ClientSecret,
	},nil
}

//
//func (r *ClusterAttachReq) ConvertToCluster() (*entity.Cluster, error) {
//	decodeBytes, err := base64.StdEncoding.DecodeString(r.Kubeconfig)
//	if err != nil {
//		return nil, err
//	}
//
//	return &entity.Cluster{
//		Name:     r.Name,
//		Provider: entity.AttachedProvider,
//		Status:   entity.StatusOk,
//		Kubeconfig: entity.Kubeconfig{
//			Kubeconfig: string(decodeBytes),
//			Context:    r.Context,
//		},
//	}, nil
//}