package repository

import "github.com/jmoiron/sqlx"

type emailRepository struct {
	db *sqlx.DB
}

type EmailRepository interface{}

func InitEmailRepository(db *sqlx.DB) EmailRepository {
	return &emailRepository{
		db,
	}
}
