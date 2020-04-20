package base

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

type Resource struct{}

// Response writes http.StatusOK to response header and result obj with json format to response.
func (r *Resource) Response(resp *restful.Response, result interface{}) {
	r.ResponseHeaderAndValue(resp, http.StatusOK, result)
}

// ResponseHeaderAndValue writes http status to response header and result obj with json format to response.
func (r *Resource) ResponseHeaderAndValue(resp *restful.Response, httpStatus int, result interface{}) {
	resp.WriteHeaderAndJson(httpStatus, result, restful.MIME_JSON) //nolint
}

type ErrorMessage struct {
	Message string `json:"message"`
}

// ResponseError writes http status to response header and error message to response.
func (r *Resource) ResponseError(resp *restful.Response, httpStatus int, err error) {
	r.ResponseHeaderAndValue(resp, httpStatus, ErrorMessage{
		Message: err.Error(),
	})
}
