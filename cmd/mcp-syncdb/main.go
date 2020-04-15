package main

import (
	"flag"
	"reflect"

	"k8s.io/klog"

	"github.com/q8s-io/mcp/pkg/biz/cluster"
	"github.com/q8s-io/mcp/pkg/biz/kubeconfig"
	"github.com/q8s-io/mcp/pkg/config"
	"github.com/q8s-io/mcp/pkg/db/mysql"
)

// regist module here to created table in database
var tables = []interface{}{
	cluster.Cluster{},
	kubeconfig.Kubeconfig{},
}

func main() {
	var (
		configFile string
	)

	flag.StringVar(
		&configFile,
		"config",
		"",
		"The config file need to be loaded.",
	)

	flag.Parse()

	// load config
	conf, err := config.LoadConfig(configFile)
	if err != nil {
		klog.Errorf("error to load config file: '%s', %v", configFile, err)
		return
	}

	mysql.InitDB(&conf.MysqlConfig)

	db := mysql.GetDB()
	defer mysql.CloseDB()

	// migrate table
	for _, table := range tables {
		module := reflect.TypeOf(table).Name()
		klog.Infof("auto migrate module[%s]", module)
		err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(table).Error
		if err != nil {
			klog.Errorf("error to migrate module[%s]", module)
			return
		}
	}

	klog.Info("success to sync db")
}
