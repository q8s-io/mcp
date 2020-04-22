package repository

import (
	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type ClusterRepository interface {
	GetAll() (entity.Clusters, error)
	GetByID(id uint) (*entity.Cluster, error)
	Add(cluster *entity.Cluster) error
}
