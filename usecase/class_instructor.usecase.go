package usecase

import "belajariah-main-service/repository"

type instructorClassUsecase struct {
	instructorClassRepository repository.InstructorClassRepository
}

type InstructorClassUsecase interface{}

func InitInstructorClassUsecase(instructorClassRepository repository.InstructorClassRepository) InstructorClassUsecase {
	return &instructorClassUsecase{
		instructorClassRepository,
	}
}
