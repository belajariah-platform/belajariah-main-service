package usecase

import "belajariah-main-service/repository"

type instructorUsecase struct {
	instructorRepository repository.InstructorRepository
}

type InstructorUsecase interface{}

func InitInstructorUsecase(instructorRepository repository.InstructorRepository) InstructorUsecase {
	return &instructorUsecase{
		instructorRepository,
	}
}
