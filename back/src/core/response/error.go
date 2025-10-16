package response

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

// ErrResponse est la structure de la réponse pour les erreurs.
type ErrResponse struct {
	Err            error `json:"-"` // Erreur détaillée (en interne, pas envoyée)
	HTTPStatusCode int   `json:"-"` // Code HTTP de l'erreur

	StatusText   string      `json:"status"`          // Statut de la réponse (error)
	AppCode      int64       `json:"code,omitempty"`  // Code d'erreur de l'application
	ErrorText    string      `json:"error,omitempty"` // Message d'erreur
	ErrorDetails interface{} `json:"errorDetails,omitempty"`
}

func newErrorWithContent(message string, code int, details interface{}) *ErrResponse {
	return &ErrResponse{
		StatusText:     "error",
		ErrorText:      message,
		HTTPStatusCode: code,
		ErrorDetails:   details,
	}
}

func ErrCreate(message string, content interface{}) render.Renderer {
	return newErrorWithContent(message, http.StatusUnauthorized, content)
}

func newErrorResponse(message string, code int) *ErrResponse {
	return &ErrResponse{
		StatusText:     "error",
		ErrorText:      message,
		HTTPStatusCode: code,
	}
}

func ErrBadRequest(err error) render.Renderer {
	return newErrorResponse(err.Error(), http.StatusBadRequest)
}

func ErrUnauthorized(err error) render.Renderer {
	return newErrorResponse(err.Error(), http.StatusUnauthorized)
}

func ErrNotFound(err error) render.Renderer {
	return newErrorResponse(err.Error(), http.StatusNotFound)
}

func ErrServer(err error) render.Renderer {
	return newErrorResponse(err.Error(), http.StatusInternalServerError)
}

func FieldsError(possibleFields []string) render.Renderer {
	fields := strings.Join(possibleFields, ",")
	errorMsg := fmt.Sprintf(
		"Un ou des champs que vous avez sélectionné n'existe pas, les champs possibles sont : %v",
		fields,
	)
	return newErrorResponse(errorMsg, http.StatusNotFound)
}

func (e *ErrResponse) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
