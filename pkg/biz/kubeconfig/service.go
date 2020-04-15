package kubeconfig

import (
	"github.com/jinzhu/gorm"
)

type Service struct{}

var serviceInstance *Service

func GetService() *Service {
	if serviceInstance == nil {
		serviceInstance = &Service{}
	}
	return serviceInstance
}

func (s *Service) GetByClusterID(db *gorm.DB, clusterID uint) (*ClusterKubeconfigResp, error) {
	kubeconfig, err := model{}.getByClusterID(db, clusterID)
	if err != nil {
		return nil, err
	}
	if kubeconfig == nil {
		return nil, nil
	}

	clusterKubeconfigResp := &ClusterKubeconfigResp{}
	if err := clusterKubeconfigResp.convertFromKubeconfig(kubeconfig); err != nil {
		return nil, err
	}
	return clusterKubeconfigResp, nil
}
