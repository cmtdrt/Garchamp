package fridge

import (
	"api/src/core/base"
	"api/src/core/utils"
	"api/src/db"
	"context"
	"encoding/json"
	"fmt"
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
	// Création de l'item
	res, err := s.repositoryManager.GetItemRepo().Create(ctx, req.Name, req.Unity, req.Quantity, &req.ExpDate)
	if err != nil {
		return err
	}
	itemID, _ := res.LastInsertId()
	// Récupération des allergènes
	client := utils.NewOllamaClient("http://localhost:11434")
	prompt := fmt.Sprintf(`
			Tu es un assistant qui détecte les allergènes dans un produit alimentaire.

			Item: "%s"

			Liste uniquement les allergènes possibles parmi : gluten, lait, oeufs, fruits à coque, soja, poisson, crustacés, arachides.

			Réponse format JSON : {"allergens": [...]}.

			⚠️ Important : si aucun allergène n'est présent dans le produit, renvoie {"allergens": []} (un tableau vide) au lieu de null ou de texte.
		`, req.Name)
	// Appel avec affichage du stream
	response, err := client.Prompt(
		ctx,
		"mistral:instruct",
		prompt,
		*s.logger,
	)

	if err != nil {
		return err
	}

	var result struct {
		Allergens []string `json:"allergens"`
	}
	if err = json.Unmarshal([]byte(response), &result); err != nil {
		return fmt.Errorf("erreur parsing JSON Mistral: %w", err)
	}

	var (
		allergenRepo    = s.repositoryManager.GetAllergenRepo()
		itemAlergenRepo = s.repositoryManager.GetitemAllergenRelationRepo()
	)

	// Lien avec les alergènes
	for _, allergenName := range result.Allergens {
		allergenID := allergenRepo.FindByName(ctx, allergenName)
		if allergenID != -1 {
			// Ajouter lien item_allergens
			_, err = itemAlergenRepo.Create(ctx, itemID, int64(allergenID))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
