package handler

import "belajariah-main-service/usecase"

type instructorHandler struct {
	instructorUsecase usecase.InstructorUsecase
}

type InstructorHandler interface {
}

func InitInstructorHandler(instructorUsecase usecase.InstructorUsecase) InstructorHandler {
	return &instructorHandler{
		instructorUsecase,
	}
}
