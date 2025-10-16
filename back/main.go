package main

import (
	"api/src/core/base"
	"api/src/core/initialization"
	"api/src/db"
	"api/src/routes"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// @title GARCHAMP API
// @version 1.0
// @description Cette documentation concerne l'API en golang avec CHI du projet GARCHAMP
// @contact.url  https://github.com/HackatonM1/Garchamp

// @host localhost:8080
// @BasePath /api/v1

const readTimeout = 10
const writeTimeout = 10
const idleTimeout = 120

func main() {
	// Chargement du .env
	err := godotenv.Load()
	tmpLogger := base.NewLogger("warn", "text")
	if err != nil {
		tmpLogger.Error("Erreur lors du chargement du fichier .env", "err", err)
		os.Exit(1)
		return
	}

	// Chargement de l'application (connexion db, etc.)
	corsOptions, serverPort, logLevel, logFormat, err := initialization.LoadConfig()
	if err != nil {
		tmpLogger.Error("Error loading config", "err", err)
		os.Exit(1)
		return
	}

	// Initialisation du logger
	finalLogger := base.NewLogger(logLevel, logFormat)

	// Initialisation connexion db
	dbMain := initialization.InitDBConn(finalLogger)
	repositoryManager := db.InitRepositories(dbMain, finalLogger)

	// Fermeture des connexions à la db lors de la fin du programme
	defer dbMain.Close()

	// Setup des routes
	r := routes.SetupRouter(corsOptions, repositoryManager, finalLogger)

	// Setup du serveur
	finalLogger.Info("Server lancé", "port", serverPort)
	srv := &http.Server{
		Addr:         serverPort,
		Handler:      r,
		ReadTimeout:  readTimeout * time.Second,  // timeout pour lire la requête
		WriteTimeout: writeTimeout * time.Second, // timeout pour écrire la réponse
		IdleTimeout:  idleTimeout * time.Second,  // timeout pour la connexion keep-alive
	}

	// Lancement du serveur
	err = srv.ListenAndServe()
	if err != nil {
		finalLogger.Error("Le lancement du serveur à échoué", "erreur", err)
		return
	}
}
