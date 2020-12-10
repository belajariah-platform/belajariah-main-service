package repository

import "github.com/jmoiron/sqlx"

type classRepository struct {
	db *sqlx.DB
}

type ClassRepository interface{}

func InitClassRepository(db *sqlx.DB) ClassRepository {
	return &classRepository{
		db,
	}
}
