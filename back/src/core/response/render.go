package response

import (
	"api/src/core/base"
	"context"
	"net/http"

	"github.com/go-chi/render"
)

func RenderAndLog(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
	errResp render.Renderer,
	renderPosition string,
	logger *base.Logger,
) {
	if err := render.Render(w, r, errResp); err != nil {
		logger.LogError(ctx, "render", err, "renderPosition", renderPosition)
	}
}
