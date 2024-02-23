package container

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"log"
	"payment-platform/config"
	"time"
)

func NewDatabase(config *config.DatabaseConfig) *pgxpool.Pool {
	ctx := context.Background()
	dbPoolConfig, err := pgxpool.ParseConfig(config.DSN)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v", err)
	}

	dbPoolConfig.MaxConns = 100
	dbPoolConfig.MinConns = 10
	dbPoolConfig.MaxConnLifetime = time.Hour
	dbPoolConfig.MaxConnIdleTime = 30 * time.Minute
	dbPoolConfig.HealthCheckPeriod = 1 * time.Minute

	dbPool, err := pgxpool.NewWithConfig(ctx, dbPoolConfig)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	err = dbPool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping the database: %v", err)
	}

	return dbPool
}
