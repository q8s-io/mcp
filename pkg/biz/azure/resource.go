package azure

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"

	"github.com/q8s-io/mcp/pkg/biz/base"
	"github.com/q8s-io/mcp/pkg/validator"
)

type Resource struct {
	base.Resource
	azureService *Service
}

func NewResource() *Resource {
	return &Resource{
		azureService: GetService(),
	}
}

// Regist cluster route to webservice.
func (r *Resource) Regist(ws *restful.WebService) {
	ws.Route(ws.GET("/azure/subscriptions").To(r.subscriptions).
		// docs
		Doc("Get subscriptions").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), SubscriptionsListResp{}))
}

// GET http://localhost:8080/api/v1/azure/subscriptions
func (r *Resource) subscriptions(req *restful.Request, resp *restful.Response) {
	subscriptionsReq := new(SubscriptionsReq)
	if err := req.ReadEntity(subscriptionsReq); err != nil {
		r.ResponseError(resp, http.StatusBadRequest, err)
		return
	}

	// validate value
	if err := validator.GetValidate().Struct(subscriptionsReq); err != nil {
		r.ResponseError(resp, http.StatusBadRequest, err)
		return
	}

	subscriptions, err := r.azureService.subscriptions(subscriptionsReq)
	if err != nil {
		r.ResponseError(resp, http.StatusBadRequest, err)
		return
	}

	r.Response(resp, subscriptions)
}
