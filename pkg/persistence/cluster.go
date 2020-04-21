package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type Cluster struct {
	base
}

func NewClusterPersistence() *Cluster {
	return &Cluster{}
}

func (c *Cluster) GetAll() (entity.Clusters, error) {
	var clusters entity.Clusters
	if err := c.GetDB().Find(&clusters).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return clusters, nil
}

func (c *Cluster) GetByID(id uint) (*entity.Cluster, error) {
	var cluster entity.Cluster
	if err := c.GetDB().Where("id=?", id).First(&cluster).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &cluster, nil
}

func (c *Cluster) Add(cluster *entity.Cluster) error {
	return c.GetDB().Create(cluster).Error
}
