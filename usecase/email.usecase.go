package usecase

import "belajariah-main-service/repository"

type emailUsecase struct {
	emailRepository repository.EmailRepository
}

type EmailUsecase interface{}

func InitEmailUsecase(emailRepository repository.EmailRepository) EmailUsecase {
	return &emailUsecase{
		emailRepository,
	}
}
