package repository

import (
	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type KubeconfigRepository interface {
	GetByClusterID(id uint) (*entity.Kubeconfig, error)
}
