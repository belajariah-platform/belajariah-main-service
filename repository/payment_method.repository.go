package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type paymentMethodRepository struct {
	db *sqlx.DB
}

type PaymentMethodRepository interface {
	GetAllPaymentMethod(skip, take int, filter string) ([]model.PaymentMethod, error)
	GetAllPaymentMethodCount(filter string) (int, error)
}

func InitPaymentMethodRepository(db *sqlx.DB) PaymentMethodRepository {
	return &paymentMethodRepository{
		db,
	}
}

func (paymentMethodRepository *paymentMethodRepository) GetAllPaymentMethod(skip, take int, filter string) ([]model.PaymentMethod, error) {
	var paymentMethodList []model.PaymentMethod
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		type,
		value,
		account_name,
		account_number,
		icon_account,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM 
		master.master_payment_method 
	WHERE 
		is_deleted=false AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := paymentMethodRepository.db.Query(query)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPaymentMethod => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var createdDate time.Time
			var isActive, isDeleted bool
			var modifiedDate sql.NullTime
			var code, types, values, createdBy string
			var accountName, accountNumber, iconAccount, modifiedBy sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&types,
				&values,
				&accountName,
				&accountNumber,
				&iconAccount,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&isDeleted,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllPaymentMethod => ", sqlError)
			} else {
				paymentMethodList = append(
					paymentMethodList,
					model.PaymentMethod{
						ID:            id,
						Code:          code,
						Type:          types,
						Value:         values,
						AccountName:   accountName,
						AccountNumber: accountNumber,
						IconAccount:   iconAccount,
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
	return paymentMethodList, sqlError
}

func (paymentMethodRepository *paymentMethodRepository) GetAllPaymentMethodCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		master.master_payment_method  
	WHERE 
		is_deleted=false AND
		is_active=true
	%s
	`, filter)

	row := paymentMethodRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPaymentMethodCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}
