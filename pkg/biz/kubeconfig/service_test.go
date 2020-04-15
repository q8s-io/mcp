package kubeconfig

import (
	"database/sql"
	"errors"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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
}

func (s *ModelSuite) TearDownSuite() {
	defer s.DB.Close()
}

func TestKubeconfig(t *testing.T) {
	suite.Run(t, new(ModelSuite))
}

func (s *ModelSuite) Test_GetByClusterID() {
	const sqlSelect = "SELECT \\* FROM `cluster_kubeconfig`"

	tests := []struct {
		name     string
		id       uint
		expected *ClusterKubeconfigResp
		err      error
		mockFun  func()
	}{
		{
			name: "getOneValue",
			id:   1,
			expected: &ClusterKubeconfigResp{
				Kubeconfig: "dGVzdA==",
				Context:    "test",
			},
			err: nil,
			mockFun: func() {
				s.mock.ExpectQuery(sqlSelect).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"kubeconfig", "context"}).AddRow("test", "test"))
			},
		},
		{
			name:     "getNoValue",
			id:       2,
			expected: nil,
			err:      nil,
			mockFun: func() {
				s.mock.ExpectQuery(sqlSelect).
					WithArgs(2).
					WillReturnRows(sqlmock.NewRows(nil))
			},
		},
		{
			name:     "getWithError",
			id:       3,
			expected: nil,
			err:      errors.New("some error"),
			mockFun: func() {
				s.mock.ExpectQuery(sqlSelect).
					WithArgs(3).
					WillReturnError(errors.New("some error"))
			},
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			test.mockFun()
			kubeconfigResp, err := GetService().GetByClusterID(s.DB, test.id)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, kubeconfigResp)
		})
	}
}
