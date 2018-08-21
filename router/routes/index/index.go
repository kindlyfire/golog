package index

import "github.com/kindlyfire/golog/modules/context"

// Index ...
func Index(ctx context.Context) {
	ctx.Render(200, "other")
}
