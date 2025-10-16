package itemallergenrelationdb

import "api/src/core/base"

type Repository struct {
	DbManager *base.DatabaseManager
	Logger    *base.Logger
}

func NewRepository(db *base.DatabaseManager, logger *base.Logger) *Repository {
	return &Repository{DbManager: db, Logger: logger}
}
