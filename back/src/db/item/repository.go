package itemdb

import (
	"api/src/core/base"
	"context"
	"database/sql"
	"errors"
)

type Repository struct {
	DBManager *base.DatabaseManager
	Logger    *base.Logger
}

func NewRepository(db *base.DatabaseManager, logger *base.Logger) *Repository {
	return &Repository{DBManager: db, Logger: logger}
}

func (r *Repository) Create(ctx context.Context, name, unit string, quantity int, expDate *string) (sql.Result, error) {
	query := `
	INSERT INTO items (name, unit, quantity, expiration_date) VALUES (?,?,?,?);
	`
	res, err := r.DBManager.DB.ExecContext(ctx, query, name, unit, quantity, expDate)

	if err != nil {
		r.Logger.ErrorContext(ctx, "impossible de créer l'item", "err", err)
		return nil, errors.New("impossible de créer l'item")
	}
	return res, nil
}
