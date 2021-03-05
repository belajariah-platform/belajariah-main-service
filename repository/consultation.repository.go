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
	GetConsultation(filter string) (model.Consultation, error)
	GetAllConsultationCount(filter, filterUser string) (int, error)
	GetAllConsultation(skip, take int, sort, search, filter, filterUser string) ([]model.Consultation, error)
	GetAllConsultationLimit(skip, take int, filter, filterUser string) ([]model.Consultation, error)

	ReadConsultation(consultation model.Consultation) (bool, error)
	InsertConsultation(consultation model.Consultation) (bool, error)
	UpdateConsultation(consultation model.Consultation, status string) (bool, error)
	ConfirmConsultation(consultation model.Consultation, status string) (bool, error)

	CheckConsultationSpam() ([]model.Consultation, error)
	CheckAllConsultationExpired() ([]model.Consultation, error)
}

func InitConsultationRepository(db *sqlx.DB) ConsultationRepository {
	return &consultationRepository{
		db,
	}
}

func (consultationRepository *consultationRepository) GetAllConsultationLimit(skip, take int, filter, filterUser string) ([]model.Consultation, error) {
	var consultationList []model.Consultation
	query := fmt.Sprintf(`
	SELECT
		DISTINCT ON (user_code) user_code,
		id,
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
	ORDER BY user_code, id DESC
	OFFSET %d
	LIMIT %d
	`, filterUser, filter, skip, take)

	rows, sqlError := consultationRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllConsultationLimit => ", sqlError)
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
				&userCode,
				&id,
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
				utils.PushLogf("SQL error on GetAllConsultationLimit => ", sqlError)
			} else {
				consultationList = append(
					consultationList,
					model.Consultation{
						UserCode:          userCode,
						ID:                id,
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

func (consultationRepository *consultationRepository) GetAllConsultation(skip, take int, sort, search, filter, filterUser string) ([]model.Consultation, error) {
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
		%s %s %s %s
	OFFSET %d
	LIMIT %d
	`, filterUser, filter, search, sort, skip, take)

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

func (consultationRepository *consultationRepository) GetConsultation(filter string) (model.Consultation, error) {
	var consultationRow model.Consultation
	query := fmt.Sprintf(`
	SELECT
		id,
		user_code,
		class_code,
		status_code,
		status,
		taken_code,
		taken_name,
		is_action_taken,
		expired_date
	FROM v_t_consultation
	WHERE 
		deleted_by IS NULL
		%s
	ORDER BY id desc
	`, filter)
	row := consultationRepository.db.QueryRow(query)

	var id, userCode int
	var takenCode sql.NullInt64
	var expiredDate sql.NullTime
	var takenName sql.NullString
	var isActionTaken sql.NullBool
	var statusCode, status, classCode string

	sqlError := row.Scan(
		&id,
		&userCode,
		&classCode,
		&statusCode,
		&status,
		&takenCode,
		&takenName,
		&isActionTaken,
		&expiredDate,
	)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetConsultation => ", sqlError)
		return model.Consultation{}, nil
	} else {
		consultationRow = model.Consultation{
			ID:            id,
			UserCode:      userCode,
			ClassCode:     classCode,
			StatusCode:    statusCode,
			Status:        status,
			TakenCode:     takenCode,
			TakenName:     takenName,
			IsActionTaken: isActionTaken,
			ExpiredDate:   expiredDate,
		}
		return consultationRow, sqlError
	}
}

func (consultationRepository *consultationRepository) InsertConsultation(consultation model.Consultation) (bool, error) {
	var err error
	var result bool

	tx, errTx := consultationRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in InsertConsultation", errTx)
	} else {
		err = insertConsultation(tx, consultation)
		if err != nil {
			utils.PushLogf("err in user-class---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to InsertConsultation", err)
	}

	return result, err
}

func insertConsultation(tx *sql.Tx, consultation model.Consultation) error {
	_, err := tx.Exec(`
	INSERT INTO transact_consultation
	(
		user_code,
		class_code,
		recording_code,
		duration,
		status_code,
		description,
		taken_code,
		is_action_taken,
		created_by,
		created_date,
		modified_by,
		modified_date,
		expired_date
	)
	VALUES(
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		$12,
		$13
	);
	`,
		consultation.UserCode,
		consultation.ClassCode,
		consultation.RecordingCode.Int64,
		consultation.RecordingDuration.Int64,
		consultation.StatusCode,
		consultation.Description.String,
		consultation.TakenCode.Int64,
		consultation.IsActionTaken.Bool,
		consultation.CreatedBy,
		consultation.CreatedDate,
		consultation.ModifiedBy.String,
		consultation.ModifiedDate.Time,
		utils.CurrentDateStringCustom(consultation.ExpiredDate.Time),
	)
	return err
}

func (consultationRepository *consultationRepository) ReadConsultation(consultation model.Consultation) (bool, error) {
	var err error
	var result bool

	tx, errTx := consultationRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in ReadConsultation", errTx)
	} else {
		err = readConsultation(tx, consultation)
		if err != nil {
			utils.PushLogf("err in user-class---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to ReadConsultation", err)
	}

	return result, err
}

func readConsultation(tx *sql.Tx, consultation model.Consultation) error {
	_, err := tx.Exec(`
	UPDATE
		transact_consultation
	SET
		is_read=true,
	WHERE
		id=$1
	);
	`,
		consultation.ID,
	)
	return err
}

func (consultationRepository *consultationRepository) UpdateConsultation(consultation model.Consultation, status string) (bool, error) {
	var err error
	var result bool

	tx, errTx := consultationRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in UpdateConsultation", errTx)
	} else {
		err = updateConsultation(tx, consultation, status)
		if err != nil {
			utils.PushLogf("err in user-class---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to UpdateConsultation", err)
	}
	return result, err
}

func updateConsultation(tx *sql.Tx, consultation model.Consultation, status string) error {
	_, err := tx.Exec(`
	UPDATE
		transact_consultation
	SET
		status_code=$1,	
		taken_code=$2,
		is_action_taken=$3,
		modified_by=$4,
		modified_date=$5
	WHERE
		user_code=$6 AND 
		class_code=$7 AND
		expired_date=$8 AND
		status_code=$9
	`,
		consultation.StatusCode,
		consultation.TakenCode.Int64,
		consultation.IsActionTaken.Bool,
		consultation.ModifiedBy.String,
		consultation.ModifiedDate.Time,
		consultation.UserCode,
		consultation.ClassCode,
		utils.CurrentDateStringCustom(consultation.ExpiredDate.Time),
		status,
	)
	return err
}

func (consultationRepository *consultationRepository) ConfirmConsultation(consultation model.Consultation, status string) (bool, error) {
	var err error
	var result bool

	tx, errTx := consultationRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in ConfirmConsultation", errTx)
	} else {
		err = confirmConsultation(tx, consultation, status)
		if err != nil {
			utils.PushLogf("err in user-class---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to ConfirmConsultation", err)
	}

	return result, err
}

func confirmConsultation(tx *sql.Tx, consultation model.Consultation, status string) error {
	_, err := tx.Exec(`
	UPDATE
		transact_consultation
	SET
		status_code=$1,
		modified_by=$2,
		modified_date=$3
	WHERE
		user_code=$4 AND 
		class_code=$5 AND
		expired_date=$6 AND
		status_code=$7
	`,
		consultation.StatusCode,
		consultation.ModifiedBy.String,
		consultation.ModifiedDate.Time,
		consultation.UserCode,
		consultation.ClassCode,
		utils.CurrentDateStringCustom(consultation.ExpiredDate.Time),
		status,
	)
	return err
}

func (consultationRepository *consultationRepository) CheckConsultationSpam() ([]model.Consultation, error) {
	var consultationList []model.Consultation

	rows, sqlError := consultationRepository.db.Query(`
	SELECT
		id,
		user_code
	FROM 
		v_t_consultation 
	ORDER BY id DESC 
	LIMIT 2;
	`)

	if sqlError != nil {
		utils.PushLogf("SQL error on CheckConsultationSpam => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var id, userCode int

			sqlError := rows.Scan(
				&id,
				&userCode,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on CheckConsultationSpam => ", sqlError)
			} else {
				consultationList = append(consultationList, model.Consultation{
					ID:       id,
					UserCode: userCode,
				})
			}
		}
	}
	return consultationList, sqlError
}

func (consultationRepository *consultationRepository) CheckAllConsultationExpired() ([]model.Consultation, error) {
	var consultationList []model.Consultation

	rows, sqlError := consultationRepository.db.Query(`
	SELECT
		id,
		user_code,
		class_code,
		expired_date,
		status_code
	FROM v_t_consultation
	WHERE  
		DATE_PART('day', now()::timestamp - modified_date::timestamp) * 24 * 60 * 60 + 
		DATE_PART('hour', now()::timestamp - modified_date::timestamp) * 60 * 60 +
		DATE_PART('minute', now()::timestamp - modified_date::timestamp) * 60 +
		DATE_PART('second', now()::timestamp - modified_date::timestamp) > 600 AND 
		status IN ('Waiting for Response')
	ORDER BY id ASC 
	LIMIT 1;
	`)

	if sqlError != nil {
		utils.PushLogf("SQL error on CheckAllConsultationExpired => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var id, userCode int
			var expiredDate sql.NullTime
			var classCode, statusCode string

			sqlError := rows.Scan(
				&id,
				&userCode,
				&classCode,
				&expiredDate,
				&statusCode,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on CheckAllConsultationExpired => ", sqlError)
			} else {
				consultationList = append(consultationList, model.Consultation{
					ID:          id,
					UserCode:    userCode,
					ClassCode:   classCode,
					ExpiredDate: expiredDate,
					StatusCode:  statusCode,
				})
			}
		}
	}
	return consultationList, sqlError
}
