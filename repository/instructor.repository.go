package repository

import "github.com/jmoiron/sqlx"

type instructorRepository struct {
	db *sqlx.DB
}

type InstructorRepository interface{}

func InitInstructorRepository(db *sqlx.DB) InstructorRepository {
	return &instructorRepository{
		db,
	}
}
