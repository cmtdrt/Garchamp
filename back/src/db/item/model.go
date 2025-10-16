package itemdb

import (
	"api/src/core/utils"
	"database/sql"
)

type Model struct {
	Name     string
	Quantity int
	Unit     string
	ExpDate  sql.NullString
}

func NewModel(name, unit string, quantity int, expDate *string) *Model {
	return &Model{
		Name:     name,
		Unit:     unit,
		Quantity: quantity,
		ExpDate:  utils.NullStringValidation(expDate),
	}
}
