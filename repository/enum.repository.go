package repository

import "github.com/jmoiron/sqlx"

type enumRepository struct {
	db *sqlx.DB
}

type EnumRepository interface{}

func InitEnumRepository(db *sqlx.DB) EnumRepository {
	return &enumRepository{
		db,
	}
}
