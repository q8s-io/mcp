package controller

import (
	"net/http"
	"github.com/emicklei/go-restful/v3"

	"github.com/q8s-io/mcp/pkg/dto"
	"github.com/q8s-io/mcp/pkg/service"
	"github.com/q8s-io/mcp/pkg/validator"
)

type AzureSecret struct {
	secretService *service.AzureSecret
}

func NewSecretController() *AzureSecret {
	return &AzureSecret{
		secretService: service.NewSecretService(),
	}
}


// Regist secret route to webservice.
func (s *AzureSecret) Regist(ws *restful.WebService) {
	ws.Route(ws.POST("/secrets/azure").To(s.create).
		// docs
		Doc("Attach Secret").
		Reads(dto.SecretAttachReq{}).
		Returns(http.StatusCreated, http.StatusText(http.StatusOK), dto.SecretAttachResp{}))
}

// POST http://localhost:8080/api/v1/secrets/azure
func (s *AzureSecret) create(req *restful.Request, resp *restful.Response) {
	attachSecret := new(dto.SecretAttachReq)
	if err := req.ReadEntity(attachSecret); err != nil {
		responseError(resp, http.StatusBadRequest, err)
		return
	}
	// validate value
	if err := validator.GetValidate().Struct(attachSecret); err != nil {
		responseError(resp, http.StatusBadRequest, err)
		return
	}
	attchResp, err := s.secretService.SecretCreate(attachSecret)
	if err != nil {
		responseError(resp, http.StatusBadRequest, err)
		return
	}

	responseHeaderAndValue(resp, http.StatusCreated, attchResp)
}

