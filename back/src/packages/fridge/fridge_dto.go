package fridge

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type createReq struct {
	Name     string `json:"name"     validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
	Unity    string `json:"unity"    validate:"required,oneof=g kg ml L"`
	ExpDate  string `json:"exp_date" validate:"omitempty"`
}

func (cr *createReq) Bind(_ *http.Request) error {
	v := validator.New()
	if err := v.Struct(cr); err != nil {
		return errors.New("structure de requête invalide")
	}
	return nil
}
