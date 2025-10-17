package allergendb

import (
	"api/src/core/base"
	"context"
	"errors"
)

type Repository struct {
	DBManager *base.DatabaseManager
	Logger    *base.Logger
}

func NewRepository(db *base.DatabaseManager, logger *base.Logger) *Repository {
	return &Repository{DBManager: db, Logger: logger}
}

func (r *Repository) FindByName(ctx context.Context, name string) int {
	query := `SELECT id FROM allergens WHERE name = ?`
	rslt, err := r.DBManager.DB.QueryContext(ctx, query, name)

	if err != nil {
		return -1
	}
	defer rslt.Close()

	var id = -1
	for rslt.Next() {
		err = rslt.Scan(&id)
		if err != nil {
			return -1
		}
	}
	return id
}

func (r *Repository) GetAllAllergensByRelation(ctx context.Context, itemID int) ([]string, error) {
	query := `SELECT name FROM allergens a INNER JOIN items_allergens_relation iar ON a.id = iar.allergen_ID  WHERE item_ID = ?;`
	rslt, err := r.DBManager.DB.QueryContext(ctx, query, itemID)

	if err != nil {
		return nil, errors.New("error")
	}
	defer rslt.Close()

	var (
		names = []string{}
		name  string
	)
	for rslt.Next() {
		err = rslt.Scan(&name)
		if err != nil {
			return nil, errors.New("error")
		}
		names = append(names, name)
	}
	return names, nil
}
