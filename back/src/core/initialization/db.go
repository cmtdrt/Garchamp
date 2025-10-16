package initialization

import (
	"api/src/core/base"
	"os"
	"sync"
)

func InitDBConn(logger *base.Logger) *base.DatabaseManager {
	var (
		dbMain *base.DatabaseManager
		err    error
	)
	// Utilisation de goroutines pour initialiser les bases de données en parallèle
	wg := sync.WaitGroup{}
	const maxConcurrentTasks = 1
	wg.Add(maxConcurrentTasks)

	go func() {
		defer wg.Done()
		dbMain, err = base.NewDatabaseManager(os.Getenv("DB_CON"), logger)
		if err != nil {
			logger.Error("Impossible de se connecter à la base de données utilisateur", "erreur", err)
			os.Exit(1)
		}
	}()

	wg.Wait()
	return dbMain
}
