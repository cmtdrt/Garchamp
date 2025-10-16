package base

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseManager struct {
	DB     *sql.DB
	logger *Logger
}

func NewDatabaseManager(dsn string, logger *Logger) (*DatabaseManager, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	const maxConTime = 5
	const maxOpenConns = 25 // SQLite gère moins bien la concurrence
	const maxIdleConns = 5  // Réduit pour SQLite
	const maxPingDurationContext = 2

	// Configuration adaptée pour SQLite
	db.SetConnMaxLifetime(maxConTime * time.Minute)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(maxConTime * time.Minute)

	// Active les pragmas SQLite pour de meilleures performances
	if _, err := db.Exec(`
		PRAGMA journal_mode = WAL;
		PRAGMA synchronous = NORMAL;
		PRAGMA cache_size = -64000;
		PRAGMA busy_timeout = 5000;
	`); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to set SQLite pragmas: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), maxPingDurationContext*time.Second)
	defer cancel()

	if err = db.PingContext(pingCtx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DatabaseManager{DB: db, logger: logger.With("errPosition", "dbManager")}, nil
}

func (dm *DatabaseManager) Close() {
	if err := dm.DB.Close(); err != nil {
		ctx := context.Background()
		dm.logger.ErrorContext(ctx, "impossible de fermer la connexion à la db", "erreur", err)
	}
}

func (dm *DatabaseManager) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return dm.DB.QueryContext(ctx, query, args...)
}

func (dm *DatabaseManager) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return dm.DB.ExecContext(ctx, query, args...)
}

func (dm *DatabaseManager) Transaction(ctx context.Context, txFunc func(*sql.Tx) error) error {
	tx, err := dm.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	if err = txFunc(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			dm.logger.ErrorContext(ctx, "impossible de rollback", "erreur", rbErr)
		}
		return err
	}

	return tx.Commit()
}
