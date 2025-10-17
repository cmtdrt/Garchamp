package perf

import (
	"api/src/core/base"
	"api/src/db"

	"github.com/go-chi/chi/v5"
)

func Route(repositoryManager *db.RepositoryManager, logger *base.Logger) chi.Router {
	r := chi.NewRouter()
	c := newController(repositoryManager, logger)
	r.Get("/", c.get)
	return r
}
