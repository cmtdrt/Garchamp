package recipe

import (
	"api/src/core/base"
	"api/src/core/response"
	"api/src/db"
	"errors"
	"net/http"

	"github.com/go-chi/render"
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

func (c *controller) create(w http.ResponseWriter, r *http.Request) {
	const controllerReference = "fridge-create"
	var err error
	req := &createReq{}

	if err = render.Bind(r, req); err != nil {
		response.RenderAndLog(r.Context(),
			w,
			r,
			response.ErrBadRequest(errors.New("structure de requête invalide")),
			controllerReference,
			c.logger,
		)
		return
	}
	var rsp *recipeAIResponse
	rsp, err = c.service.create(r.Context(), req)
	if err == nil {
		response.RenderAndLog(r.Context(),
			w,
			r,
			response.NewSuccessCreatedResponse(rsp, "item created successfully"),
			controllerReference,
			c.logger,
		)
		return
	}

	response.RenderAndLog(r.Context(),
		w,
		r,
		response.ErrServer(base.ErrServerText),
		controllerReference,
		c.logger,
	)
}
