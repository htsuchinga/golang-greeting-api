package hello

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/htsuchinga/golang-greeting-api/api/internal/apierror"
)

type Request struct {
	Name string `validate:"required"`
}

func ParseRequest(r *http.Request, v *validator.Validate) (*Request, error) {
	if r.Method != http.MethodPost {
		return nil, apierror.NewHTTPMethodNotAllowed([]string{http.MethodPost})
	}

	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	req := &Request{
		Name: r.Form.Get("name"),
	}

	if err := v.Struct(req); err != nil {
		return nil, fmt.Errorf("validation error: %s", err)
	}

	return req, nil
}
