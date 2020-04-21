package entity

import (
	"github.com/jinzhu/gorm"
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
	Kubeconfig Kubeconfig
	Name       string `gorm:"size:20;NOT NULL;unique"`
	Provider   string `gorm:"size:20"`
	Status     Status `gorm:"type:tinyint unsigned NOT NULL"`
}
