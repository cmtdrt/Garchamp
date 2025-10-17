package recipe

import (
	"api/src/core/base"
	"api/src/core/utils"
	"api/src/db"
	"context"
	"encoding/json"
	"fmt"
	"strings"
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

type ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

type recipeAIResponse struct {
	Title         string       `json:"title"`
	Description   string       `json:"description"`
	Ingredients   []ingredient `json:"ingredients"`
	Steps         []string     `json:"steps"`
	EstimatedTime string       `json:"estimated_time"`
	Difficulty    string       `json:"difficulty"`
	Error         string       `json:"error,omitempty"`
}

func (s *service) create(ctx context.Context, req *createReq) (*recipeAIResponse, error) {
	// Préparation des datas
	ingredientsList := ""
	if len(req.Items) == 0 {
		ingredientsList = "Libre : utilisez les ingrédients que vous souhaitez."
	} else {
		for _, item := range req.Items {
			ingredientsList += fmt.Sprintf("- %d %s de %s\n", item.Quantity, item.Unit, item.Name)
		}
	}
	allergenList := "aucun"
	if len(req.Allergens) > 0 {
		allergenList = strings.Join(req.Allergens, ", ")
	}
	// Récupération des allergènes
	client := utils.NewOllamaClient("http://localhost:11434")
	prompt := fmt.Sprintf(`
Tu es un chef cuisinier expert en nutrition et en sécurité alimentaire.

Ta mission est de créer une recette détaillée et réalisable à partir des informations suivantes :

Ingrédients disponibles : %s  
Nombre de personnes : %d  
Allergènes à éviter : %s  

⚙️ Contraintes :
- Utilise uniquement les ingrédients listés (ou des variantes très proches si nécessaire).
- Respecte strictement la liste des allergènes à éviter.
- Ajuste les quantités pour le nombre de personnes indiqué.
- Si un ingrédient est manquant pour équilibrer la recette, propose une alternative sans allergène.

🧾 Réponse attendue au format JSON :
{
  "title": "Nom de la recette",
  "description": "Courte description appétissante du plat",
  "ingredients": [
    {"name": "nom de l’ingrédient", "quantity": "quantité et unité"}
  ],
  "steps": [
    "Étape 1 ...",
    "Étape 2 ...",
    "Étape 3 ..."
  ],
  "estimated_time": "durée estimée en minutes",
  "difficulty": "facile | moyen | difficile"
}

⚠️ Important :
- il faut répondre en français
- Si aucune recette sûre (sans allergène) ne peut être faite avec ces ingrédients, renvoie :
  {"error": "Aucune recette possible sans les allergènes indiqués."}
		`, ingredientsList, req.PeopleNumber, allergenList)

	// Appel avec affichage du stream
	response, err := client.Prompt(
		ctx,
		"mistral:instruct",
		prompt,
		*s.logger,
	)

	if err != nil {
		return nil, err
	}

	var AIResponse = recipeAIResponse{}
	if err = json.Unmarshal([]byte(response), &AIResponse); err != nil {
		return nil, fmt.Errorf("erreur parsing JSON Mistral: %w", err)
	}
	return &AIResponse, nil
}
