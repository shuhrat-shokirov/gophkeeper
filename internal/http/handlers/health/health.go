package health

import (
	"github.com/aliftechuz/pkg/i18n/translation"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"go-template/internal/response"
	"go-template/pkg/reply"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
}

type Handler interface {
	Health(c *gin.Context)
}

type handler struct {
}

func New(p Params) Handler {
	return &handler{}
}

func (h *handler) Health(c *gin.Context) {
	resp := response.Ok()
	reply.JSONWithTranslate(c, &resp, translation.Ru)
}
