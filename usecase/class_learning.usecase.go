package usecase

import "belajariah-main-service/repository"

type learningUsecase struct {
	learningRepository repository.LearningRepository
}

type LearningUsecase interface{}

func InitLearningUsecase(learningRepository repository.LearningRepository) LearningUsecase {
	return &learningUsecase{
		learningRepository,
	}
}
