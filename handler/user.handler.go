package handler

import "belajariah-main-service/usecase"

type userHandler struct {
	userUsecase usecase.UserUsecase
}

type UserHandler interface {
}

func InitUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase,
	}
}
