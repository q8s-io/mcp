package service

import (
	"github.com/q8s-io/mcp/pkg/domain/repository"
	"github.com/q8s-io/mcp/pkg/dto"
	"github.com/q8s-io/mcp/pkg/persistence"
)

type Kubeconfig struct {
	kubeconfigRepo repository.KubeconfigRepository
}

func NewKubeconfigService() *Kubeconfig {
	return &Kubeconfig{
		kubeconfigRepo: persistence.GetRepositories().KubeconfigRepo,
	}
}

func (k *Kubeconfig) GetByClusterID(clusterID uint) (*dto.KubeconfigResp, error) {
	kubeconfig, err := k.kubeconfigRepo.GetByClusterID(clusterID)
	if err != nil {
		return nil, err
	}
	if kubeconfig == nil {
		return nil, nil
	}

	clusterKubeconfigResp := &dto.KubeconfigResp{}
	if err := clusterKubeconfigResp.ConvertFromKubeconfig(kubeconfig); err != nil {
		return nil, err
	}
	return clusterKubeconfigResp, nil
}
