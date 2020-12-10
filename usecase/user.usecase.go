package usecase

import "belajariah-main-service/repository"

type userUsecase struct {
	userRepository repository.UserRepository
}

type UserUsecase interface{}

func InitUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository,
	}
}
