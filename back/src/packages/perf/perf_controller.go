package perf

import (
	"api/src/core/base"
	"api/src/core/response"
	"api/src/db"
	"net/http"
)

const controllerName = "fridge"

type controller struct {
	service *service
	logger  *base.Logger
}

func newController(repositoryManager *db.RepositoryManager, logger *base.Logger) *controller {
	cLogger := logger.With("controller", controllerName)
	return &controller{
		service: newService(repositoryManager, cLogger),
		logger:  cLogger,
	}
}

func (c *controller) get(w http.ResponseWriter, r *http.Request) {
	items, err := c.service.getPerf()

	if err != nil {
		response.RenderAndLog(r.Context(), w, r, response.ErrServer(err), "fridge-getall", c.logger)
		return
	}

	response.RenderAndLog(
		r.Context(),
		w,
		r,
		response.NewSuccessResponse(items, "Fetched successfully"),
		"fridge-getAll",
		c.logger,
	)
}
