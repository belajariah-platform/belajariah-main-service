package handler

import "belajariah-main-service/usecase"

type ratingHandler struct {
	ratingUsecase usecase.RatingUsecase
}

type RatingHandler interface {
}

func InitRatingHandler(ratingUsecase usecase.RatingUsecase) RatingHandler {
	return &ratingHandler{
		ratingUsecase,
	}
}
