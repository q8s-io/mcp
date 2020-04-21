package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful/v3"

	"github.com/q8s-io/mcp/pkg/dto"
	"github.com/q8s-io/mcp/pkg/service"
	"github.com/q8s-io/mcp/pkg/validator"
)

type Cluster struct {
	clusterService *service.Cluster
}

func NewClusterController() *Cluster {
	return &Cluster{
		clusterService: service.NewClusterService(),
	}
}

// Regist cluster route to webservice.
func (c *Cluster) Regist(ws *restful.WebService) {
	ws.Route(ws.GET("/clusters").To(c.all).
		// docs
		Doc("Get clusters").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []dto.ClusterListResp{}))

	ws.Route(ws.GET("/clusters/{id}").To(c.get).
		// docs
		Doc("Get cluster by id").
		Param(ws.PathParameter("id", "Id of the cluster").DataType("int")).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), dto.ClusterDetailResp{}))

	ws.Route(ws.POST("/clusters/attach").To(c.attach).
		// docs
		Doc("Attach cluster").
		Param(ws.PathParameter("id", "Id of the cluster").DataType("int")).
		Reads(dto.ClusterAttachReq{}).
		Returns(http.StatusCreated, http.StatusText(http.StatusOK), dto.ClusterAttachResp{}))
}

// GET http://localhost:8080/api/v1/clusters
func (c *Cluster) all(req *restful.Request, resp *restful.Response) {
	clusterResps, err := c.clusterService.All()
	if err != nil {
		responseError(resp, http.StatusInternalServerError, err)
		return
	}

	response(resp, clusterResps)
}

// GET http://localhost:8080/api/v1/clusters/{id}
func (c *Cluster) get(req *restful.Request, resp *restful.Response) {
	id, err := strconv.Atoi(req.PathParameter("id"))
	if err != nil {
		responseError(resp, http.StatusBadRequest, errors.New("param id should be a number"))
		return
	}

	clusterResp, err := c.clusterService.GetByID(uint(id))
	if err != nil {
		responseError(resp, http.StatusInternalServerError, err)
		return
	}
	if clusterResp == nil {
		responseError(resp, http.StatusNotFound, errors.New("model not found"))
		return
	}

	response(resp, clusterResp)
}

// POST http://localhost:8080/api/v1/clusters/attach
func (c *Cluster) attach(req *restful.Request, resp *restful.Response) {
	attachCluster := new(dto.ClusterAttachReq)
	if err := req.ReadEntity(attachCluster); err != nil {
		responseError(resp, http.StatusBadRequest, err)
		return
	}

	// validate value
	if err := validator.GetValidate().Struct(attachCluster); err != nil {
		responseError(resp, http.StatusBadRequest, err)
		return
	}

	attchResp, err := c.clusterService.Attach(attachCluster)
	if err != nil {
		responseError(resp, http.StatusBadRequest, err)
		return
	}

	responseHeaderAndValue(resp, http.StatusCreated, attchResp)
}
