package container

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5/pgxpool"
	"payment-platform/config"
)

type Container struct {
	Db        *pgxpool.Pool
	DbDialect *goqu.DialectWrapper
	Config    *config.Config
}

func NewContainer(config *config.Config) *Container {
	return &Container{
		Db:        NewDatabase(config.Db),
		DbDialect: NewDatabaseDialect(config.Db),
		Config:    config,
	}
}
