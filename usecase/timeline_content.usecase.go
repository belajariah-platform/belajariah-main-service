package usecase

import "belajariah-main-service/repository"

type contentUsecase struct {
	contentRepository repository.ContentRepository
}

type ContentUsecase interface{}

func InitContentUsecase(contentRepository repository.ContentRepository) ContentUsecase {
	return &contentUsecase{
		contentRepository,
	}
}
