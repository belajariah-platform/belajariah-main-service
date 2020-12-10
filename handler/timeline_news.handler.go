package handler

import "belajariah-main-service/usecase"

type newsHandler struct {
	newsUsecase usecase.NewsUsecase
}

type NewsHandler interface {
}

func InitNewsHandler(newsUsecase usecase.NewsUsecase) NewsHandler {
	return &newsHandler{
		newsUsecase,
	}
}
