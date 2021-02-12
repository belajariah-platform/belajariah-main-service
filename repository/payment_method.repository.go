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
		method_image,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM 
		v_m_payment_method 
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := paymentMethodRepository.db.Query(query)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPaymentMethod => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var isActive bool
			var createdDate time.Time
			var modifiedDate, deletedDate sql.NullTime
			var code, types, values, createdBy string
			var accountName, accountNumber, methodImage, modifiedBy, deletedBy sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&types,
				&values,
				&accountName,
				&accountNumber,
				&methodImage,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
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
						MethodImage:   methodImage,
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
	return paymentMethodList, sqlError
}

func (paymentMethodRepository *paymentMethodRepository) GetAllPaymentMethodCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		v_m_payment_method  
	WHERE 
		deleted_by IS NULL AND
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
