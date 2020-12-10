package repository

import "github.com/jmoiron/sqlx"

type instructorClassRepository struct {
	db *sqlx.DB
}

type InstructorClassRepository interface{}

func InitInstructorClassRepository(db *sqlx.DB) InstructorClassRepository {
	return &instructorClassRepository{
		db,
	}
}
