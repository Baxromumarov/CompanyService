package v1

import (
	"github.com/baxromumarov/CompanyService/config"
	"github.com/baxromumarov/CompanyService/internal/db"
	"github.com/baxromumarov/CompanyService/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type handlerV1 struct {
	log             logger.Logger
	cfg             *config.Config
	storagePostgres db.StorageI
}

type HandlerV1Options struct {
	Log logger.Logger
	Cfg *config.Config
	Db  *sqlx.DB
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		log:             options.Log,
		cfg:             options.Cfg,
		storagePostgres: db.NewStorage(options.Db, options.Log),
	}
}

func (h *handlerV1) Log() logger.Logger {
	return h.log
}

func (h *handlerV1) Config() *config.Config {
	return h.cfg
}
