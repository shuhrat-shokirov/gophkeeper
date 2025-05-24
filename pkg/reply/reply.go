package reply

import (
	"encoding/json"
	"net/http"

	"github.com/aliftechuz/pkg/i18n/translation"
	"github.com/aliftechuz/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"go-template/internal/response"
)

var Module = fx.Options(fx.Invoke(New))

type Params struct {
	fx.In
	Logger logger.Logger
}

var params Params

func New(p Params) {
	params = p
}

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json;charset=utf-8"
)

func JSON(ginCtx *gin.Context, response *response.Response) {
	reply, err := json.Marshal(response)
	if err != nil {
		params.Logger.Error(ginCtx, "err on marshal data", zap.Error(err))
		http.Error(ginCtx.Writer, err.Error(), response.HeaderCode)
		return
	}

	ginCtx.Writer.Header().Set(contentType, applicationJSON)
	ginCtx.Writer.WriteHeader(response.HeaderCode)
	_, err = ginCtx.Writer.Write(reply)
	if err != nil {
		params.Logger.Error(ginCtx, "write response error", zap.Error(err))
		return
	}
}

func JSONWithTranslate(ginCtx *gin.Context, response *response.Response, language translation.Language) {
	response.Message = translation.Translate(language, response.Message)

	reply, err := json.Marshal(response)
	if err != nil {
		params.Logger.Error(ginCtx, "err on marshal data", zap.Any("data", response), zap.Error(err))
		JSON(ginCtx, response)
		return
	}

	ginCtx.Writer.Header().Set(contentType, applicationJSON)
	ginCtx.Writer.WriteHeader(response.HeaderCode)
	_, err = ginCtx.Writer.Write(reply)
	if err != nil {
		params.Logger.Error(ginCtx, "write response error", zap.Error(err))
		return
	}
}
