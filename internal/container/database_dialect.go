package container

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"payment-platform/config"
)

func NewDatabaseDialect(_ *config.DatabaseConfig) *goqu.DialectWrapper {
	dialect := goqu.Dialect("postgres")
	return &dialect
}
