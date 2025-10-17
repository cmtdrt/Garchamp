package recipe

import (
	"api/src/core/utils"
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type createReq struct {
	Allergens    []string `json:"allergens"     validate:"required"`
	Items        []Item   `json:"items" validate:"required"`
	PeopleNumber int      `json:"people_number"    validate:"required"`
}

func (cr *createReq) Bind(_ *http.Request) error {
	v := validator.New()
	if err := v.Struct(cr); err != nil {
		return errors.New("structure de requête invalide")
	}
	return nil
}

type Item struct {
	ID       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
	Unit     string `json:"unit" validate:"required"`
	ExpDate  string `json:"exp_date" validate:"omitempty"`
}

func newItem(id, quantity int, name, unit string, expDate sql.NullString, allergens []string) *Item {
	return &Item{
		ID:       id,
		Quantity: quantity,
		Unit:     unit,
		Name:     name,
		ExpDate:  utils.NullStringToString(expDate),
	}
}
