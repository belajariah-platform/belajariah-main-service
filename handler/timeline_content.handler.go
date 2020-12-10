package handler

import "belajariah-main-service/usecase"

type contentHandler struct {
	contentUsecase usecase.ContentUsecase
}

type ContentHandler interface {
}

func InitContentHandler(contentUsecase usecase.ContentUsecase) ContentHandler {
	return &contentHandler{
		contentUsecase,
	}
}
