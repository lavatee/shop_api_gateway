package repository

import "github.com/jmoiron/sqlx"

const usersTable = "users"

type Repository struct {
	Auth
}

type Auth interface {
	SignUp(name string, email string, passwordHash string) (int, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
	}
}
