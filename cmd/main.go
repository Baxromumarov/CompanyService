package main

import (
	"github.com/baxromumarov/CompanyService/config"
	"github.com/baxromumarov/CompanyService/internal/api"
	"github.com/baxromumarov/CompanyService/internal/db"
	"github.com/baxromumarov/CompanyService/pkg/logger"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "CompanyService")
	db, err := db.InitPostgres(cfg)
	if err != nil {
		log.Fatal("error while connecting to postgres: " + err.Error())
	}

	server := api.New(&api.RouterOptions{
		Log: log,
		Cfg: cfg,
		Db:  db,
	})

	server.Run(":" + cfg.HttpPort)
}
