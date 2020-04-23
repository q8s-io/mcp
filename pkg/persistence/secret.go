package persistence

import (
	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type Secret_Azure struct {
	base
}

func NewSecretPersistence() *Secret_Azure {
	return &Secret_Azure{}
}

func (s *Secret_Azure) Add(secret *entity.Secret_Azure) error{
	return s.GetDB().Create(secret).Error
}