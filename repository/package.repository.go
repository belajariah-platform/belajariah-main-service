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
		deleted_by,
		deleted_date,
		duration
	FROM master_class_package
	WHERE 
		deleted_by IS NULL AND
		is_active=true AND
		code=$1
	`, codes)

	var isActive bool
	var id, duration int
	var createdDate time.Time
	var modifiedDate, deletedDate sql.NullTime
	var PriceDiscount, modifiedBy, deletedBy sql.NullString
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
		&deletedBy,
		&deletedDate,
		&duration,
	)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetPackage => ", sqlError)
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
			DeletedBy:     deletedBy,
			DeletedDate:   deletedDate,
			Duration:      duration,
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
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date,
		duration
	FROM master_class_package
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := packageRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPackage => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var isActive bool
			var id, duration int
			var createdDate time.Time
			var modifiedDate, deletedDate sql.NullTime
			var PriceDiscount, modifiedBy, deletedBy sql.NullString
			var types, classCode, pricePackage, code, createdBy string

			sqlError := rows.Scan(
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
				&deletedBy,
				&deletedDate,
				&duration,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllPackage => ", sqlError)
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
						IsActive:      isActive,
						CreatedBy:     createdBy,
						CreatedDate:   createdDate,
						ModifiedBy:    modifiedBy,
						ModifiedDate:  modifiedDate,
						DeletedBy:     deletedBy,
						DeletedDate:   deletedDate,
						Duration:      duration,
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
		master_class_package  
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	`, filter)

	row := packageRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPackageCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}
