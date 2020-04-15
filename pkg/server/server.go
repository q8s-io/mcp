package server

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"k8s.io/klog"

	"github.com/q8s-io/mcp/cmd/mcp-server/app/options"
	"github.com/q8s-io/mcp/pkg/filters"
	"github.com/q8s-io/mcp/pkg/routers"
)

func NewContainer() *restful.Container {
	container := restful.NewContainer()

	ws := new(restful.WebService)
	// import routers to webservice
	routers.AddRouters(ws)

	container.Add(ws)

	return container
}

// StartServer starts web server for mcp.
func StartServer(opts *options.Options) error {
	container := NewContainer()

	// add filters to container
	filters.AddFilters(container)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", opts.Port),
		Handler: container,
	}

	klog.Infof("start server, addr: %s", server.Addr)
	return server.ListenAndServe()
}
