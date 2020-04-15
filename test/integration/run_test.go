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
	"github.com/q8s-io/mcp/pkg/db/mysql"
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
	mysql.InitDB(mysqlConfig)

	container = server.NewContainer()
}

func after(t *testing.T) {
	mysql.CloseDB()
}

func TestRun(t *testing.T) {
	before(t)
	defer after(t)

	newClusterTest().runTest(t)
}
