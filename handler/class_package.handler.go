package handler

import "belajariah-main-service/usecase"

type packageHandler struct {
	packageUsecase usecase.PackageUsecase
}

type PackageHandler interface {
}

func InitPackageHandler(packageUsecase usecase.PackageUsecase) PackageHandler {
	return &packageHandler{
		packageUsecase,
	}
}
