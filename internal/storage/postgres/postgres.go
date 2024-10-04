package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"ministry/config"
	"ministry/pkg/logger"
)

func NewPostgresDB(logger *logger.Logger, cfg *config.Config) *sql.DB {
	database := cfg.Database

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		database.Host, database.Port, database.User, database.Password, database.Name, database.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}

	_, err = db.Exec("SET timezone TO 'Asia/Dushanbe'")
	if err != nil {
		logger.Println(err)
	}

	return db
}
