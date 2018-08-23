package admin

import (
	"github.com/kindlyfire/golog/modules/context"
)

func Index(ctx *context.Context) {
	ctx.AdminHTML(200, "panel/index")
}
