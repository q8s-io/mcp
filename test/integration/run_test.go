// +build integration

package integration

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/emicklei/go-restful/v3"
	"github.com/stretchr/testify/assert"

	"github.com/q8s-io/mcp/pkg/config"
	"github.com/q8s-io/mcp/pkg/persistence"
	"github.com/q8s-io/mcp/pkg/server"
)

var (
	container *restful.Container
)

type tester interface {
	runTest(t *testing.T)
}

func before(t *testing.T) {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile != "" && !filepath.IsAbs(configFile) {
		configFile = path.Join("..", "..", configFile)
	}

	// load config from env
	conf, err := config.LoadConfig(configFile)
	assert.NoError(t, err)

	mysqlConfig := &conf.MysqlConfig
	if _, ok := persistence.InitDB(mysqlConfig); !ok {
		t.Errorf("error to init DB")
		return
	}

	container = server.NewContainer()
}

func after(t *testing.T) {
	persistence.CloseDB()
}

func TestRun(t *testing.T) {
	before(t)
	defer after(t)

	newClusterTest().runTest(t)
}
