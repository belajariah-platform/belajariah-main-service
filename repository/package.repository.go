package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	_getPackage = `
		SELECT
			id,
			code,
			class_code,
			type,
			price_package,
			price_discount,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted,
			duration,
			duration_frequence,
			consultation,
			webinar
		FROM master.master_package
		WHERE 
			is_deleted = false AND
			is_active=true AND
			code=$1
	`
	_getAllPackage = `
		SELECT
			id,
			code,
			class_code,
			type,
			price_package,
			price_discount,
			description,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted,
			duration,
			duration_frequence,
			consultation,
			webinar
		FROM master.master_package
		WHERE 
			is_deleted = false AND
			is_active=true
		%s
		OFFSET %d
		LIMIT %d
	`
	_getAllPackageCount = `
		SELECT COUNT(*) FROM 
			master.master_package  
		WHERE 
			is_deleted=false AND
			is_active=true
		%s
	`
	_getAllPackageQuran = `
	SELECT
			id,
			code,
			class_code,
			mentor_code,
			type,
			price_package,
			price_discount,
			description,
			duration,
			duration_frequence
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM master.master_package_quran
		WHERE 
			is_deleted = false AND
			is_active=true
		%s
	`
	_getAllBenefit = `
		SELECT
			id,
			code,
			class_code,
			description,
			icon_benefit,
			sequence,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM master.master_benefit
		WHERE 
			is_deleted = false AND
			is_active=true
		%s
		ORDER BY sequence ASC
		OFFSET %d
		LIMIT %d
	`
	_getAllBenefitQuran = `
		SELECT
			id,
			code,
			class_code,
			description,
			icon_benefit,
			sequence,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM master.master_benefit_quran
		WHERE 
			is_deleted = false AND
			is_active=true
		%s
		ORDER BY sequence ASC
	`
	_getAllTermConditionQuran = `
		SELECT
			id,
			code,
			class_code,
			description,
			icon_term,
			sequence,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM master.master_term_condition_quran
		WHERE 
			is_deleted = false AND
			is_active=true
		%s
		ORDER BY sequence ASC
	`
)

type packageRepository struct {
	db *sqlx.DB
}

type PackageRepository interface {
	GetPackage(code string) (model.Package, error)
	GetAllPackage(skip, take int, filter string) ([]model.Package, error)
	GetAllPackageCount(filter string) (int, error)

	GetAllPackageQuran(filter string) (*[]model.PackageQuran, error)

	GetAllBenefit(skip, take int, filter string) ([]model.Benefit, error)
	GetAllBenefitQuran(filter string) (*[]model.BenefitQuran, error)

	GetAllTermConditionQuran(filter string) (*[]model.TermConditionQuran, error)
}

func InitPackageRepository(db *sqlx.DB) PackageRepository {
	return &packageRepository{
		db,
	}
}

func (r *packageRepository) GetPackage(codes string) (model.Package, error) {
	var packageRow model.Package
	row := r.db.QueryRow(_getPackage, codes)

	var id, duration int
	var createdDate time.Time
	var isActive, isDeleted bool
	var modifiedDate sql.NullTime
	var PriceDiscount, modifiedBy sql.NullString
	var consultation, webinar, durationFrequence sql.NullInt64
	var types, classCode, pricePackage, code, createdBy string

	sqlError := row.Scan(
		&id,
		&code,
		&classCode,
		&types,
		&pricePackage,
		&PriceDiscount,
		&isActive,
		&createdBy,
		&createdDate,
		&modifiedBy,
		&modifiedDate,
		&isDeleted,
		&duration,
		&durationFrequence,
		&consultation,
		&webinar,
	)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetPackage => ", sqlError.Error())
		return model.Package{}, nil
	} else {
		packageRow = model.Package{
			ID:                id,
			Code:              code,
			ClassCode:         classCode,
			Type:              types,
			PricePackage:      pricePackage,
			PriceDiscount:     PriceDiscount,
			IsActive:          isActive,
			CreatedBy:         createdBy,
			CreatedDate:       createdDate,
			ModifiedBy:        modifiedBy,
			ModifiedDate:      modifiedDate,
			IsDeleted:         isDeleted,
			Duration:          duration,
			DurationFrequence: durationFrequence,
			Consultation:      consultation,
			Webinar:           webinar,
		}
		return packageRow, sqlError
	}
}

func (r *packageRepository) GetAllPackage(skip, take int, filter string) ([]model.Package, error) {
	var packageList []model.Package
	query := fmt.Sprintf(_getAllPackage, filter, skip, take)

	rows, sqlError := r.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPackage => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id, duration int
			var createdDate time.Time
			var isActive, isDeleted bool
			var modifiedDate sql.NullTime
			var PriceDiscount, modifiedBy, description sql.NullString
			var consultation, webinar, durationFrequence sql.NullInt64
			var types, classCode, pricePackage, code, createdBy string

			sqlError := rows.Scan(
				&id,
				&code,
				&classCode,
				&types,
				&pricePackage,
				&PriceDiscount,
				&description,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&isDeleted,
				&duration,
				&durationFrequence,
				&consultation,
				&webinar,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllPackage => ", sqlError.Error())
			} else {
				packageList = append(
					packageList,
					model.Package{
						ID:                id,
						Code:              code,
						ClassCode:         classCode,
						Type:              types,
						PricePackage:      pricePackage,
						PriceDiscount:     PriceDiscount,
						Description:       description,
						IsActive:          isActive,
						CreatedBy:         createdBy,
						CreatedDate:       createdDate,
						ModifiedBy:        modifiedBy,
						ModifiedDate:      modifiedDate,
						IsDeleted:         isDeleted,
						Duration:          duration,
						DurationFrequence: durationFrequence,
						Consultation:      consultation,
						Webinar:           webinar,
					},
				)
			}
		}
	}
	return packageList, sqlError
}

func (r *packageRepository) GetAllPackageCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(_getAllPackageCount, filter)

	row := r.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPackageCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (r *packageRepository) GetAllBenefit(skip, take int, filter string) ([]model.Benefit, error) {
	var packageList []model.Benefit
	query := fmt.Sprintf(_getAllBenefit, filter, skip, take)

	rows, sqlError := r.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllBenefit => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id, sequence int
			var createdDate time.Time
			var isActive, isDeleted bool
			var modifiedDate sql.NullTime
			var code, value, classCode, createdBy string
			var iconBenefit, modifiedBy sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&classCode,
				&value,
				&iconBenefit,
				&sequence,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&isDeleted,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllBenefit => ", sqlError.Error())
			} else {
				packageList = append(
					packageList,
					model.Benefit{
						ID:           id,
						Code:         code,
						ClassCode:    classCode,
						Description:  value,
						IconBenefit:  iconBenefit,
						Sequence:     sequence,
						IsActive:     isActive,
						CreatedBy:    createdBy,
						CreatedDate:  createdDate,
						ModifiedBy:   modifiedBy,
						ModifiedDate: modifiedDate,
						IsDeleted:    isDeleted,
					},
				)
			}
		}
	}
	return packageList, sqlError
}

func (r *packageRepository) GetAllPackageQuran(filter string) (*[]model.PackageQuran, error) {
	var result []model.PackageQuran
	query := fmt.Sprintf(_getAllPackageQuran, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "packageRepository.GetAllPackageQuran :  error get")
	}

	return &result, nil
}

func (r *packageRepository) GetAllBenefitQuran(filter string) (*[]model.BenefitQuran, error) {
	var result []model.BenefitQuran
	query := fmt.Sprintf(_getAllBenefitQuran, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "packageRepository.GetAllBenefitQuran :  error get")
	}

	return &result, nil
}

func (r *packageRepository) GetAllTermConditionQuran(filter string) (*[]model.TermConditionQuran, error) {
	var result []model.TermConditionQuran
	query := fmt.Sprintf(_getAllTermConditionQuran, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "packageRepository.GetAllTermConditionQuran :  error get")
	}

	return &result, nil
}
