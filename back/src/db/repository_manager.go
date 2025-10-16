package db

import (
	"api/src/core/base"
	itemdb "api/src/db/item"
	itemallergenrelationdb "api/src/db/item_allergen_relation"
)

type RepositoryManager struct {
	logger                         *base.Logger
	dbMain                         *base.DatabaseManager
	itemRepository                 *itemdb.Repository
	itemAllergenRelationRepository *itemallergenrelationdb.Repository
}

func NewRepositoryManager(
	dbMain *base.DatabaseManager,
	itemRepository *itemdb.Repository,
	itemAllergenRelationRepository *itemallergenrelationdb.Repository,
) *RepositoryManager {
	return &RepositoryManager{
		dbMain:                         dbMain,
		itemAllergenRelationRepository: itemAllergenRelationRepository,
		itemRepository:                 itemRepository,
	}
}

func InitRepositories(dbMain *base.DatabaseManager, logger *base.Logger) *RepositoryManager {
	return NewRepositoryManager(
		dbMain,
		itemdb.NewRepository(dbMain, logger),
		itemallergenrelationdb.NewRepository(dbMain, logger),
	)
}

// GetItemRepo permet l'accès à itemRepository.
func (rm *RepositoryManager) GetItemRepo() *itemdb.Repository {
	if rm.itemRepository == nil {
		rm.logger.Error("itemAllergenRelationRepository non initialisé")
	}
	return rm.itemRepository
}

// GetitemAllergenRelationRepo permet l'accès à itemRepository.
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
