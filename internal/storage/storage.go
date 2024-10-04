package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log"
	"ministry/config"
	"ministry/internal/storage/postgres"
	r "ministry/internal/storage/redis"
	"ministry/pkg/logger"
)

type Storage struct {
	Postgres *sql.DB
	Redis    *redis.Client
}

var Store Storage

func New(log *logger.Logger, cfg *config.Config) *Storage {
	Store.Postgres = postgres.NewPostgresDB(log, cfg)
	Store.Redis = r.NewRedisClient(log, cfg)
	return &Store
}

func CloseStorage() {
	if err := Store.Postgres.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("Postgres DB is closed")
	if err := Store.Redis.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("Redis client is closed")
}
