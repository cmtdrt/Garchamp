package allergendb

import (
	"api/src/core/base"
	"context"
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
