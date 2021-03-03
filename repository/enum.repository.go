package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type enumRepository struct {
	db *sqlx.DB
}

type EnumRepository interface {
	GetAllEnum(skip, take int, filter string) ([]model.Enum, error)
	GetAllEnumCount(filter string) (int, error)
	GetEnum(value string) (model.Enum, error)
	GetEnumSplit(value string) (model.Enum, error)
}

func InitEnumRepository(db *sqlx.DB) EnumRepository {
	return &enumRepository{
		db,
	}
}

func (enumRepository *enumRepository) GetAllEnum(skip, take int, filter string) ([]model.Enum, error) {
	var enumList []model.Enum
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		type,
		value,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM master_enum
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := enumRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllEnum => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var isActive bool
			var id int
			var createdDate time.Time
			var modifiedDate, deletedDate sql.NullTime
			var modifiedBy, deletedBy sql.NullString
			var types, value, code, createdBy string

			sqlError := rows.Scan(
				&id,
				&code,
				&types,
				&value,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllEnum => ", sqlError)
			} else {
				enumList = append(
					enumList,
					model.Enum{
						ID:           id,
						Code:         code,
						Type:         types,
						Value:        value,
						IsActive:     isActive,
						CreatedBy:    createdBy,
						CreatedDate:  createdDate,
						ModifiedBy:   modifiedBy,
						ModifiedDate: modifiedDate,
						DeletedBy:    deletedBy,
						DeletedDate:  deletedDate,
					},
				)
			}
		}
	}
	return enumList, sqlError
}

func (enumRepository *enumRepository) GetAllEnumCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		master_enum  
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	`, filter)

	row := enumRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllEnumCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}

func (enumRepository *enumRepository) GetEnum(values string) (model.Enum, error) {
	var enumRow model.Enum
	row := enumRepository.db.QueryRow(`
	SELECT
		id,
		code,
		type,
		value
	FROM master_enum
	WHERE 
		deleted_by IS NULL AND
		is_active=true AND
		value=$1
	`, values)

	var id int
	var types, value, code string

	sqlError := row.Scan(
		&id,
		&code,
		&types,
		&value,
	)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetEnum => ", sqlError)
		return model.Enum{}, nil
	} else {
		enumRow = model.Enum{
			ID:    id,
			Code:  code,
			Type:  types,
			Value: value,
		}
		return enumRow, sqlError
	}
}

func (enumRepository *enumRepository) GetEnumSplit(values string) (model.Enum, error) {
	var enumRow model.Enum
	row := enumRepository.db.QueryRow(`
	SELECT
		id,
		code,
		type,
		value
	FROM 
		master_enum
	WHERE 
		deleted_by IS NULL AND
		is_active=true AND
		split_part(value, '|', 1)=$1
	`, values)

	var id int
	var types, value, code string

	sqlError := row.Scan(
		&id,
		&code,
		&types,
		&value,
	)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetEnumSplit => ", sqlError)
		return model.Enum{}, nil
	} else {
		enumRow = model.Enum{
			ID:    id,
			Code:  code,
			Type:  types,
			Value: value,
		}
		return enumRow, sqlError
	}
}
