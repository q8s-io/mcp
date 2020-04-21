package persistence

import (
	"github.com/jinzhu/gorm"
)

type base struct {
	db *gorm.DB
}

func (b *base) GetDB() *gorm.DB {
	if b.db == nil {
		b.db = dbPool
	}
	return b.db
}
