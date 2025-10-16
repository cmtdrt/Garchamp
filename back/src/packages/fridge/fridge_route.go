package fridge

import (
	"api/src/core/base"
	"api/src/db"

	"github.com/go-chi/chi/v5"
)

func Route(repositoryManager *db.RepositoryManager, logger *base.Logger) chi.Router {
	r := chi.NewRouter()
	groupController := newController(repositoryManager, logger)
	r.Get("/", groupController.getAll)
	// r.Get("/{itemID}", groupController.get)
	r.Delete("/{itemID}", groupController.delete)
	r.Post("/", groupController.create)
	return r
}
