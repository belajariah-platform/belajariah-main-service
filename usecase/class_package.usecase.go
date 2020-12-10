package usecase

import "belajariah-main-service/repository"

type packageUsecase struct {
	packageRepository repository.PackageRepository
}

type PackageUsecase interface{}

func InitPackageUsecase(packageRepository repository.PackageRepository) PackageUsecase {
	return &packageUsecase{
		packageRepository,
	}
}
