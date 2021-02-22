package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type exerciseRepository struct {
	db *sqlx.DB
}

type ExerciseRepository interface {
	GetAllExercise(skip, take int, filter string) ([]model.Exercise, error)
	GetAllExerciseCount(filter string) (int, error)
}

func InitExerciseRepository(db *sqlx.DB) ExerciseRepository {
	return &exerciseRepository{
		db,
	}
}

func (exerciseRepository *exerciseRepository) GetAllExercise(skip, take int, filter string) ([]model.Exercise, error) {
	var exerciseList []model.Exercise
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		subtitle_code,
		image_code,
		exercise_image,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM v_m_exercise
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := exerciseRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllExercise => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var isActive bool
			var id, imageCode int
			var createdDate time.Time
			var modifiedBy, deletedBy sql.NullString
			var modifiedDate, deletedDate sql.NullTime
			var exerciseImage, subtitleCode, code, createdBy string

			sqlError := rows.Scan(
				&id,
				&code,
				&subtitleCode,
				&imageCode,
				&exerciseImage,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllExercise => ", sqlError)
			} else {
				exerciseList = append(
					exerciseList,
					model.Exercise{
						ID:            id,
						Code:          code,
						SubtitleCode:  subtitleCode,
						ImageCode:     imageCode,
						ExerciseImage: exerciseImage,
						IsActive:      isActive,
						CreatedBy:     createdBy,
						CreatedDate:   createdDate,
						ModifiedBy:    modifiedBy,
						ModifiedDate:  modifiedDate,
						DeletedBy:     deletedBy,
						DeletedDate:   deletedDate,
					},
				)
			}
		}
	}
	return exerciseList, sqlError
}

func (exerciseRepository *exerciseRepository) GetAllExerciseCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		master_exercise  
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	`, filter)

	row := exerciseRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllExerciseCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}
