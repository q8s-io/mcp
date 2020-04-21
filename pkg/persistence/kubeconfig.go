package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type Kubeconfig struct {
	base
}

func NewKubeconfigPersistence() *Kubeconfig {
	return &Kubeconfig{}
}

func (k *Kubeconfig) GetByClusterID(id uint) (*entity.Kubeconfig, error) {
	var kubeconfig entity.Kubeconfig
	if err := k.GetDB().Where("cluster_id=?", id).First(&kubeconfig).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &kubeconfig, nil
}
