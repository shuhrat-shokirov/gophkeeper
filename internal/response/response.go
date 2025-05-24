package response

import (
	"net/http"

	"github.com/aliftechuz/pkg/errorsx"

	"go-template/internal/response/codes"
)

type Response struct {
	Payload    any    `json:"payload"`
	Message    string `json:"message" example:"success"`
	HeaderCode int    `json:"-"`
	Code       int    `json:"code" example:"200"`
}

func Ok() Response {
	return Response{
		Message:    "Success",
		HeaderCode: http.StatusOK,
		Code:       codes.OkCode,
	}
}

func Err(err error) Response {
	var result = newResponse(codes.OkCode, "Success")

	if err == nil {
		return result
	}

	if errX, ok := errorsx.As(err); ok {
		result = newResponse(errX.Code(), errX.Message())
		result.HeaderCode = errX.StatusCode()
	} else {
		result = newResponse(codes.InternalErrCode, "InternalErr")
	}

	return result
}

func newResponse(code int, message string) Response {
	return Response{
		HeaderCode: codes.OkCode,
		Code:       code,
		Message:    message,
	}
}

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json;charset=utf-8"
)

func (r *Response) WithPayload(payload any) *Response {
	r.Payload = payload

	return r
}
