package repository

import "github.com/jmoiron/sqlx"

type learningRepository struct {
	db *sqlx.DB
}

type LearningRepository interface{}

func InitLearningRepository(db *sqlx.DB) LearningRepository {
	return &learningRepository{
		db,
	}
}
