package repository

import "github.com/jmoiron/sqlx"

type newsRepository struct {
	db *sqlx.DB
}

type NewsRepository interface{}

func InitNewsRepository(db *sqlx.DB) NewsRepository {
	return &newsRepository{
		db,
	}
}
