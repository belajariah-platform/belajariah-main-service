package usecase

import "belajariah-main-service/repository"

type classUsecase struct {
	classRepository repository.ClassRepository
}

type ClassUsecase interface{}

func InitClassUsecase(classRepository repository.ClassRepository) ClassUsecase {
	return &classUsecase{
		classRepository,
	}
}
