// Package utils contient des fonctions utilitaires globales pour le projet.
//
//revive:disable:var-naming
package utils

import (
	"api/src/core/base"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// OllamaClient gère les appels à Ollama.
type OllamaClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// OllamaRequest structure de la requête Ollama.
type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

// OllamaResponse structure de chaque ligne de réponse en streaming.
type OllamaResponse struct {
	Model     string `json:"model"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"created_at,omitempty"`
}

const timeout = 500

// NewOllamaClient crée un nouveau client Ollama.
func NewOllamaClient(baseURL string) *OllamaClient {
	return &OllamaClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: timeout * time.Second, // Timeout plus long pour le streaming
		},
	}
}

// Prompt envoie un prompt à Ollama et récupère la réponse complète.
func (c *OllamaClient) Prompt(ctx context.Context, model, prompt string, logger base.Logger) (string, error) {
	logger.InfoContext(ctx, "🟢 Requête envoyée à Ollama...")
	// Prépare la requête
	reqBody := OllamaRequest{
		Model:  model,
		Prompt: prompt,
		Stream: true,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("erreur marshalling: %w", err)
	}

	// Crée la requête HTTP
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+"/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("erreur création requête: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Envoie la requête
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("erreur requête HTTP: %w", err)
	}
	defer resp.Body.Close()

	// Vérifie le status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("erreur HTTP %d: %s", resp.StatusCode, string(body))
	}

	logger.InfoContext(ctx, "🟢 Réponse reçue :\n")

	// Lit le stream ligne par ligne
	fullText := ""
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		// Décode chaque ligne JSON
		var ollamaResp OllamaResponse
		if err = json.Unmarshal(line, &ollamaResp); err != nil {
			return "", fmt.Errorf("erreur unmarshalling: %w", err)
		}

		// Affiche et accumule la réponse
		if ollamaResp.Response != "" {
			// Print obligatoire pour l'affichage direct
			//
			//revive:disable:forbidigo
			fmt.Print(ollamaResp.Response)
			//revive:enable:forbidigo
			fullText += ollamaResp.Response
		}

		// Si done=true, c'est fini
		if ollamaResp.Done {
			break
		}
	}

	if err = scanner.Err(); err != nil {
		return "", fmt.Errorf("erreur lecture stream: %w", err)
	}

	return fullText, nil
}

// PromptSilent version sans affichage en temps réel (juste retour final).
func (c *OllamaClient) PromptSilent(ctx context.Context, model, prompt string) (string, error) {
	reqBody := OllamaRequest{
		Model:  model,
		Prompt: prompt,
		Stream: true,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("erreur marshalling: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+"/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("erreur création requête: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("erreur requête HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("erreur HTTP %d: %s", resp.StatusCode, string(body))
	}

	fullText := ""
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		var ollamaResp OllamaResponse
		if err = json.Unmarshal(line, &ollamaResp); err != nil {
			return "", fmt.Errorf("erreur unmarshalling: %w", err)
		}

		fullText += ollamaResp.Response

		if ollamaResp.Done {
			break
		}
	}

	if err = scanner.Err(); err != nil {
		return "", fmt.Errorf("erreur lecture stream: %w", err)
	}

	return fullText, nil
}
