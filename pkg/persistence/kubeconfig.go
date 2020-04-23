package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type KubeconfigRepo struct {
	db *gorm.DB
}

func newKubeconfigRepository(db *gorm.DB) *KubeconfigRepo {
	return &KubeconfigRepo{db}
}

func (r *KubeconfigRepo) GetByClusterID(id uint) (*entity.Kubeconfig, error) {
	var kubeconfig entity.Kubeconfig
	if err := r.db.Where("cluster_id=?", id).First(&kubeconfig).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &kubeconfig, nil
}
