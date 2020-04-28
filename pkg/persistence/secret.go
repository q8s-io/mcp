package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type AzureSecretRepo struct {
	db *gorm.DB
}

//type AzureSecret struct {
//	base
//}
func newSecretRepository(db *gorm.DB) *AzureSecretRepo {
	return &AzureSecretRepo{db}
}

func (s *AzureSecretRepo) Add(secret *entity.AzureSecret) error{
	return s.db.Create(secret).Error
}

