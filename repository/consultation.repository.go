package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type consultationRepository struct {
	db *sqlx.DB
}

type ConsultationRepository interface {
	GetAllConsultation(skip, take int, filter, filterUser string) ([]model.Consultation, error)
	GetAllConsultationCount(filter, filterUser string) (int, error)
}

func InitConsultationRepository(db *sqlx.DB) ConsultationRepository {
	return &consultationRepository{
		db,
	}
}

func (consultationRepository *consultationRepository) GetAllConsultation(skip, take int, filter, filterUser string) ([]model.Consultation, error) {
	var consultationList []model.Consultation
	query := fmt.Sprintf(`
	SELECT
		id,
		user_code,
		user_name,
		class_code,
		class_name,
		recording_code,
		recording_path,
		recording_name,
		recording_duration,
		status_code,
		status,
		description,
		taken_code,
		taken_name,
		is_play,
		is_read,
		is_action_taken,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM v_t_consultation
	WHERE 
		deleted_by IS NULL
		%s
	%s
	OFFSET %d
	LIMIT %d
	`, filterUser, filter, skip, take)

	rows, sqlError := consultationRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllConsultation => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var isActive bool
			var id, userCode int
			var isPlay, isRead, isActionTaken sql.NullBool
			var createdDate time.Time
			var modifiedDate, deletedDate sql.NullTime
			var recordingCode, recordingDuration, takenCode sql.NullInt64
			var recordingPath, recordingName, description, takenName, modifiedBy, deletedBy sql.NullString
			var userName, statusCode, status, classCode, className, createdBy string

			sqlError := rows.Scan(
				&id,
				&userCode,
				&userName,
				&classCode,
				&className,
				&recordingCode,
				&recordingPath,
				&recordingName,
				&recordingDuration,
				&statusCode,
				&status,
				&description,
				&takenCode,
				&takenName,
				&isPlay,
				&isRead,
				&isActionTaken,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllConsultation => ", sqlError)
			} else {
				consultationList = append(
					consultationList,
					model.Consultation{
						ID:                id,
						UserCode:          userCode,
						UserName:          userName,
						ClassCode:         classCode,
						ClassName:         className,
						RecordingCode:     recordingCode,
						RecordingPath:     recordingPath,
						RecordingName:     recordingName,
						RecordingDuration: recordingDuration,
						StatusCode:        statusCode,
						Status:            status,
						Description:       description,
						TakenCode:         takenCode,
						TakenName:         takenName,
						IsPlay:            isPlay,
						IsRead:            isRead,
						IsActionTaken:     isActionTaken,
						IsActive:          isActive,
						CreatedBy:         createdBy,
						CreatedDate:       createdDate,
						ModifiedBy:        modifiedBy,
						ModifiedDate:      modifiedDate,
						DeletedBy:         deletedBy,
						DeletedDate:       deletedDate,
					},
				)
			}
		}
	}
	return consultationList, sqlError
}

func (consultationRepository *consultationRepository) GetAllConsultationCount(filter, filterUser string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		v_t_consultation  
	WHERE 
		deleted_by IS NULL 
		%s
	%s
	`, filterUser, filter)

	row := consultationRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllConsultationCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}
