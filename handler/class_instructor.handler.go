package handler

import "belajariah-main-service/usecase"

type instructorClassHandler struct {
	instructorClassUsecase usecase.InstructorClassUsecase
}

type InstructorClassHandler interface {
}

func InitInstructorClassHandler(instructorClassUsecase usecase.InstructorClassUsecase) InstructorClassHandler {
	return &instructorClassHandler{
		instructorClassUsecase,
	}
}
