package greeting

import (
	"time"

	"github.com/htsuchinga/golang-greeting-api/api/internal/greeting/hello"
	"github.com/htsuchinga/golang-greeting-api/internal/datetime"
)

type (
	HelloUseCase interface {
		Exec() *hello.Response
	}

	helloUseCase struct {
		received time.Time
		req      *hello.Request
	}
)

func NewHelloUseCase(received time.Time, req *hello.Request) HelloUseCase {
	return &helloUseCase{
		received: received,
		req:      req,
	}
}

func (u *helloUseCase) Exec() *hello.Response {
	message := "Hello, " + u.req.Name + "!"
	receiveDate := datetime.NowInJST().Format("20060102")
	receiveTime := datetime.NowInJST().Format("150405")

	return hello.CreateSuccessResponse(message, receiveDate, receiveTime)
}
