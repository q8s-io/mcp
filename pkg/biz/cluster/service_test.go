package cluster

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
	s.DB.Close()
}

func TestCluster(t *testing.T) {
	suite.Run(t, new(ModelSuite))
}

func (s *ModelSuite) Test_GetByID() {
	const sqlSelect = "SELECT \\* FROM `cluster`"

	tests := []struct {
		name     string
		id       uint
		expected *DetailResp
		err      error
		mockFun  func()
	}{
		{
			name: "getOneResult",
			id:   1,
			expected: &DetailResp{
				Name: "test",
			},
			err: nil,
			mockFun: func() {
				s.mock.ExpectQuery(sqlSelect).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("test"))
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
			// prepare mock sql
			test.mockFun()
			cluster, err := GetService().GetByID(s.DB, test.id)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, cluster)
		})
	}
}

func (s *ModelSuite) Test_All() {
	sql := "SELECT \\* FROM `cluster`"

	tests := []struct {
		name     string
		expected []ListResp
		err      error
		mockFun  func()
	}{
		{
			name: "getOneValue",
			expected: []ListResp{
				{
					Name: "test",
				},
			},
			err: nil,
			mockFun: func() {
				s.mock.ExpectQuery(sql).
					WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("test"))
			},
		},
		{
			name: "getTwoValues",
			expected: []ListResp{
				{
					Name: "test1",
				},
				{
					Name: "test2",
				},
			},
			err: nil,
			mockFun: func() {
				s.mock.ExpectQuery(sql).
					WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("test1").AddRow("test2"))
			},
		},
		{
			name:     "getNoValue",
			expected: []ListResp{},
			err:      nil,
			mockFun: func() {
				s.mock.ExpectQuery(sql).
					WillReturnRows(sqlmock.NewRows(nil))
			},
		},
		{
			name:     "getWithError",
			expected: nil,
			err:      errors.New("some error"),
			mockFun: func() {
				s.mock.ExpectQuery(sql).
					WillReturnError(errors.New("some error"))
			},
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			// prepare mock sql
			test.mockFun()
			clusters, err := GetService().All(s.DB)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, clusters)
		})
	}
}

func (s *ModelSuite) Test_Attach() {
	tests := []struct {
		name     string
		req      *AttachReq
		expected *AttachResp
		err      error
		mockFun  func()
	}{
		{
			name: "attachSuccess",
			req: &AttachReq{
				Name:       "success_cluster",
				Kubeconfig: "a3ViZWNvbmZpZy1zdHJpbmc=",
				Context:    "default-context",
			},
			expected: &AttachResp{
				ID:           5,
				KubeconfigID: 40,
			},
			err: nil,
			mockFun: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec("INSERT INTO `cluster`").
					WillReturnResult(sqlmock.NewResult(5, 1))
				s.mock.ExpectExec("INSERT INTO `cluster_kubeconfig`").
					WillReturnResult(sqlmock.NewResult(40, 1))
				s.mock.ExpectCommit()
			},
		},
		{
			name: "attachRollback",
			req: &AttachReq{
				Name:       "rollback_cluster",
				Kubeconfig: "a3ViZWNvbmZpZy1zdHJpbmc=",
				Context:    "default-context",
			},
			expected: nil,
			err:      errors.New("rollback error"),
			mockFun: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec("INSERT INTO `cluster`").
					WillReturnResult(sqlmock.NewResult(5, 1))
				s.mock.ExpectExec("INSERT INTO `cluster_kubeconfig`").
					WillReturnError(errors.New("rollback error"))
				s.mock.ExpectRollback()
			},
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			// prepare mock sql
			test.mockFun()
			resp, err := GetService().Attach(s.DB, test.req)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, resp)
		})
	}
}
