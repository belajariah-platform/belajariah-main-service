package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type paymentsRepository struct {
	db *sqlx.DB
}

type PaymentsRepository interface {
	GetAllPayment(skip, take int, filter, filterUser string) ([]model.Payment, error)
	GetAllPaymentCount(filter, filterUser string) (int, error)
	InsertPayment(payment model.Payment) (bool, error)
	UploadPayment(payment model.Payment) (bool, error)
	ConfirmPayment(payment model.Payment) (bool, error)
	CheckAllPaymentExpired() ([]model.Payment, error)
}

func InitPaymentsRepository(db *sqlx.DB) PaymentsRepository {
	return &paymentsRepository{
		db,
	}
}

func (paymentsRepository *paymentsRepository) GetAllPayment(skip, take int, filter, filterUser string) ([]model.Payment, error) {
	var paymentList []model.Payment
	query := fmt.Sprintf(`
	SELECT
		id,
		user_code,
		user_name,
		class_code,
		class_name,
		class_initial,
		package_code,
		package_type,
		payment_method_code,
		payment_method,
		invoice_number,
		status_payment_code,
		status_payment,
		total_transfer,
		sender_bank,
		destination_account,
		image_proof,
		payment_type_code,
		payment_type,
		remarks,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date
	FROM v_t_payment
		%s
	%s
	OFFSET %d
	LIMIT %d
	`, filterUser, filter, skip, take)

	rows, sqlError := paymentsRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPayment => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var isActive bool
			var createdDate time.Time
			var modifiedDate sql.NullTime
			var id, userCode, totalTransfer int
			var remarks, senderBank, senderName, imageProof, modifiedBy sql.NullString
			var userName, classCode, className, classInitial, paymentMethodCode, paymentMethod, invoiceNumber, paymentTypeCode, paymentType, statusPaymentCode, statusPayment, packageCode, packageType, createdBy string

			sqlError := rows.Scan(
				&id,
				&userCode,
				&userName,
				&classCode,
				&className,
				&classInitial,
				&packageCode,
				&packageType,
				&paymentMethodCode,
				&paymentMethod,
				&invoiceNumber,
				&statusPaymentCode,
				&statusPayment,
				&totalTransfer,
				&senderBank,
				&senderName,
				&imageProof,
				&paymentTypeCode,
				&paymentType,
				&remarks,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllPayment => ", sqlError)
			} else {
				paymentList = append(
					paymentList,
					model.Payment{
						ID:                id,
						UserCode:          userCode,
						UserName:          userName,
						ClassCode:         classCode,
						ClassName:         className,
						ClassInitial:      classInitial,
						PackageCode:       packageCode,
						PackageType:       packageType,
						PaymentMethodCode: paymentMethodCode,
						PaymentMethod:     paymentMethod,
						InvoiceNumber:     invoiceNumber,
						StatusPaymentCode: statusPaymentCode,
						StatusPayment:     statusPayment,
						TotalTransfer:     totalTransfer,
						SenderBank:        senderBank,
						SenderName:        senderName,
						ImageProof:        imageProof,
						PaymentTypeCode:   paymentTypeCode,
						PaymentType:       paymentType,
						Remarks:           remarks,
						IsActive:          isActive,
						CreatedBy:         createdBy,
						CreatedDate:       createdDate,
						ModifiedBy:        modifiedBy,
						ModifiedDate:      modifiedDate,
					},
				)
			}
		}
	}
	return paymentList, sqlError
}

func (paymentsRepository *paymentsRepository) GetAllPaymentCount(filter, filterUser string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM
		v_t_payment
		%s
	%s
	`, filterUser, filter)

	row := paymentsRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPaymentCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}

func (paymentsRepository *paymentsRepository) InsertPayment(payment model.Payment) (bool, error) {
	var err error
	var result bool

	tx, errTx := paymentsRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in InsertPayment", errTx)
	} else {
		err = insertPayment(tx, payment)
		if err != nil {
			utils.PushLogf("err in payment---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to InsertPayment", err)
	}

	return result, err
}

func insertPayment(tx *sql.Tx, payment model.Payment) error {
	_, err := tx.Exec(`
	INSERT INTO transact_payment
	(
		user_code,
		class_code,
		package_code,
		payment_method_code,
		invoice_number,
		status_payment,
		total_transfer,
		payment_type,
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
		$11,
		$12
	);
	`,
		payment.UserCode,
		payment.ClassCode,
		payment.PackageCode,
		payment.PaymentMethod,
		payment.InvoiceNumber,
		payment.StatusPaymentCode,
		payment.TotalTransfer,
		payment.PaymentTypeCode,
		payment.CreatedBy,
		payment.CreatedDate,
		payment.ModifiedBy.String,
		payment.ModifiedDate.Time,
	)
	return err
}

func (paymentsRepository *paymentsRepository) UploadPayment(payment model.Payment) (bool, error) {
	var err error
	var result bool

	tx, errTx := paymentsRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in UploadPayment", errTx)
	} else {
		err = uploadPayment(tx, payment)
		if err != nil {
			utils.PushLogf("err in payment---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to uploadPayment", err)
	}

	return result, err
}

func uploadPayment(tx *sql.Tx, payment model.Payment) error {
	_, err := tx.Exec(`
	UPDATE
		transact_payment
	 SET
		payment_method_code=$1,
		status_payment=$2,
		sender_bank=$3,
		sender_name=$4,
		image_code=$5,
		modified_by=$6,
		modified_date=$7
 	WHERE
 		id=$8
	`,
		payment.PaymentMethod,
		payment.StatusPaymentCode,
		payment.SenderBank.String,
		payment.SenderName.String,
		payment.ImageCode.Int64,
		payment.ModifiedBy.String,
		payment.ModifiedDate.Time,
		payment.ID,
	)
	return err
}

func (paymentsRepository *paymentsRepository) ConfirmPayment(payment model.Payment) (bool, error) {
	var err error
	var result bool

	tx, errTx := paymentsRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in ConfirmPayment", errTx)
	} else {
		err = confirmPayment(tx, payment)
		if err != nil {
			utils.PushLogf("err in payment---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to ConfirmPayment", err)
	}

	return result, err
}

func confirmPayment(tx *sql.Tx, payment model.Payment) error {
	_, err := tx.Exec(`
	UPDATE
		transact_payment
	 SET
		status_payment=$1,
		payment_type=$2,
		remarks=$3,
		modified_by=$4,
		modified_date=$5
 	WHERE
 		id=$6
	`,
		payment.StatusPaymentCode,
		payment.PaymentTypeCode,
		payment.Remarks.String,
		payment.ModifiedBy.String,
		payment.ModifiedDate.Time,
		payment.ID,
	)
	return err
}

func (paymentsRepository *paymentsRepository) CheckAllPaymentExpired() ([]model.Payment, error) {
	var paymentList []model.Payment

	rows, sqlError := paymentsRepository.db.Query(`
	SELECT
		id,
		status_payment_code
	FROM v_m_payments
	WHERE  
		created_date <= now() AND 
		status_payment IN ('Waiting for Payment');
	`)

	if sqlError != nil {
		utils.PushLogf("SQL error on CheckAllPaymentExpired => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var statusPaymentCode string
			var id int

			sqlError := rows.Scan(&id, &statusPaymentCode)

			if sqlError != nil {
				utils.PushLogf("SQL error on CheckAllPaymentExpired => ", sqlError)
			} else {
				paymentList = append(paymentList, model.Payment{
					ID:                id,
					StatusPaymentCode: statusPaymentCode,
				})
			}
		}
	}
	return paymentList, sqlError
}
