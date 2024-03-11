package repository

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Repository {
	return &Repository{db}
}
