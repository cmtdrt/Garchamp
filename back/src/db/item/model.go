package itemdb

import (
	"api/src/core/utils"
	"database/sql"
)

type Model struct {
	ID       int
	Name     string
	Quantity int
	Unit     string
	ExpDate  sql.NullString
}

func NewModel(id int, name, unit string, quantity int, expDate *string) *Model {
	return &Model{
		ID:       id,
		Name:     name,
		Unit:     unit,
		Quantity: quantity,
		ExpDate:  utils.NullStringValidation(expDate),
	}
}
