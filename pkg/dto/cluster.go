package dto

import (
	"encoding/base64"
	"time"

	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type ClusterAttachReq struct {
	Name       string `json:"name" validate:"required" description:"Name of attached cluster."`
	Kubeconfig string `json:"kubeconfig" validate:"required" description:"Kubeconfig info with Base64 format."`
	Context    string `json:"context" validate:"required" description:"Context of kubeconfig."`
}

func (r *ClusterAttachReq) ConvertToClusterKubeconfig() (*entity.Kubeconfig, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(r.Kubeconfig)
	if err != nil {
		return nil, err
	}

	return &entity.Kubeconfig{
		Kubeconfig: string(decodeBytes),
		Context:    r.Context,
	}, nil
}

func (r *ClusterAttachReq) ConvertToCluster() (*entity.Cluster, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(r.Kubeconfig)
	if err != nil {
		return nil, err
	}

	return &entity.Cluster{
		Name:     r.Name,
		Provider: entity.AttachedProvider,
		Status:   entity.StatusOk,
		Kubeconfig: entity.Kubeconfig{
			Kubeconfig: string(decodeBytes),
			Context:    r.Context,
		},
	}, nil
}

type ClusterListResp struct {
	ID        uint      `json:"id" description:"ID of attached cluster."`
	Name      string    `json:"name" description:"Name of attached cluster."`
	Status    string    `json:"status" description:"Status of attached cluster."`
	CreatedAt time.Time `json:"createdAt" description:"Created time of attached cluster."`
	Provider  string    `json:"provider" description:"Provider of attached cluster. Empty if no provider."`
}

func (r *ClusterListResp) ConvertFromCluster(cluster *entity.Cluster) {
	r.ID = cluster.ID
	r.Name = cluster.Name
	r.Status = cluster.Status.StatusString()
	r.Provider = cluster.Provider
	r.CreatedAt = cluster.CreatedAt
}

type ClusterDetailResp struct {
	ID        uint      `json:"id" description:"ID of attached cluster."`
	Name      string    `json:"name" description:"Name of attached cluster."`
	Status    string    `json:"status" description:"Status of attached cluster."`
	CreatedAt time.Time `json:"createdAt" description:"Created time of attached cluster."`
	Provider  string    `json:"provider" description:"Provider of attached cluster. Empty if no provider."`
}

func (r *ClusterDetailResp) ConvertFromCluster(cluster *entity.Cluster) {
	r.ID = cluster.ID
	r.Name = cluster.Name
	r.Status = cluster.Status.StatusString()
	r.Provider = cluster.Provider
	r.CreatedAt = cluster.CreatedAt
}

type ClusterAttachResp struct {
	ID           uint `json:"id" description:"ID of cluster."`
	KubeconfigID uint `json:"kubeconfigId" description:"ID of kubeconfig."`
}

func (r *ClusterAttachResp) ConvertFromCluster(cluster *entity.Cluster) {
	r.ID = cluster.ID
	r.KubeconfigID = cluster.Kubeconfig.ID
}
