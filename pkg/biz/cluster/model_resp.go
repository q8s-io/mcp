package cluster

import (
	"time"
)

type ListResp struct {
	ID        uint      `json:"id" description:"ID of attached cluster."`
	Name      string    `json:"name" description:"Name of attached cluster."`
	Status    string    `json:"status" description:"Status of attached cluster."`
	CreatedAt time.Time `json:"createdAt" description:"Created time of attached cluster."`
	Provider  string    `json:"provider" description:"Provider of attached cluster. Empty if no provider."`
}

func (r *ListResp) convertFromCluster(cluster *Cluster) {
	r.ID = cluster.ID
	r.Name = cluster.Name
	r.Status = cluster.Status.StatusString()
	r.Provider = cluster.Provider
	r.CreatedAt = cluster.CreatedAt
}

type DetailResp struct {
	ID        uint      `json:"id" description:"ID of attached cluster."`
	Name      string    `json:"name" description:"Name of attached cluster."`
	Status    string    `json:"status" description:"Status of attached cluster."`
	CreatedAt time.Time `json:"createdAt" description:"Created time of attached cluster."`
	Provider  string    `json:"provider" description:"Provider of attached cluster. Empty if no provider."`
}

func (r *DetailResp) convertFromCluster(cluster *Cluster) {
	r.ID = cluster.ID
	r.Name = cluster.Name
	r.Status = cluster.Status.StatusString()
	r.Provider = cluster.Provider
	r.CreatedAt = cluster.CreatedAt
}

type AttachResp struct {
	ID           uint `json:"id" description:"ID of cluster."`
	KubeconfigID uint `json:"kubeconfigId" description:"ID of kubeconfig."`
}

func (r *AttachResp) convertFromCluster(cluster *Cluster) {
	r.ID = cluster.ID
	r.KubeconfigID = cluster.Kubeconfig.ID
}
