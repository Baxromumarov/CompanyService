package db

import (
	"fmt"

	"github.com/baxromumarov/CompanyService/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPostgres(cfg *config.Config) (*sqlx.DB, error) {
	// connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
	// 	cfg.PostgresUser,
	// 	cfg.PostgresPassword,
	// 	cfg.PostgresHost,
	// 	cfg.PostgresPort,
	// 	cfg.PostgresDatabase,
	// )
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
		cfg.PostgresHost,
		cfg.PostgresPort,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil

}
