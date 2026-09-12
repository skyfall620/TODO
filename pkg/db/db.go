package db

import (
	"database/sql"
	"log"
	"todo/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgreDb(conf *config.Config) *sql.DB {
	db, err := sql.Open("pgx", conf.DBConnString())
	if err != nil {
		log.Fatalf("couldn't open the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Ping failed: %v", err)
	}

	return db
}
