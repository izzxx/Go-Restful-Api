package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/izzxx/Go-Restful-Api/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

func OpenDbConnection() (*pgxpool.Pool, error) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)

	cfg, err := pgxpool.ParseConfig(psqlconn)
	if err != nil {
		return nil, err
	}

	cfg.MaxConnIdleTime = 3 * time.Minute
	cfg.MaxConnLifetime = 5 * time.Minute
	cfg.MinConns = 5
	cfg.MaxConns = 20

	ctx := context.Background()
	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

// Migrate
func Create(ctx context.Context) {
	db, err := OpenDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	queryUser := `
		CREATE TABLE IF NOT EXISTS users(
			id UUID NOT NULL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			email VARCHAR(50) NOT NULL UNIQUE,
			password TEXT NOT NULL,
			isAdmin BOOLEAN
		)
	`

	if _, err = db.Exec(ctx, queryUser); err != nil {
		log.Fatal(err)
	}

	queryProduct := `
		CREATE TABLE IF NOT EXISTS products(
			id VARCHAR(20) NOT NULL PRIMARY KEY,
			name VARCHAR(50) NOT NULL UNIQUE,
			price DOUBLE PRECISION NOT NULL,
			quantity SMALLINT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err = db.Exec(ctx, queryProduct)
	if err != nil {
		log.Fatal(err)
	}
}
