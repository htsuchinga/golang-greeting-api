package apierror

import (
	"fmt"
	"strings"
)

type (
	HTTPMethodNotAllowed struct {
		errorMsg     string
		AllowMethods []string
	}
)

func NewHTTPMethodNotAllowed(allowMethods []string) *HTTPMethodNotAllowed {
	return &HTTPMethodNotAllowed{
		errorMsg: fmt.Sprintf("%s以外の許可されていないHTTPメソッド", strings.Join(allowMethods, ", ")),
	}
}

func (e *HTTPMethodNotAllowed) Error() string {
	return e.errorMsg
}
