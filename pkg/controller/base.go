package controller

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

// response writes http.StatusOK to response header and result obj with json format to response.
func response(resp *restful.Response, result interface{}) {
	responseHeaderAndValue(resp, http.StatusOK, result)
}

// responseError writes http status to response header and error message to response.
func responseError(resp *restful.Response, httpStatus int, err error) {
	responseHeaderAndValue(resp, httpStatus, ErrorMessage{
		Message: err.Error(),
	})
}

// responseHeaderAndValue writes http status to response header and result obj with json format to response.
func responseHeaderAndValue(resp *restful.Response, httpStatus int, result interface{}) {
	resp.WriteHeaderAndJson(httpStatus, result, restful.MIME_JSON) //nolint
}
