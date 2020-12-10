package handler

import "belajariah-main-service/usecase"

type learningHandler struct {
	learningUsecase usecase.LearningUsecase
}

type LearningHandler interface {
}

func InitLearningHandler(learningUsecase usecase.LearningUsecase) LearningHandler {
	return &learningHandler{
		learningUsecase,
	}
}
