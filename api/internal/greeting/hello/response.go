package hello

import (
	"encoding/json"
	"net/http"

	"github.com/htsuchinga/golang-greeting-api/api/internal"
)

type (
	ResponseContents struct {
		Message     string `json:"message"`
		ReceiveDate string `json:"receiveDate"`
		ReceiveTime string `json:"receiveTime"`
	}

	Response struct {
		ResultCd internal.ResultCd `json:"resultCd"`
		Contents *ResponseContents `json:"contents,omitempty"`
	}
)

func (r *Response) Write(w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	buf, _ := json.MarshalIndent(r, "", "  ")
	_, _ = w.Write(buf)
}

func createResponse(resultCd internal.ResultCd, contents *ResponseContents) *Response {
	return &Response{
		ResultCd: resultCd,
		Contents: contents,
	}
}

func CreateSuccessResponse(message string, receiveDate string, receiveTime string) *Response {
	return createResponse(internal.ResultCdSuccess, &ResponseContents{
		Message:     message,
		ReceiveDate: receiveDate,
		ReceiveTime: receiveTime,
	})
}

func CreateValidationErrorResponse() *Response {
	return createResponse(internal.ResultCdValidationError, nil)
}

func CreateSystemErrorResponse() *Response {
	return createResponse(internal.ResultCdSystemError, nil)
}
