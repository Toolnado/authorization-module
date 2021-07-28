package database

import (
	"context"
	"os"

	"github.com/Toolnado/authorization-module/internal/model"
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

func (p *PostgresStore) CreateUser(ctx context.Context, user *model.User) (string, error) {
	return "", nil
}
