package handler

import "belajariah-main-service/usecase"

type testHandler struct {
	testUsecase usecase.TestUsecase
}

type TestHandler interface {
}

func InitTestHandler(testUsecase usecase.TestUsecase) TestHandler {
	return &testHandler{
		testUsecase,
	}
}
