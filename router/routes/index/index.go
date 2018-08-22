package index

import "github.com/kindlyfire/golog/modules/context"

// Index ...
func Index(ctx context.Context) {
	ctx.ThemeHTML(200, "index")
}
