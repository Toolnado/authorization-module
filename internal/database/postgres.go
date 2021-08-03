package database

import (
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/Toolnado/authorization-module/internal/model"
	"github.com/jmoiron/sqlx"
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

func (p *PostgresStore) CreateUser(user *model.User) (int, error) {
	hashPassword := hashPassword(user.Password)
	id := 0

	query := fmt.Sprintln("INSERT INTO users (name, username, hashpassword) VALUES ($1, $2, $3) RETURNING id;")

	p.db.QueryRow(query, user.Name, user.Username, hashPassword).Scan(&id)

	return id, nil
}

func (p *PostgresStore) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintln("SELECT id FROM users WHERE username=$1 AND hashpassword=$2")

	if err := p.db.Get(&user, query, username, hashPassword(password)); err != nil {
		return user, err
	}

	return user, nil
}

func hashPassword(password string) []byte {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	hashPassword := hasher.Sum([]byte(password))

	return hashPassword
}
