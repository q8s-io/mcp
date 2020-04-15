package kubeconfig

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful/v3"

	"github.com/q8s-io/mcp/pkg/biz/base"
	"github.com/q8s-io/mcp/pkg/db/mysql"
)

type Resource struct {
	base.Resource
	service *Service
}

func NewResource() *Resource {
	return &Resource{
		service: GetService(),
	}
}

// Regist kubeconfig route to webservice.
func (r *Resource) Regist(ws *restful.WebService) {
	ws.Route(ws.GET("/clusters/{id}/kubeconfig").To(r.get).
		// docs
		Doc("Get cluster kubeconfig by cluster id").
		Param(ws.PathParameter("id", "Id of the cluster").DataType("int")).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), ClusterKubeconfigResp{}))
}

// GET http://localhost:8080/api/v1/clusters/{id}/kubeconfig
func (r *Resource) get(req *restful.Request, resp *restful.Response) {
	id, err := strconv.Atoi(req.PathParameter("id"))
	if err != nil {
		r.ResponseError(resp, http.StatusBadRequest, errors.New("param id should be a number"))
		return
	}

	kubeconfigResp, err := r.service.GetByClusterID(mysql.GetDB(), uint(id))
	if err != nil {
		r.ResponseError(resp, http.StatusInternalServerError, err)
		return
	}
	if kubeconfigResp == nil {
		r.ResponseError(resp, http.StatusNotFound, errors.New("model not found"))
		return
	}

	r.Response(resp, kubeconfigResp)
}
