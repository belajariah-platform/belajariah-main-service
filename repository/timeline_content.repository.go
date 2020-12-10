package repository

import "github.com/jmoiron/sqlx"

type contentRepository struct {
	db *sqlx.DB
}

type ContentRepository interface{}

func InitContentRepository(db *sqlx.DB) ContentRepository {
	return &contentRepository{
		db,
	}
}
