package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type learningRepository struct {
	db *sqlx.DB
}

type LearningRepository interface {
	GetAllLearning(skip, take int, filter string) ([]model.Learning, error)
	GetAllSubLearning(titleCode string) ([]model.SubLearning, error)
	GetAllSubLearningCount(filter string) (int, error)
}

func InitLearningRepository(db *sqlx.DB) LearningRepository {
	return &learningRepository{
		db,
	}
}

func (learningRepository *learningRepository) GetAllLearning(skip, take int, filter string) ([]model.Learning, error) {
	var learningList []model.Learning
	query := fmt.Sprintf(`
	SELECT
		id,	
		code,
		class_code,
		title,
		document,
		sequence,	
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM v_m_class_learning
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := learningRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllLearning => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var isActive bool
			var id int
			var createdDate time.Time
			var sequenced sql.NullInt64
			var modifiedDate, deletedDate sql.NullTime
			var document, modifiedBy, deletedBy sql.NullString
			var classCode, title, code, createdBy string

			sqlError := rows.Scan(
				&id,
				&code,
				&classCode,
				&title,
				&document,
				&sequenced,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllLearning => ", sqlError)
			} else {
				learningList = append(
					learningList,
					model.Learning{
						ID:           id,
						Code:         code,
						ClassCode:    classCode,
						Title:        title,
						Document:     document,
						Sequence:     sequenced,
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
	return learningList, sqlError
}

func (learningRepository *learningRepository) GetAllSubLearning(titleCode string) ([]model.SubLearning, error) {
	var learningList []model.SubLearning
	query := fmt.Sprintf(`
	SELECT
		id,	
		code,
		title_code,
		sub_title,
		video_duration,
		video,
		document,
		exercise_image,
		sequence,	
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM v_m_class_sub_learning
	WHERE 
		deleted_by IS NULL AND
		is_active=true AND
		title_code='%s'
	`, titleCode)

	rows, sqlError := learningRepository.db.Query(query)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllSubLearning => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var isActive bool
			var createdDate time.Time
			var sequenced sql.NullInt64
			var videoDuration sql.NullFloat64
			var titleCode, code, createdBy string
			var modifiedDate, deletedDate sql.NullTime
			var exerciseImage, subTitle, document, video, modifiedBy, deletedBy sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&titleCode,
				&subTitle,
				&videoDuration,
				&video,
				&document,
				&exerciseImage,
				&sequenced,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
			)
			fmt.Println(sqlError)
			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllSubLearning => ", sqlError)
			} else {
				learningList = append(
					learningList,
					model.SubLearning{
						ID:            id,
						Code:          code,
						TitleCode:     titleCode,
						SubTitle:      subTitle,
						VideoDuration: videoDuration,
						Video:         video,
						Document:      document,
						ExerciseImage: exerciseImage,
						Sequence:      sequenced,
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
	return learningList, sqlError
}

func (learningRepository *learningRepository) GetAllSubLearningCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		v_m_class_sub_learning  
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	`, filter)

	row := learningRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllSubLearningCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}
