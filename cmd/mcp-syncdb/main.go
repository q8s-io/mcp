package main

import (
	"flag"
	"reflect"

	"k8s.io/klog"

	"github.com/q8s-io/mcp/pkg/config"
	"github.com/q8s-io/mcp/pkg/domain/entity"
	"github.com/q8s-io/mcp/pkg/persistence"
)

// regist module here to created table in database
var tables = []interface{}{
	entity.Cluster{},
	entity.Kubeconfig{},
	entity.AzureSecret{},
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

	repositories, ok := persistence.InitDB(&conf.MysqlConfig)
	if !ok {
		klog.Errorf("error to init DB, %v", err)
		return
	}
	defer persistence.CloseDB()

	// migrate table
	for _, table := range tables {
		module := reflect.TypeOf(table).Name()
		klog.Infof("auto migrate module[%s]", module)
		err := repositories.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(table).Error
		if err != nil {
			klog.Errorf("error to migrate module[%s]", module)
			return
		}
	}

	klog.Info("success to sync db")
}
