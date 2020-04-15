package cluster

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful/v3"

	"github.com/q8s-io/mcp/pkg/biz/base"
	"github.com/q8s-io/mcp/pkg/db/mysql"
	"github.com/q8s-io/mcp/pkg/validator"
)

type Resource struct {
	base.Resource
	clusterService *Service
}

func NewResource() *Resource {
	return &Resource{
		clusterService: GetService(),
	}
}

// Regist cluster route to webservice.
func (r *Resource) Regist(ws *restful.WebService) {
	ws.Route(ws.GET("/clusters").To(r.all).
		// docs
		Doc("Get clusters").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []ListResp{}))

	ws.Route(ws.GET("/clusters/{id}").To(r.get).
		// docs
		Doc("Get cluster by id").
		Param(ws.PathParameter("id", "Id of the cluster").DataType("int")).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), DetailResp{}))

	ws.Route(ws.POST("/clusters/attach").To(r.attach).
		// docs
		Doc("Attach cluster").
		Param(ws.PathParameter("id", "Id of the cluster").DataType("int")).
		Reads(AttachReq{}).
		Returns(http.StatusCreated, http.StatusText(http.StatusOK), AttachResp{}))
}

// GET http://localhost:8080/api/v1/clusters
func (r *Resource) all(req *restful.Request, resp *restful.Response) {
	clusterResps, err := r.clusterService.All(mysql.GetDB())
	if err != nil {
		r.ResponseError(resp, http.StatusInternalServerError, err)
		return
	}

	r.Response(resp, clusterResps)
}

// GET http://localhost:8080/api/v1/clusters/{id}
func (r *Resource) get(req *restful.Request, resp *restful.Response) {
	id, err := strconv.Atoi(req.PathParameter("id"))
	if err != nil {
		r.ResponseError(resp, http.StatusBadRequest, errors.New("param id should be a number"))
		return
	}

	clusterResp, err := r.clusterService.GetByID(mysql.GetDB(), uint(id))
	if err != nil {
		r.ResponseError(resp, http.StatusInternalServerError, err)
		return
	}
	if clusterResp == nil {
		r.ResponseError(resp, http.StatusNotFound, errors.New("model not found"))
		return
	}

	r.Response(resp, clusterResp)
}

// POST http://localhost:8080/api/v1/clusters/attach
func (r *Resource) attach(req *restful.Request, resp *restful.Response) {
	attachCluster := new(AttachReq)
	if err := req.ReadEntity(attachCluster); err != nil {
		r.ResponseError(resp, http.StatusBadRequest, err)
		return
	}

	// validate value
	if err := validator.GetValidate().Struct(attachCluster); err != nil {
		r.ResponseError(resp, http.StatusBadRequest, err)
		return
	}

	attchResp, err := r.clusterService.Attach(mysql.GetDB(), attachCluster)
	if err != nil {
		r.ResponseError(resp, http.StatusBadRequest, err)
		return
	}

	r.ResponseHeaderAndValue(resp, http.StatusCreated, attchResp)
}
