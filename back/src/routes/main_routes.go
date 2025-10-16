package routes

import (
	"net/http"
	"time"

	// Importing generated Swagger docs for side effects only.
	_ "api/docs" // revive:disable-line:blank-imports
	"api/src/core/base"
	"api/src/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter(
	corsOptions *cors.Options,
	repositoryManager *db.RepositoryManager,
	logger *base.Logger,
) chi.Router {
	const maxRequestLimit = 100
	const maxReqDuration = 120 * time.Second
	r := chi.NewRouter()
	r.Use(cors.Handler(*corsOptions)) // Gestion des CORS policy
	r.Use(
		middleware.RequestID,
	) // Ajoute un identifiant unique à chaque demande pour une meilleure traçabilité des logs
	r.Use(
		middleware.Logger,
	) // Gère la reprise sur panique pour éviter que le serveur ne se bloque en cas d'erreur.
	r.Use(
		middleware.Recoverer,
	) // Permet de ne pas crash apres un pannic
	r.Use(middleware.URLFormat) // Normalise les URL pour améliorer le routage
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(maxReqDuration))

	// Décommenter cette ligne pour afficher la documention OpenAPI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// r.Handle("/ws", http.HandlerFunc(ws.WSHandleConnections))
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
			if _, err := w.Write([]byte("pong")); err != nil {
				logger.Error("Erreur lors de l'écriture de la réponse", "erreur", err)
			}
		})
	})

	return r
}
