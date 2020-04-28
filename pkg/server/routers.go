package server

import (
	"github.com/emicklei/go-restful/v3"

	"github.com/q8s-io/mcp/pkg/controller"
)

func addRouters(ws *restful.WebService) {
	ws.Path("/api/v1")
	// just support json
	ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	controller.NewClusterController().Regist(ws)
	controller.NewKubeconfigController().Regist(ws)
	controller.NewAzureController().Regist(ws)
	controller.NewSecretController().Regist(ws)
}
