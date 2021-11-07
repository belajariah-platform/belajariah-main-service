package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	_getUserClassHistory = `
		SELECT
			id,
			code,
			user_class_code,
			package_code,
			promo_code,
			price,
			start_date,
			expired_date
		FROM 
			log.user_class_history
		WHERE 
			is_deleted=false
		%s
	`
	_insertUserClassHistory = `
		INSERT INTO log.user_class_history
		(
			user_class_code,
			package_code,
			promo_code,
			payment_method_code,
			price,
			start_date,
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
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11
		) returning code
	`
	_deleteUserClassHistory = `
		UPDATE
			log.user_class_history
		SET
			modified_by=$1,
			modified_date=$2,
			is_active=false,
			is_deleted=true
		WHERE
			user_class_code=$3 AND
			expired_date=$4
	`
)

type userClassHistoryRepository struct {
	db *sqlx.DB
}

type UserClassHistoryRepository interface {
	GetUserClassHistory(filter string) (model.UserClassHistory, error)
	DeleteUserClassHistory(userClass model.UserClassHistory) (bool, error)
	InsertUserClassHistory(userClassHistory model.UserClassHistory) (model.UserClassHistory, bool, error)
}

func InitUserClassHistoryRepository(db *sqlx.DB) UserClassHistoryRepository {
	return &userClassHistoryRepository{
		db,
	}
}

func (userClassHistoryRepository *userClassHistoryRepository) GetUserClassHistory(filter string) (model.UserClassHistory, error) {
	var userClassRow model.UserClassHistory
	query := fmt.Sprintf(_getUserClassHistory, filter)
	row := userClassHistoryRepository.db.QueryRow(query)

	var id, price int
	var expiredDate, startDate sql.NullTime
	var promoCode, packageCode, userClassCode, code string

	sqlError := row.Scan(
		&id,
		&code,
		&userClassCode,
		&packageCode,
		&promoCode,
		&price,
		&startDate,
		&expiredDate,
	)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetUserClassHistory => ", sqlError)
		return model.UserClassHistory{}, nil
	} else {
		userClassRow = model.UserClassHistory{
			ID:            id,
			Code:          code,
			UserClassCode: userClassCode,
			PackageCode:   packageCode,
			PromoCode:     promoCode,
			Price:         price,
			StartDate:     startDate,
			ExpiredDate:   expiredDate,
		}
		return userClassRow, sqlError
	}
}

func (r *userClassHistoryRepository) InsertUserClassHistory(data model.UserClassHistory) (model.UserClassHistory, bool, error) {
	var code string
	var history model.UserClassHistory

	tx, err := r.db.Beginx()
	if err != nil {
		return model.UserClassHistory{}, false, errors.New("userClassHistoryRepository: InsertUserClassHistory: error begin transaction")
	}

	err = tx.QueryRow(_insertUserClassHistory,
		data.UserClassCode,
		data.PackageCode,
		data.PromoCode,
		data.PaymentMethodCode,
		data.Price,
		data.StartDate.Time,
		data.ExpiredDate.Time,
		data.CreatedBy,
		data.CreatedDate,
		data.ModifiedBy.String,
		data.ModifiedDate.Time,
	).Scan(&code)

	if err != nil {
		tx.Rollback()
		return model.UserClassHistory{}, false, utils.WrapError(err, "userClassHistoryRepository: InsertUserClassHistory: error insert")
	}

	history = model.UserClassHistory{Code: code}

	tx.Commit()
	return history, err == nil, nil
}

func (r *userClassHistoryRepository) DeleteUserClassHistory(data model.UserClassHistory) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("userClassHistoryRepository: DeleteUserClassHistory: error begin transaction")
	}

	_, err = tx.Exec(_deleteUserClassHistory,
		data.UserClassCode,
		data.ModifiedBy.String,
		data.ModifiedDate.Time,
		data.ExpiredDate.Time,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "userClassHistoryRepository: DeleteUserClassHistory: error delete")
	}

	tx.Commit()
	return err == nil, nil
}
