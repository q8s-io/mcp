package controller

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"

	"github.com/q8s-io/mcp/pkg/dto"
	"github.com/q8s-io/mcp/pkg/service"
	"github.com/q8s-io/mcp/pkg/validator"
)

type Azure struct {
	azureService *service.Azure
}

func NewAzureController() *Azure {
	return &Azure{
		azureService: service.NewAzureService(),
	}
}

// Regist cluster route to webservice.
func (a *Azure) Regist(ws *restful.WebService) {
	ws.Route(ws.GET("/azure/subscriptions").To(a.subscriptions).
		// docs
		Doc("Get subscriptions").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), dto.AzureSubscriptionsResp{}))
}

// GET http://localhost:8080/api/v1/azure/subscriptions
func (a *Azure) subscriptions(req *restful.Request, resp *restful.Response) {
	subscriptionsReq := new(dto.AzureSubscriptionsReq)
	if err := req.ReadEntity(subscriptionsReq); err != nil {
		responseError(resp, http.StatusBadRequest, err)
		return
	}

	// validate value
	if err := validator.GetValidate().Struct(subscriptionsReq); err != nil {
		responseError(resp, http.StatusBadRequest, err)
		return
	}

	subscriptions, err := a.azureService.Subscriptions(subscriptionsReq)
	if err != nil {
		responseError(resp, http.StatusBadRequest, err)
		return
	}

	response(resp, subscriptions)
}
