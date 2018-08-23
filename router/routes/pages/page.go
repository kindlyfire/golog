package pages

import (
	"strings"

	"github.com/kindlyfire/golog/modules/context"
)

// Page handlers the rendering of a page
func Page(ctx *context.Context) string {
	// Filter out any URL that starts with /_/, because we do not serve that
	// Eg, when a file named /_/gl-assets/css/main.css is not found, this is called anyway
	// We filter that out
	if strings.HasPrefix(ctx.RealCtx.Req.URL.Path, "/_/") {
		ctx.Status(404)
		return ""
	}

	return "Page to render here !"
}
