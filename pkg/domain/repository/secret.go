package repository

import (
	"github.com/q8s-io/mcp/pkg/domain/entity"
)

type AzureSecretRepository interface {
	// Add creates Cluster struct to db.
	Add(secret *entity.AzureSecret) error
}