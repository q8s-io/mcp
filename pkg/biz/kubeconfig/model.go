package kubeconfig

import (
	"github.com/jinzhu/gorm"
)

type Kubeconfigs []Kubeconfig

type Kubeconfig struct {
	gorm.Model
	ClusterID  uint
	Kubeconfig string `gorm:"type:text NOT NULL"`
	Context    string `gorm:"size:30;NOT NULL"`
}

func (k Kubeconfig) TableName() string {
	return "cluster_kubeconfig"
}

type model struct{}

func (model) getByClusterID(db *gorm.DB, id uint) (*Kubeconfig, error) {
	var kubeconfig Kubeconfig
	if err := db.Where("cluster_id=?", id).First(&kubeconfig).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &kubeconfig, nil
}
