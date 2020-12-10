package usecase

import "belajariah-main-service/repository"

type testUsecase struct {
	testRepository repository.TestRepository
}

type TestUsecase interface{}

func InitTestUsecase(testRepository repository.TestRepository) TestUsecase {
	return &testUsecase{
		testRepository,
	}
}
