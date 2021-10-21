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
		image_exercise,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM master.master_exercise_writing
	WHERE 
		is_deleted=false AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := exerciseRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllExerciseWriting => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var createdDate time.Time
			var isActive, isDeleted bool
			var modifiedDate sql.NullTime
			var subtitleCode, code, createdBy string
			var modifiedBy, imageExercise sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&subtitleCode,
				&imageExercise,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&isDeleted,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllExerciseWriting => ", sqlError.Error())
			} else {
				exerciseList = append(
					exerciseList,
					model.Exercise{
						ID:            id,
						Code:          code,
						SubtitleCode:  subtitleCode,
						ImageExercise: imageExercise,
						IsActive:      isActive,
						CreatedBy:     createdBy,
						CreatedDate:   createdDate,
						ModifiedBy:    modifiedBy,
						ModifiedDate:  modifiedDate,
						IsDeleted:     isDeleted,
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
		master.master_exercise_writing  
	WHERE 
		is_deleted=false AND
		is_active=true
	%s
	`, filter)

	row := exerciseRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllExerciseWritingCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}
