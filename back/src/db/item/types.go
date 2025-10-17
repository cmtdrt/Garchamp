package itemdb

import (
	"api/src/core/utils"
	"database/sql"
)

type ItemLite struct {
	ID       int
	Name     string
	Quantity int
	Unit     string
	ExpDate  sql.NullString
}

func NewItemLite(id int, name, unit string, quantity int, expDate *string) *ItemLite {
	return &ItemLite{
		ID:       id,
		Name:     name,
		Unit:     unit,
		Quantity: quantity,
		ExpDate:  utils.NullStringValidation(expDate),
	}
}
