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

func (r *Repository) Create(
	ctx context.Context,
	name, unit string,
	quantity int,
	kcal, protein, fat, carbohydrate, fiber, sugar, salt float64,
	expDate *string,
) (sql.Result, error) {
	query := `
	INSERT INTO items (name, unit, quantity, expiration_date, energy_kcal, protein_g, fat_g, carbohydrate_g, fiber_g, sugar_g, salt_g ) VALUES (?,?,?,?,?,?,?,?,?,?,?);;
	`
	res, err := r.DBManager.DB.ExecContext(
		ctx,
		query,
		name,
		unit,
		quantity,
		expDate,
		kcal,
		protein,
		fat,
		carbohydrate,
		fiber,
		sugar,
		salt,
	)

	if err != nil {
		r.Logger.ErrorContext(ctx, "impossible de créer l'item", "err", err)
		return nil, errors.New("impossible de créer l'item")
	}
	return res, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]ItemLite, error) {
	query := `SELECT id, name, quantity, unit, expiration_date FROM items i `
	rslt, err := r.DBManager.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, errors.New("error")
	}
	defer rslt.Close()

	var (
		items        = []ItemLite{}
		id, quantity int
		name, unit   string
		expDate      *string
	)
	for rslt.Next() {
		err = rslt.Scan(&id, &name, &quantity, &unit, &expDate)
		if err != nil {
			return nil, errors.New("error")
		}
		items = append(items, *NewItemLite(id, name, unit, quantity, expDate))
	}
	if rslt.Err() != nil {
		return nil, errors.New("erreur")
	}
	return items, nil
}

func (r *Repository) Delete(ctx context.Context, itemID string) error {
	query := "DELETE FROM items WHERE id = ?"

	_, err := r.DBManager.DB.ExecContext(ctx, query, itemID)

	if err != nil {
		r.Logger.ErrorContext(ctx, "impossible de delete", "err", err)
		return errors.New("error")
	}
	return nil
}

func (r *Repository) GetByID(ctx context.Context, itemID string) (*Model, error) {
	query := `SELECT id, name, unit, quantity, expiration_date, energy_kcal, protein_g, fat_g, carbohydrate_g, fiber_g, sugar_g, salt_g FROM items i WHERE id = ?`
	rslt, err := r.DBManager.DB.QueryContext(ctx, query, itemID)

	if err != nil {
		return nil, errors.New("error")
	}
	defer rslt.Close()

	var (
		item                                                               *Model
		id, quantity, kcal, protein, fat, carbohydrate, fiber, sugar, salt int
		name, unit                                                         string
		expDate                                                            *string
	)
	if rslt.Next() {
		err = rslt.Scan(
			&id,
			&name,
			&unit,
			&quantity,
			&expDate,
			&kcal,
			&protein,
			&fat,
			&carbohydrate,
			&fiber,
			&sugar,
			&salt,
		)
		if err != nil {
			return nil, errors.New("error")
		}

		item = NewModel(id, name, unit, quantity, kcal, protein, fat, carbohydrate, fiber, sugar, salt, expDate)
		return item, nil
	}

	if rslt.Err() != nil {
		return nil, errors.New("erreur")
	}
	return nil, errors.New("no item find")
}
