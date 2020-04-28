package entity

import (
	"github.com/jinzhu/gorm"
)

type AzureSecrets []AzureSecret

type AzureSecret struct {
	gorm.Model
	Name          string    `gorm:"size:20;NOT NULL;unique"`
	TenantID      string    `gorm:"size:30;NOT NULL"`
	ClientID      string    `gorm:"size:30;NOT NULL"`
	ClientSecret  string    `gorm:"type:text NOT NULL"`
}
