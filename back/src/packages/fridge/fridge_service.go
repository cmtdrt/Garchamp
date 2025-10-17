package fridge

import (
	"api/src/core/base"
	"api/src/core/utils"
	"api/src/db"
	"context"
	"encoding/json"
	"errors"
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

func (s *service) create(ctx context.Context, req *createReq) error {
	// Récupération des allergènes
	client := utils.NewOllamaClient("http://localhost:11434")
	prompt1 := fmt.Sprintf(`
Tu es un assistant qui détecte les allergènes dans un produit alimentaire.

Item: "%s"

Liste uniquement les allergènes possibles parmi : gluten, lait, oeufs, fruits à coque, soja, poisson, crustacés, arachides.

Réponse format JSON : {"allergens": [...]}.

⚠️ Important : si aucun allergène n'est présent, renvoie [].
		`, fmt.Sprintf("%s %d %s", req.Name, req.Quantity, req.Unity))

	prompt2 := fmt.Sprintf(`
Tu es un assistant qui fournit les valeurs nutritionnelles d'un aliment.

Item: "%s"

Réponse format JSON :
{
"Kcal": 0,
"Protein": 0,
"Fat": 0,
"Carbohydrate": 0,
"Fiber": 0,
"Sugar": 0,
"Salt": 0
}

⚠️ Important : 
- Si une valeur nutritionnelle est inconnue, renvoie 0.
- Donne les valeurs pour la quantité donnée, arrondies si nécessaire.
- Ne renvoie aucun commentaire, juste le JSON.
		`, fmt.Sprintf("%s %d %s", req.Name, req.Quantity, req.Unity))
	// Appel avec affichage du stream
	response, err := client.Prompt(
		ctx,
		"mistral:instruct",
		prompt2,
		*s.logger,
	)

	if err != nil {
		return err
	}

	var result2 struct {
		Kcal         float64 `json:"Kcal"`
		Protein      float64 `json:"Protein"`
		Fat          float64 `json:"Fat"`
		Carbohydrate float64 `json:"Carbohydrate"`
		Fiber        float64 `json:"Fiber"`
		Sugar        float64 `json:"Sugar"`
		Salt         float64 `json:"Salt"`
	}
	if err = json.Unmarshal([]byte(response), &result2); err != nil {
		return fmt.Errorf("erreur parsing JSON Mistral: %w", err)
	}

	var (
		allergenRepo    = s.repositoryManager.GetAllergenRepo()
		itemAlergenRepo = s.repositoryManager.GetitemAllergenRelationRepo()
		itemRepo        = s.repositoryManager.GetItemRepo()
	)

	// Création de l'item
	res, err := itemRepo.Create(
		ctx,
		req.Name,
		req.Unity,
		req.Quantity,
		result2.Kcal,
		result2.Protein,
		result2.Fat,
		result2.Carbohydrate,
		result2.Fiber,
		result2.Sugar,
		result2.Salt,
		&req.ExpDate,
	)
	if err != nil {
		return err
	}
	itemID, _ := res.LastInsertId()

	var result1 struct {
		Allergens []string `json:"allergens"`
	}
	// Appel avec affichage du stream
	response, err = client.Prompt(
		ctx,
		"mistral:instruct",
		prompt1,
		*s.logger,
	)
	if err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(response), &result1); err != nil {
		return fmt.Errorf("erreur parsing JSON Mistral: %w", err)
	}
	// Lien avec les alergènes
	for _, allergenName := range result1.Allergens {
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

func (s *service) getAll(ctx context.Context) ([]Item, error) {
	var (
		itemRepo      = s.repositoryManager.GetItemRepo()
		allergenRepo  = s.repositoryManager.GetAllergenRepo()
		item          *Item
		result        = []Item{}
		allergensName []string
		err           error
	)

	items, err := itemRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	for _, i := range items {
		allergensName, err = allergenRepo.GetAllAllergensByRelation(ctx, i.ID)
		if err != nil {
			return nil, err
		}
		item = newItem(i.ID, i.Quantity, i.Name, i.Unit, i.ExpDate, allergensName)
		result = append(result, *item)
	}
	fmt.Println("resultat du getAll", result)
	return result, nil
}

func (s *service) deleteItemByID(ctx context.Context, itemID string) error {
	return s.repositoryManager.GetItemRepo().Delete(ctx, itemID)
}

func (s *service) getItemMacroByID(ctx context.Context, itemID string) (*itemDetails, error) {
	item, err := s.repositoryManager.GetItemRepo().GetByID(ctx, itemID)
	if err != nil || item == nil {
		return nil, errors.New("error")
	}
	return newItemDetails(item.Kcal, item.Protein, item.Fat, item.Carbohydrate, item.Fiber, item.Sugar, item.Salt), nil
}
