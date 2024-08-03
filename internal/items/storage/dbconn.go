package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/ruziba3vich/boock/internal/items/config"
)

func ConnectDB(config *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		config.Database.User,
		config.Database.DBName,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
