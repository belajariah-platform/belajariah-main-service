package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type exerciseReadingRepository struct {
	db *sqlx.DB
}

type ExerciseReadingRepository interface {
	GetAllExerciseReading(skip, take int, filter string) ([]model.ExerciseReading, error)
	GetExerciseReading(titleCode string) (model.ExerciseReading, error)
	GetAllExerciseReadingCount(filter string) (int, error)
}

func InitExerciseReadingRepository(db *sqlx.DB) ExerciseReadingRepository {
	return &exerciseReadingRepository{
		db,
	}
}

func (exerciseReadingRepository *exerciseReadingRepository) GetAllExerciseReading(skip, take int, filter string) ([]model.ExerciseReading, error) {
	var exerciseReadingList []model.ExerciseReading
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		title_code,
		surat_code,
		ayat_start,
		ayat_end,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM master.master_exercise_reading
	WHERE 
		is_deleted=false AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := exerciseReadingRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllExerciseReading => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var createdDate time.Time
			var isActive, isDeleted bool
			var modifiedBy sql.NullString
			var modifiedDate sql.NullTime
			var titleCode, code, createdBy string
			var id, ayatStart, ayatEnd, suratCode int

			sqlError := rows.Scan(
				&id,
				&code,
				&titleCode,
				&suratCode,
				&ayatStart,
				&ayatEnd,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&isDeleted,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllExerciseReading => ", sqlError.Error())
			} else {
				exerciseReadingList = append(
					exerciseReadingList,
					model.ExerciseReading{
						ID:           id,
						Code:         code,
						TitleCode:    titleCode,
						SuratCode:    suratCode,
						AyatStart:    ayatStart,
						AyatEnd:      ayatEnd,
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
	return exerciseReadingList, sqlError
}

func (exerciseReadingRepository *exerciseReadingRepository) GetExerciseReading(titleCodes string) (model.ExerciseReading, error) {
	var exerciseReadRow model.ExerciseReading
	row := exerciseReadingRepository.db.QueryRow(`
	SELECT
		id,
		code,
		title_code,
		surat_code,
		ayat_start,
		ayat_end,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM master.master_exercise_reading
	WHERE 
		is_deleted=false AND
		is_active=true AND
		title_code=$1
	`, titleCodes)

	var createdDate time.Time
	var isActive, isDeleted bool
	var modifiedBy sql.NullString
	var modifiedDate sql.NullTime
	var titleCode, code, createdBy string
	var id, suratCode, ayatStart, ayatEnd int

	sqlError := row.Scan(
		&id,
		&code,
		&titleCode,
		&suratCode,
		&ayatStart,
		&ayatEnd,
		&isActive,
		&createdBy,
		&createdDate,
		&modifiedBy,
		&modifiedDate,
		&isDeleted,
	)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetExerciseReading => ", sqlError.Error())
		return model.ExerciseReading{}, nil
	} else {
		exerciseReadRow = model.ExerciseReading{
			ID:           id,
			Code:         code,
			TitleCode:    titleCode,
			SuratCode:    suratCode,
			AyatStart:    ayatStart,
			AyatEnd:      ayatEnd,
			IsActive:     isActive,
			CreatedBy:    createdBy,
			CreatedDate:  createdDate,
			ModifiedBy:   modifiedBy,
			ModifiedDate: modifiedDate,
			IsDeleted:    isDeleted,
		}
		return exerciseReadRow, sqlError
	}
}

func (exerciseReadingRepository *exerciseReadingRepository) GetAllExerciseReadingCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		master.master_exercise_reading  
	WHERE 
		is_deleted=false AND
		is_active=true
	%s
	`, filter)

	row := exerciseReadingRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllExerciseReadingCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}
