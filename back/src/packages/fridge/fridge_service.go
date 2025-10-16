package fridge

import (
	"api/src/core/base"
	"api/src/db"
	"context"
)

type service struct {
	repositoryManager *db.RepositoryManager
	logger            *base.Logger
}

func newService(repositoryManager *db.RepositoryManager, logger *base.Logger) *service {
	sLogger := logger.With("service", "fridge")
	return &service{
		repositoryManager: repositoryManager,
		logger:            sLogger,
	}
}

// func (s *service) get(ctx context.Context, req *createReq) error {
// 	nItem := itemdb.NewModel(req.Name, req.Unity, req.Quantity, &req.ExpDate)
// }

func (s *service) create(ctx context.Context, req *createReq) error {
	err := s.repositoryManager.GetItemRepo().Create(ctx, req.Name, req.Unity, req.Quantity, &req.ExpDate)
	return err
}
