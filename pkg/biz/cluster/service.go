package cluster

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

func (s *Service) All(db *gorm.DB) ([]ListResp, error) {
	clusters, err := model{}.getAll(db)
	if err != nil {
		return nil, err
	}

	response := make([]ListResp, len(clusters))
	for index, cluster := range clusters {
		clusterResp := ListResp{}
		clusterResp.convertFromCluster(&cluster)
		response[index] = clusterResp
	}
	return response, nil
}

func (s *Service) GetByID(db *gorm.DB, id uint) (*DetailResp, error) {
	cluster, err := model{}.getByID(db, id)
	if err != nil {
		return nil, err
	}
	if cluster == nil {
		return nil, nil
	}

	response := &DetailResp{}
	response.convertFromCluster(cluster)
	return response, nil
}

func (s *Service) Attach(db *gorm.DB, attachCluster *AttachReq) (*AttachResp, error) {
	cluster, err := attachCluster.toCluster()
	if err != nil {
		return nil, err
	}

	err = model{}.add(db, cluster)
	if err != nil {
		return nil, err
	}

	response := &AttachResp{}
	response.convertFromCluster(cluster)
	return response, nil
}
