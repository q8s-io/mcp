package entity

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
