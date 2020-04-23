// +build integration

package integration

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/q8s-io/mcp/pkg/dto"
)

type secretTest struct {
	name string
}

func newSecretTest() tester {
	return &secretTest{
		name: "secret_test",
	}
}

func (se *secretTest) runTest(t *testing.T) {
	t.Run(se.name, test_post_ok)
}

func test_post_ok(t *testing.T) {
	data,_ := json.Marshal(dto.SecretAttachReq{NAME:"success",TenantID:"success_secret",ClientID:"success_secret-context",ClientSecret:"success_secret"})
	req, err := http.NewRequest(http.MethodPost, "/api/v1/secret", bytes.NewReader(data))
	assert.NoError(t, err)
	req.Header.Add("content-type","application/json")
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	container.Dispatch(rec, req)

	res := rec.Result()
	assert.Equal(t, 201, res.StatusCode)
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
}