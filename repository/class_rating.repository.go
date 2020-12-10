package repository

import "github.com/jmoiron/sqlx"

type ratingRepository struct {
	db *sqlx.DB
}

type RatingRepository interface{}

func InitRatingRepository(db *sqlx.DB) RatingRepository {
	return &ratingRepository{
		db,
	}
}
