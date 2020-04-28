package service

import (
	"github.com/q8s-io/mcp/pkg/domain/repository"
	"github.com/q8s-io/mcp/pkg/dto"
	"github.com/q8s-io/mcp/pkg/persistence"
)

type Cluster struct {
	clusterRepo repository.ClusterRepository
}

func NewClusterService() *Cluster {
	return &Cluster{
		clusterRepo: persistence.GetRepositories().ClusterRepo,
	}
}

func (c *Cluster) All() ([]dto.ClusterListResp, error) {
	clusters, err := c.clusterRepo.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.ClusterListResp, len(clusters))
	for index, cluster := range clusters {
		clusterResp := dto.ClusterListResp{}
		clusterResp.ConvertFromCluster(&cluster)
		response[index] = clusterResp
	}
	return response, nil
}

func (c *Cluster) GetByID(id uint) (*dto.ClusterDetailResp, error) {
	cluster, err := c.clusterRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if cluster == nil {
		return nil, nil
	}

	response := &dto.ClusterDetailResp{}
	response.ConvertFromCluster(cluster)
	return response, nil
}

func (c *Cluster) Attach(attachCluster *dto.ClusterAttachReq) (*dto.ClusterAttachResp, error) {
	cluster, err := attachCluster.ConvertToCluster()
	if err != nil {
		return nil, err
	}

	err = c.clusterRepo.Add(cluster)
	if err != nil {
		return nil, err
	}

	response := &dto.ClusterAttachResp{}
	response.ConvertFromCluster(cluster)
	return response, nil
}
