package controller

import (
	"net/http"
	"github.com/emicklei/go-restful/v3"

	"github.com/q8s-io/mcp/pkg/dto"
	"github.com/q8s-io/mcp/pkg/service"
	"github.com/q8s-io/mcp/pkg/validator"
)

type Secret_Azure struct {
	secretService *service.Secret_Azure
}

func NewSecretController() *Secret_Azure {
	return &Secret_Azure{
		secretService: service.NewSecretService(),
	}
}


// Regist secret route to webservice.
func (s *Secret_Azure) Regist(ws *restful.WebService) {
	ws.Route(ws.POST("/secret").To(s.attach).
		// docs
		Doc("Attach Secret").
		Reads(dto.SecretAttachReq{}).
		Returns(http.StatusCreated, http.StatusText(http.StatusOK), dto.SecretAttachResp{}))
}

// POST http://localhost:8080/api/v1/azure/attach
func (s *Secret_Azure) attach(req *restful.Request, resp *restful.Response) {
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
	attchResp, err := s.secretService.Attach(attachSecret)
	if err != nil {
		responseError(resp, http.StatusBadRequest, err)
		return
	}

	responseHeaderAndValue(resp, http.StatusCreated, attchResp)
}

