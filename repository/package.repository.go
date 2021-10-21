package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type packageRepository struct {
	db *sqlx.DB
}

type PackageRepository interface {
	GetPackage(code string) (model.Package, error)
	GetAllPackage(skip, take int, filter string) ([]model.Package, error)
	GetAllPackageCount(filter string) (int, error)
	GetAllBenefit(skip, take int, filter string) ([]model.Benefit, error)
}

func InitPackageRepository(db *sqlx.DB) PackageRepository {
	return &packageRepository{
		db,
	}
}

func (packageRepository *packageRepository) GetPackage(codes string) (model.Package, error) {
	var packageRow model.Package
	row := packageRepository.db.QueryRow(`
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
		consultation,
		webinar
	FROM master.master_package
	WHERE 
		is_deleted = false AND
		is_active=true AND
		code=$1
	`, codes)

	var id, duration int
	var createdDate time.Time
	var isActive, isDeleted bool
	var consultation, webinar sql.NullInt64
	var modifiedDate sql.NullTime
	var PriceDiscount, modifiedBy sql.NullString
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
		&consultation,
		&webinar,
	)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetPackage => ", sqlError.Error())
		return model.Package{}, nil
	} else {
		packageRow = model.Package{
			ID:            id,
			Code:          code,
			ClassCode:     classCode,
			Type:          types,
			PricePackage:  pricePackage,
			PriceDiscount: PriceDiscount,
			IsActive:      isActive,
			CreatedBy:     createdBy,
			CreatedDate:   createdDate,
			ModifiedBy:    modifiedBy,
			ModifiedDate:  modifiedDate,
			IsDeleted:     isDeleted,
			Duration:      duration,
			Consultation:  consultation,
			Webinar:       webinar,
		}
		return packageRow, sqlError
	}
}

func (packageRepository *packageRepository) GetAllPackage(skip, take int, filter string) ([]model.Package, error) {
	var packageList []model.Package
	query := fmt.Sprintf(`
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
		consultation,
		webinar
	FROM master.master_package
	WHERE 
		is_deleted = false AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := packageRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPackage => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id, duration int
			var createdDate time.Time
			var isActive, isDeleted bool
			var modifiedDate sql.NullTime
			var consultation, webinar sql.NullInt64
			var PriceDiscount, modifiedBy, description sql.NullString
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
				&consultation,
				&webinar,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllPackage => ", sqlError.Error())
			} else {
				packageList = append(
					packageList,
					model.Package{
						ID:            id,
						Code:          code,
						ClassCode:     classCode,
						Type:          types,
						PricePackage:  pricePackage,
						PriceDiscount: PriceDiscount,
						Description:   description,
						IsActive:      isActive,
						CreatedBy:     createdBy,
						CreatedDate:   createdDate,
						ModifiedBy:    modifiedBy,
						ModifiedDate:  modifiedDate,
						IsDeleted:     isDeleted,
						Duration:      duration,
						Consultation:  consultation,
						Webinar:       webinar,
					},
				)
			}
		}
	}
	return packageList, sqlError
}

func (packageRepository *packageRepository) GetAllPackageCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		master.master_package  
	WHERE 
		is_deleted=false AND
		is_active=true
	%s
	`, filter)

	row := packageRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPackageCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (packageRepository *packageRepository) GetAllBenefit(skip, take int, filter string) ([]model.Benefit, error) {
	var packageList []model.Benefit
	query := fmt.Sprintf(`
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
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := packageRepository.db.Query(query)

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
