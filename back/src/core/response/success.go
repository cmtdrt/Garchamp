package response

import (
	"net/http"

	"github.com/go-chi/render"
)

// SuccessResponse est la structure de la réponse pour les succès.
type SuccessResponse struct {
	HTTPCode int         `json:"-"`
	Status   string      `json:"status"`            // Statut de la réponse (success)
	Message  string      `json:"message,omitempty"` // Message descriptif
	Data     interface{} `json:"data,omitempty"`    // Données retournées (si applicable)
}

func NewSuccessResponse(data interface{}, message string) *SuccessResponse {
	return &SuccessResponse{
		HTTPCode: http.StatusOK,
		Status:   "success",
		Message:  message,
		Data:     data,
	}
}

func NewSuccessCreatedResponse(data interface{}, message string) *SuccessResponse {
	return &SuccessResponse{
		HTTPCode: http.StatusCreated,
		Status:   "success",
		Message:  message,
		Data:     data,
	}
}

func NewNoContentRes() *SuccessResponse {
	return &SuccessResponse{
		HTTPCode: http.StatusNoContent,
		Status:   "success",
	}
}

func (sr *SuccessResponse) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, sr.HTTPCode)
	return nil
}
