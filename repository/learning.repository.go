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
		document_path,
		document_name,
		sequence,	
		is_exercise,
		is_direct_learning,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM master.v_m_learning
	WHERE 
		is_deleted=false AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := learningRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllLearning => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id, sequenced int
			var isExercise bool
			var createdDate time.Time
			var modifiedDate sql.NullTime
			var classCode, title, code, createdBy string
			var isActive, isDirectLearning, isDeleted bool
			var documentPath, documentName, modifiedBy sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&classCode,
				&title,
				&documentPath,
				&documentName,
				&sequenced,
				&isExercise,
				&isDirectLearning,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&isDeleted,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllLearning => ", sqlError.Error())
			} else {
				learningList = append(
					learningList,
					model.Learning{
						ID:               id,
						Code:             code,
						ClassCode:        classCode,
						Title:            title,
						DocumentPath:     documentPath,
						DocumentName:     documentName,
						Sequence:         sequenced,
						IsExercise:       isExercise,
						IsDirectLearning: isDirectLearning,
						IsActive:         isActive,
						CreatedBy:        createdBy,
						CreatedDate:      createdDate,
						ModifiedBy:       modifiedBy,
						ModifiedDate:     modifiedDate,
						IsDeleted:        isDeleted,
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
		poster,
		sequence,	
		is_exercise,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM master.v_m_sublearning
	WHERE 
		is_deleted=false AND
		is_active=true AND
		title_code='%s'
	ORDER BY id ASC
	`, titleCode)

	rows, sqlError := learningRepository.db.Query(query)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllSubLearning => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var id, sequenced int
			var createdDate time.Time
			var isActive, isDeleted, isExercise bool
			var modifiedDate sql.NullTime
			var videoDuration sql.NullFloat64
			var titleCode, code, createdBy string
			var poster, subTitle, document, video, modifiedBy sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&titleCode,
				&subTitle,
				&videoDuration,
				&video,
				&document,
				&poster,
				&sequenced,
				&isExercise,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&isDeleted,
			)
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
						Poster:        poster,
						Sequence:      sequenced,
						IsExercise:    isExercise,
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
	return learningList, sqlError
}
