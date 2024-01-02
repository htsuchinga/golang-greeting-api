package greeting

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/htsuchinga/golang-greeting-api/api/internal/apierror"
	"github.com/htsuchinga/golang-greeting-api/api/internal/greeting/hello"
	"github.com/htsuchinga/golang-greeting-api/api/usecase/greeting"

	"github.com/htsuchinga/golang-greeting-api/internal/datetime"
	"github.com/htsuchinga/golang-greeting-api/internal/logger"
)

type v1GreetingHandler struct {
	validate *validator.Validate
}

func NewV1GreetingHandler(validate *validator.Validate) http.Handler {
	return &v1GreetingHandler{
		validate: validate,
	}
}

func (h *v1GreetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	moduleName := "v1GreetingHello"
	log := logger.ModuleName(moduleName)

	req, err := hello.ParseRequest(r, h.validate)
	if err != nil {
		log.Warn("Request parse error: %s", err)
		switch err.(type) {
		case *apierror.HTTPMethodNotAllowed:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		hello.CreateValidationErrorResponse().Write(w)
		return
	}

	received := datetime.NowInJST()

	useCase := greeting.NewHelloUseCase(received, req)
	response := useCase.Exec()
	response.Write(w)

}
