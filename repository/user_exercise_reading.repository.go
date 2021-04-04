package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type userExerciseReadingRepository struct {
	db *sqlx.DB
}

type UserExerciseReadingRepository interface {
	GetAllUserExerciseReadingCount(filter string) (int, error)
	InsertUserExerciseReading(userExercise model.UserExerciseReading) (bool, error)
}

func InitUserExerciseReadingRepository(db *sqlx.DB) UserExerciseReadingRepository {
	return &userExerciseReadingRepository{
		db,
	}
}

func (userExerciseReadingRepository *userExerciseReadingRepository) GetAllUserExerciseReadingCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		transact_exercise_reading  
	WHERE 
		deleted_by IS NULL
		%s
	`, filter)

	row := userExerciseReadingRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllUserExerciseReadingCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}

func (userExerciseReadingRepository *userExerciseReadingRepository) InsertUserExerciseReading(userExercise model.UserExerciseReading) (bool, error) {
	var err error
	var result bool

	tx, errTx := userExerciseReadingRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in InsertUserExerciseReading", errTx)
	} else {
		err = insertUserExerciseReading(tx, userExercise)
		if err != nil {
			utils.PushLogf("err in user-class-history---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to InsertUserExerciseReading", err)
	}

	return result, err
}

func insertUserExerciseReading(tx *sql.Tx, userExercise model.UserExerciseReading) error {
	_, err := tx.Exec(`
	INSERT INTO transact_exercise_reading
	(
		user_code,
		class_code,
		recording_code,
		duration,
		expired_date,
		title_code,
		created_by,
		created_date,
		modified_by,
		modified_date
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
		$10
	);
	`,
		userExercise.UserCode,
		userExercise.ClassCode,
		userExercise.RecordingCode,
		userExercise.Duration,
		utils.CurrentDateStringCustom(userExercise.ExpiredDate),
		userExercise.TitleCode.String,
		userExercise.CreatedBy,
		userExercise.CreatedDate,
		userExercise.ModifiedBy.String,
		userExercise.ModifiedDate.Time,
	)
	return err
}
