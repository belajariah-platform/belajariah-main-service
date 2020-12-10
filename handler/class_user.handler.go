package handler

import "belajariah-main-service/usecase"

type userClassHandler struct {
	userClassUsecase usecase.UserClassUsecase
}

type UserClassHandler interface {
}

func InitUserClassHandler(userClassUsecase usecase.UserClassUsecase) UserClassHandler {
	return &userClassHandler{
		userClassUsecase,
	}
}
