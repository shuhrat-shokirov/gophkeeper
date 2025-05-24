package exceptions

import (
	"github.com/aliftechuz/pkg/errorsx"

	"go-template/internal/response/codes"
)

var (
	ErrNotFound   = errorsx.New().WithCode(codes.NotFoundCode).WithMessage("NotFound")
	ErrBadRequest = errorsx.New().WithCode(codes.BadRequestCode).WithMessage("BadRequest")
	ErrInternal   = errorsx.New().
			WithCode(codes.InternalErrCode).
			WithMessage("InternalError").
			WithInternal()
	ErrUnauthorized = errorsx.New().
			WithCode(codes.UnauthorizedCode).
			WithMessage("AuthDataWrong").
			WithStatusCode(codes.UnauthorizedCode)
	ErrTimeout = errorsx.New().
			WithCode(codes.InternalErrCode).
			WithMessage("Timeout").
			WithInternal()
)
