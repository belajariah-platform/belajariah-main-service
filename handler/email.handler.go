package handler

import "belajariah-main-service/usecase"

type emailHandler struct {
	emailUsecase usecase.EmailUsecase
}

type EmailHandler interface {
}

func InitEmailHandler(emailUsecase usecase.EmailUsecase) EmailHandler {
	return &emailHandler{
		emailUsecase,
	}
}
