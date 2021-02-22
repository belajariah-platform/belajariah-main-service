package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type userClassHistoryRepository struct {
	db *sqlx.DB
}

type UserClassHistoryRepository interface {
	InsertUserClassHistory(userClassHistory model.UserClass) (bool, error)
}

func InitUserClassHistoryRepository(db *sqlx.DB) UserClassHistoryRepository {
	return &userClassHistoryRepository{
		db,
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
	INSERT INTO transact_user_class
	(
		user_code,
		class_code,
		package_code,
		type_code,
		status_code,
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
			WHERE lower(value)=lower('new class') LIMIT 1),
		(SELECT code 
			FROM master_enum me 
			WHERE lower(value)=lower('start') LIMIT 1),
		$4,
		$5,
		$6,
		$7
	);
	`,
		userClass.UserCode,
		userClass.ClassCode,
		userClass.PackageCode,
		userClass.CreatedBy,
		userClass.CreatedDate,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
	)
	return err
}
