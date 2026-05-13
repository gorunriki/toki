package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(dbUrl string) *pgxpool.Pool {
	log.Println(dbUrl)
	db, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatal("Failed connect DB:", err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal("DB not reachable:", err)
	}

	log.Println("Connected to PostgreSQL")
	return db
}
