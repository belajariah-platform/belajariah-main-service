package usecase

import "belajariah-main-service/repository"

type ratingUsecase struct {
	ratingRepository repository.RatingRepository
}

type RatingUsecase interface{}

func InitRatingUsecase(ratingRepository repository.RatingRepository) RatingUsecase {
	return &ratingUsecase{
		ratingRepository,
	}
}
