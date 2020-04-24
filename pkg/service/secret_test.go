package service

import (
	"errors"
	"testing"
	
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/q8s-io/mcp/pkg/dto"
)


func (s *ModelSuite) Test_CreateSecret() {
	tests := []struct {
		name     string
		req      *dto.SecretAttachReq
		expected *dto.SecretAttachResp
		err      error
		mockFun  func()
	}{
		{
			name: "SecretCreateSuccess",
			req: &dto.SecretAttachReq{
				Name:         "success_create",
				TenantID:     "a3ViZWNvbmZpZy1zdHJpbmc",
				ClientID:     "mZpZy1zdHJiZWNvb",
				ClientSecret: "default-contex",
			},
			expected: &dto.SecretAttachResp{
				ID: 1,
			},
			err: nil,
			mockFun: func() {

				s.mock.ExpectBegin()
				s.mock.ExpectExec("INSERT INTO `azure_secret`").
					WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectCommit()

			},
		},

		{
			name: "secretRollback",
			req: &dto.SecretAttachReq{
				Name:         "rollback_secret",
				TenantID:     "a3ViZWNvbmZpZy1zdHJpbmc",
				ClientID:     "mZpZy1zdHJiZWNvb",
				ClientSecret: "default-contex",
			},
			expected: nil,
			err:      errors.New("some error"),
			mockFun: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec("INSERT INTO `azure_secret`").
					WillReturnError(errors.New("some error"))
				s.mock.ExpectRollback()
			},
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			// prepare mock sql
			test.mockFun()
			resp, err := NewSecretService().SecretCreate(test.req)

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, resp)
		})
	}
}