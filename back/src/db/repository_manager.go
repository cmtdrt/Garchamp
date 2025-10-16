package db

import (
	"api/src/core/base"
	itemallergenrelationdb "api/src/db/item_allergen_relation"
)

type RepositoryManager struct {
	logger                         *base.Logger
	dbMain                         *base.DatabaseManager
	itemAllergenRelationRepository *itemallergenrelationdb.Repository
}

func NewRepositoryManager(
	dbMain *base.DatabaseManager,
	itemAllergenRelationRepository *itemallergenrelationdb.Repository,
) *RepositoryManager {
	return &RepositoryManager{
		dbMain:                         dbMain,
		itemAllergenRelationRepository: itemAllergenRelationRepository,
	}
}

func InitRepositories(dbMain *base.DatabaseManager, logger *base.Logger) *RepositoryManager {
	return NewRepositoryManager(
		dbMain,
		itemallergenrelationdb.NewRepository(dbMain, logger),
	)
}

// GetitemAllergenRelationRepo permet l'accès à itemAllergenRelationRepository.
func (rm *RepositoryManager) GetitemAllergenRelationRepo() *itemallergenrelationdb.Repository {
	if rm.itemAllergenRelationRepository == nil {
		rm.logger.Error("itemAllergenRelationRepository non initialisé")
	}
	return rm.itemAllergenRelationRepository
}

// GetDBMain permet l'accès au db manager de la base de données utilisateur.
func (rm *RepositoryManager) GetDBMain() *base.DatabaseManager {
	return rm.dbMain
}
