package fridge

import (
	"api/src/core/utils"
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type createReq struct {
	Name     string `json:"name"     validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
	Unity    string `json:"unity"    validate:"required,oneof=g kg ml L unité"`
	ExpDate  string `json:"exp_date" validate:"omitempty"`
}

func (cr *createReq) Bind(_ *http.Request) error {
	v := validator.New()
	if err := v.Struct(cr); err != nil {
		return errors.New("structure de requête invalide")
	}
	return nil
}

type Item struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Quantity  int      `json:"quantity"`
	Unit      string   `json:"unit"`
	ExpDate   string   `json:"exp_date"`
	Allergens []string `json:"allergens"`
}

func newItem(id, quantity int, name, unit string, expDate sql.NullString, allergens []string) *Item {
	return &Item{
		ID:        id,
		Quantity:  quantity,
		Unit:      unit,
		Name:      name,
		ExpDate:   utils.NullStringToString(expDate),
		Allergens: allergens,
	}
}

type itemDetails struct {
	Kcal         int `json:"kcal"`
	Protein      int `json:"protein"`
	Fat          int `json:"fat"`
	Carbohydrate int `json:"carbohydrate"`
	Fiber        int `json:"siber"`
	Sugar        int `json:"sugar"`
	Salt         int `json:"salt"`
}

func newItemDetails(kcal,
	protein,
	fat,
	carbohydrate,
	fiber,
	sugar,
	salt int) *itemDetails {
	return &itemDetails{
		Kcal:         kcal,
		Fat:          fat,
		Carbohydrate: carbohydrate,
		Fiber:        fiber,
		Sugar:        sugar,
		Salt:         salt,
	}
}
