// +build integration

package integration

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type clusterTest struct {
	name string
}

func newClusterTest() tester {
	return &clusterTest{
		name: "cluster_test",
	}
}

func (ct *clusterTest) runTest(t *testing.T) {
	t.Run(ct.name, test_get_ok)
}

func test_get_ok(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		method     string
		statusCode int
	}{
		{
			name:       "get_all_clusters_ok",
			url:        "/api/v1/clusters",
			method:     http.MethodGet,
			statusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, test.url, nil)
			assert.NoError(t, err)
			rec := httptest.NewRecorder()
			container.Dispatch(rec, req)

			res := rec.Result()
			assert.Equal(t, test.statusCode, res.StatusCode)
			defer res.Body.Close()

			_, err = ioutil.ReadAll(res.Body)
			assert.NoError(t, err)
		})
	}
}
