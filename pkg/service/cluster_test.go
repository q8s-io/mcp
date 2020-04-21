package service

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/q8s-io/mcp/pkg/dto"
)

func (s *ModelSuite) Test_GetByID() {
	const sqlSelect = "SELECT \\* FROM `cluster`"

	tests := []struct {
		name     string
		id       uint
		expected *dto.ClusterDetailResp
		err      error
		mockFun  func()
	}{
		{
			name: "getOneResult",
			id:   1,
			expected: &dto.ClusterDetailResp{
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
			cluster, err := NewClusterService().GetByID(test.id)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, cluster)
		})
	}
}

func (s *ModelSuite) Test_All() {
	sql := "SELECT \\* FROM `cluster`"

	tests := []struct {
		name     string
		expected []dto.ClusterListResp
		err      error
		mockFun  func()
	}{
		{
			name: "getOneValue",
			expected: []dto.ClusterListResp{
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
			expected: []dto.ClusterListResp{
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
			expected: []dto.ClusterListResp{},
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
			clusters, err := NewClusterService().All()
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, clusters)
		})
	}
}

func (s *ModelSuite) Test_Attach() {
	tests := []struct {
		name     string
		req      *dto.ClusterAttachReq
		expected *dto.ClusterAttachResp
		err      error
		mockFun  func()
	}{
		{
			name: "attachSuccess",
			req: &dto.ClusterAttachReq{
				Name:       "success_cluster",
				Kubeconfig: "a3ViZWNvbmZpZy1zdHJpbmc=",
				Context:    "default-context",
			},
			expected: &dto.ClusterAttachResp{
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
			req: &dto.ClusterAttachReq{
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
			resp, err := NewClusterService().Attach(test.req)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, resp)
		})
	}
}
