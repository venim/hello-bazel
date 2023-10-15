package web

import (
	"embed"
)

//go:embed *
var Static embed.FS
