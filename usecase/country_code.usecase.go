package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/utils"
)

type countryCodeUsecase struct {
	countryCodeRepository repository.CountryCodeRepository
}

type CountryCodeUsecase interface {
	GetAllCountryCode(r model.CountryCodeRequest) ([]model.CountryCode, error)
}

func InitCountryCodeUsecase(cp repository.CountryCodeRepository) CountryCodeUsecase {
	return &countryCodeUsecase{
		cp,
	}
}

func (u *countryCodeUsecase) GetAllCountryCode(r model.CountryCodeRequest) ([]model.CountryCode, error) {
	var orderDefault = "ORDER BY country asc"
	var filterDefault = "is_deleted = false and is_active = true"
	filterFinal := utils.GetFilterOrderHandler(filterDefault, orderDefault, r.Query)

	result, err := u.countryCodeRepository.GetAllCountryCode(filterFinal)
	if err != nil {
		return nil, utils.WrapError(err, "countryCodeUsecase.GetAllCountryCode")
	}

	return *result, nil
}
