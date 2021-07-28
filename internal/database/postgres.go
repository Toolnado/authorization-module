package database

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PostgresStore struct {
	db *sqlx.DB
}

func NewStore() *PostgresStore {
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		logrus.Fatalf("Failed to establish a connection to the database: %s", err.Error())
	}
	return &PostgresStore{
		db: db,
	}
}
