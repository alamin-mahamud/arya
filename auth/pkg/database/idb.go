package database

import (
	"github.com/alamin-mahamud/arya/auth/pkg/config"
	"github.com/go-pg/pg"
)

type db interface {
}

func New(cfg *config.Configuration) (*pg.DB, error) {
	return NewPostgres(cfg.DB.DSN, cfg.DB.Timeout, cfg.DB.LogQueries)
}
