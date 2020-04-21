package service

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/q8s-io/mcp/pkg/dto"
)

func (s *ModelSuite) Test_GetByClusterID() {
	const sqlSelect = "SELECT \\* FROM `cluster_kubeconfig`"

	tests := []struct {
		name     string
		id       uint
		expected *dto.KubeconfigResp
		err      error
		mockFun  func()
	}{
		{
			name: "getOneValue",
			id:   1,
			expected: &dto.KubeconfigResp{
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
			kubeconfigResp, err := NewKubeconfigService().GetByClusterID(test.id)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, kubeconfigResp)
		})
	}
}
