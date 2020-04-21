package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful/v3"

	"github.com/q8s-io/mcp/pkg/dto"
	"github.com/q8s-io/mcp/pkg/service"
)

type Kubeconfig struct {
	kubeconfigService *service.Kubeconfig
}

func NewKubeconfigController() *Kubeconfig {
	return &Kubeconfig{
		kubeconfigService: service.NewKubeconfigService(),
	}
}

// Regist kubeconfig route to webservice.
func (k *Kubeconfig) Regist(ws *restful.WebService) {
	ws.Route(ws.GET("/clusters/{id}/kubeconfig").To(k.get).
		// docs
		Doc("Get cluster kubeconfig by cluster id").
		Param(ws.PathParameter("id", "Id of the cluster").DataType("int")).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), dto.KubeconfigResp{}))
}

// GET http://localhost:8080/api/v1/clusters/{id}/kubeconfig
func (k *Kubeconfig) get(req *restful.Request, resp *restful.Response) {
	id, err := strconv.Atoi(req.PathParameter("id"))
	if err != nil {
		responseError(resp, http.StatusBadRequest, errors.New("param id should be a number"))
		return
	}

	kubeconfigResp, err := k.kubeconfigService.GetByClusterID(uint(id))
	if err != nil {
		responseError(resp, http.StatusInternalServerError, err)
		return
	}
	if kubeconfigResp == nil {
		responseError(resp, http.StatusNotFound, errors.New("model not found"))
		return
	}

	response(resp, kubeconfigResp)
}
