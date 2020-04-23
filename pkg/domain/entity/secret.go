package entity

import (
	"github.com/jinzhu/gorm"
)

type Secret_Azures []Secret_Azure

type Secret_Azure struct {
	gorm.Model
	NAME          string    `gorm:"type:text NOT NULL"`
	TenantID      string    `gorm:"type:text NOT NULL"`
	ClientID      string    `gorm:"type:text NOT NULL"`
	ClientSecret  string    `gorm:"type:text NOT NULL"`
}
