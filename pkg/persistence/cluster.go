package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type ClusterRepo struct {
	db *gorm.DB
}

func newClusterRepository(db *gorm.DB) *ClusterRepo {
	return &ClusterRepo{db}
}

func (r *ClusterRepo) GetAll() (entity.Clusters, error) {
	var clusters entity.Clusters
	if err := r.db.Find(&clusters).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return clusters, nil
}

func (r *ClusterRepo) GetByID(id uint) (*entity.Cluster, error) {
	var cluster entity.Cluster
	if err := r.db.Where("id=?", id).First(&cluster).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &cluster, nil
}

func (r *ClusterRepo) Add(cluster *entity.Cluster) error {
	return r.db.Create(cluster).Error
}
