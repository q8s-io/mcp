package repository

import (
	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type ClusterRepository interface {
	// GetAll gets all Cluster in db.
	GetAll() (entity.Clusters, error)
	// GetByID gets Cluster with specified id in db, nil if not found
	GetByID(id uint) (*entity.Cluster, error)
	// Add creates Cluster struct to db.
	Add(cluster *entity.Cluster) error
}
