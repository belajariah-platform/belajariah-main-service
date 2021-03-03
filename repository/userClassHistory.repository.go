package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type userClassHistoryRepository struct {
	db *sqlx.DB
}

type UserClassHistoryRepository interface {
	GetUserClassHistory(filter string) (model.UserClass, error)
	InsertUserClassHistory(userClassHistory model.UserClass) (bool, error)
	DeleteUserClassHistory(userClass model.UserClass) (bool, error)
}

func InitUserClassHistoryRepository(db *sqlx.DB) UserClassHistoryRepository {
	return &userClassHistoryRepository{
		db,
	}
}

func (userClassHistoryRepository *userClassHistoryRepository) GetUserClassHistory(filter string) (model.UserClass, error) {
	var userClassRow model.UserClass
	query := fmt.Sprintf(`
	SELECT
		id,
		user_code,
		class_code,
		package_code,
		type_code,
		status_code,
		expired_date
	FROM 
		transact_user_class_history
	WHERE 
		deleted_by IS NULL
		%s
	`, filter)
	row := userClassHistoryRepository.db.QueryRow(query)

	var id, userCode int
	var expiredDate time.Time
	var typeCode, packageCode, statusCode, classCode string

	sqlError := row.Scan(
		&id,
		&userCode,
		&classCode,
		&typeCode,
		&packageCode,
		&statusCode,
		&expiredDate,
	)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetUserClassHistory => ", sqlError)
		return model.UserClass{}, nil
	} else {
		userClassRow = model.UserClass{
			ID:          id,
			UserCode:    userCode,
			ClassCode:   classCode,
			TypeCode:    typeCode,
			PackageCode: packageCode,
			StatusCode:  statusCode,
			ExpiredDate: expiredDate,
		}
		return userClassRow, sqlError
	}
}

func (userClassHistoryRepository *userClassHistoryRepository) InsertUserClassHistory(userClass model.UserClass) (bool, error) {
	var err error
	var result bool

	tx, errTx := userClassHistoryRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in InsertUserClassHistory", errTx)
	} else {
		err = insertUserClassHistory(tx, userClass)
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
		utils.PushLogf("failed to InsertUserClassHistory", err)
	}

	return result, err
}

func insertUserClassHistory(tx *sql.Tx, userClass model.UserClass) error {
	_, err := tx.Exec(`
	INSERT INTO transact_user_class_history
	(
		user_code,
		class_code,
		package_code,
		type_code,
		status_code,
		expired_date,
		created_by,
		created_date,
		modified_by,
		modified_date
	)
	VALUES(	
		$1,
		$2,
		$3,
		(SELECT code 
			FROM master_enum me 
			WHERE lower(value)=lower($4) LIMIT 1),
		(SELECT code 
			FROM master_enum me 
			WHERE lower(value)=lower('start') LIMIT 1),
		$5,
		$6,
		$7,
		$8,
		$9
	);
	`,
		userClass.UserCode,
		userClass.ClassCode,
		userClass.PackageCode,
		userClass.TypeCode,
		userClass.ExpiredDate,
		userClass.CreatedBy,
		userClass.CreatedDate,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
	)
	return err
}

func (userClassHistoryRepository *userClassHistoryRepository) DeleteUserClassHistory(userClass model.UserClass) (bool, error) {
	var err error
	var result bool

	tx, errTx := userClassHistoryRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in DeleteUserClassHistory", errTx)
	} else {
		err = deleteUserClassHistory(tx, userClass)
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
		utils.PushLogf("failed to DeleteUserClassHistory", err)
	}

	return result, err
}

func deleteUserClassHistory(tx *sql.Tx, userClass model.UserClass) error {
	_, err := tx.Exec(`
	UPDATE
		transact_user_class_history
	 SET
		deleted_by=$1,
		deleted_date=$2
 	 WHERE
 		user_code=$3 AND
		class_code=$4 AND
		expired_date=$5
	`,
		userClass.DeletedBy.String,
		userClass.DeletedDate.Time,
		userClass.UserCode,
		userClass.ClassCode,
		userClass.ExpiredDate,
	)
	return err
}
