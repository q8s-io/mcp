package cluster

import (
	"github.com/jinzhu/gorm"

	"github.com/q8s-io/mcp/pkg/biz/kubeconfig"
)

type Status uint

const (
	StatusInit Status = iota + 1
	StatusOk
)

const AttachedProvider = ""

func (s Status) StatusString() string {
	switch s {
	case StatusInit:
		return "init"
	case StatusOk:
		return "ok"
	default:
		return ""
	}
}

type Clusters []Cluster

type Cluster struct {
	gorm.Model
	Kubeconfig kubeconfig.Kubeconfig
	Name       string `gorm:"size:20;NOT NULL;unique"`
	Provider   string `gorm:"size:20"`
	Status     Status `gorm:"type:tinyint unsigned NOT NULL"`
}

type model struct{}

func (model) getAll(db *gorm.DB) (Clusters, error) {
	var clusters Clusters
	if err := db.Find(&clusters).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return clusters, nil
}

func (model) getByID(db *gorm.DB, id uint) (*Cluster, error) {
	var cluster Cluster
	if err := db.Where("id=?", id).First(&cluster).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &cluster, nil
}

func (model) add(db *gorm.DB, cluster *Cluster) error {
	return db.Create(cluster).Error
}
