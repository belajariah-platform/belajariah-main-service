package usecase

import "belajariah-main-service/repository"

type newsUsecase struct {
	newsRepository repository.NewsRepository
}

type NewsUsecase interface{}

func InitNewsUsecase(newsRepository repository.NewsRepository) NewsUsecase {
	return &newsUsecase{
		newsRepository,
	}
}
