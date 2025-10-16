package itemallergenrelationdb

import (
	"api/src/core/base"
	"context"
	"database/sql"
)

type Repository struct {
	DBManager *base.DatabaseManager
	Logger    *base.Logger
}

func NewRepository(db *base.DatabaseManager, logger *base.Logger) *Repository {
	return &Repository{DBManager: db, Logger: logger}
}

func (r *Repository) Create(ctx context.Context, idItem, idAllergen int64) (sql.Result, error) {
	query := "INSERT INTO items_allergens_relation (item_ID, allergen_ID) VALUES (?,?);"
	return r.DBManager.DB.ExecContext(ctx, query, idItem, idAllergen)
}
