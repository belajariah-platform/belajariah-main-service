package repository

import "github.com/jmoiron/sqlx"

type userRepository struct {
	db *sqlx.DB
}

type UserRepository interface{}

func InitUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db,
	}
}
