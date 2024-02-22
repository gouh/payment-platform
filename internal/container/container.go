package container

import (
	"database/sql"
	"payment-platform/config"
)

type Container struct {
	Db     *sql.DB
	Config *config.Config
}

func NewContainer(config *config.Config) *Container {
	return &Container{
		Db:     NewDatabase(config.Db),
		Config: config,
	}
}
