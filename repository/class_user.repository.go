package repository

import "github.com/jmoiron/sqlx"

type userClassRepository struct {
	db *sqlx.DB
}

type UserClassRepository interface{}

func InitUserClassRepository(db *sqlx.DB) UserClassRepository {
	return &userClassRepository{
		db,
	}
}
