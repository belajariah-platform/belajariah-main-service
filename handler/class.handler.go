package handler

import "belajariah-main-service/usecase"

type classHandler struct {
	classUsecase usecase.ClassUsecase
}

type ClassHandler interface {
}

func InitClassHandler(classUsecase usecase.ClassUsecase) ClassHandler {
	return &classHandler{
		classUsecase,
	}
}
