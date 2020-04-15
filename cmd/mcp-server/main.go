package main

import (
	"flag"

	"k8s.io/klog"

	"github.com/q8s-io/mcp/cmd/mcp-server/app/options"
	"github.com/q8s-io/mcp/pkg/config"
	"github.com/q8s-io/mcp/pkg/db/mysql"
	"github.com/q8s-io/mcp/pkg/server"
)

func main() {
	// load options
	opts := options.NewOptions()
	flag.Parse()
	klog.Infof("server options: %+v", opts)

	// load config
	conf, err := config.LoadConfig(opts.Config)
	if err != nil {
		klog.Errorf("error to load config: %s, %v", opts.Config, err)
		return
	}

	if err := mysql.InitDB(&conf.MysqlConfig); err != nil {
		klog.Errorf("error to init mysql, %v", err)
		return
	}
	defer mysql.CloseDB()

	// start server
	if err := server.StartServer(opts); err != nil {
		klog.Errorf("error to start server, %v", err)
	}
}
