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
	if err != nil {
		return err
	}
	// client := utils.NewOllamaClient("http://localhost:11434")

	// // Appel avec affichage du stream
	// response, err := client.Prompt(
	// 	"mistral:instruct",
	// 	"Salut, peux-tu me donner les alergènes de l'ingrédient : "+req.Name,
	// )

	// if err != nil {
	// 	fmt.Printf("❌ Erreur: %v\n", err)
	// 	return err
	// }
	// fmt.Println(response)
	return nil
}
