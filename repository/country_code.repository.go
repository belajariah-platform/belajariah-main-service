package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	_getAllCountryCodeSql = `
	SELECT 
		id, 
		code, 
		country,
		number_code,
		flag,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM
		master.master_country_number_code
	%s
`
)

type countryCodeRepository struct {
	db *sqlx.DB
}

type CountryCodeRepository interface {
	GetAllCountryCode(filter string) (*[]model.CountryCode, error)
}

func InitCountryCodeRepository(db *sqlx.DB) CountryCodeRepository {
	return &countryCodeRepository{
		db,
	}
}

func (r *countryCodeRepository) GetAllCountryCode(filter string) (*[]model.CountryCode, error) {
	var result []model.CountryCode
	query := fmt.Sprintf(_getAllCountryCodeSql, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "countryCodeRepository.GetAllCountryCode :  error get")
	}

	return &result, nil
}
