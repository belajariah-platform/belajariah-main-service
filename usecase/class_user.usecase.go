package usecase

import "belajariah-main-service/repository"

type userClassUsecase struct {
	userClassRepository repository.UserClassRepository
}

type UserClassUsecase interface{}

func InitUserClassUsecase(userClassRepository repository.UserClassRepository) UserClassUsecase {
	return &userClassUsecase{
		userClassRepository,
	}
}
