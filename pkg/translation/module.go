package translation

import (
	"github.com/aliftechuz/pkg/i18n/translation"
	"go.uber.org/fx"

	"go-template/i18n"
)

var Module = fx.Invoke(func() {
	translation.New(i18n.EmbedFs)
})
