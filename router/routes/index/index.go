package index

import (
	"gopkg.in/macaron.v1"
)

// Index ...
func Index(ctx *macaron.Context) {
	ctx.HTML(200, "index")
}
