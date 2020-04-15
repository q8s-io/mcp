package routers

import (
	"github.com/emicklei/go-restful/v3"

	"github.com/q8s-io/mcp/pkg/biz/cluster"
	"github.com/q8s-io/mcp/pkg/biz/kubeconfig"
)

func AddRouters(ws *restful.WebService) {
	ws.Path("/api/v1")
	// just support json
	ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	cluster.NewResource().Regist(ws)
	kubeconfig.NewResource().Regist(ws)
}
