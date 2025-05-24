package i18n

import "embed"

//go:embed messages/messages.*.yaml
var EmbedFs embed.FS
