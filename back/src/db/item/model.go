package itemdb

import (
	"api/src/core/utils"
	"database/sql"
)

type Model struct {
	ID           int
	Name         string
	Quantity     int
	Unit         string
	ExpDate      sql.NullString
	Kcal         int
	Protein      int
	Fat          int
	Carbohydrate int
	Fiber        int
	Sugar        int
	Salt         int
}

func NewModel(id int, name, unit string, quantity, Kcal,
	protein,
	fat,
	carbohydrate,
	fiber,
	sugar,
	salt int, expDate *string) *Model {
	return &Model{
		ID:           id,
		Name:         name,
		Unit:         unit,
		Quantity:     quantity,
		Protein:      protein,
		Fat:          fat,
		Carbohydrate: carbohydrate,
		Fiber:        fiber,
		Sugar:        sugar,
		Salt:         salt,
		ExpDate:      utils.NullStringValidation(expDate),
	}
}
