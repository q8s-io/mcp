package service

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/q8s-io/mcp/pkg/persistence"
)

type ModelSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
}

func (s *ModelSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	isDebug := os.Getenv("debug")
	if isDebug == "true" {
		s.DB.LogMode(true)
	}
	s.DB.SingularTable(true)

	persistence.NewRepositories(s.DB)
}

func (s *ModelSuite) TearDownSuite() {
	s.DB.Close()
}

func TestService(t *testing.T) {
	suite.Run(t, new(ModelSuite))
}
