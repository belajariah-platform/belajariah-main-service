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
	GetPayment(filter string) (model.Payment, error)
	GetAllPaymentCount(filter, filterUser string) (int, error)
	GetAllPayment(skip, take int, sort, search, filter, filterUser string) ([]model.Payment, error)

	UploadPayment(payment model.Payment) (bool, error)
	ConfirmPayment(payment model.Payment) (bool, error)
	InsertPayment(payment model.Payment) (model.Payment, bool, error)

	CheckAllPaymentExpired() ([]model.Payment, error)
	CheckAllPaymentBeforeExpired() ([]model.Payment, error)
}

func InitPaymentsRepository(db *sqlx.DB) PaymentsRepository {
	return &paymentsRepository{
		db,
	}
}

func (paymentsRepository *paymentsRepository) GetPayment(filter string) (model.Payment, error) {
	var paymentRow model.Payment
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		user_code,
		user_name,
		class_code,
		class_name,
		class_image,
		class_initial,
		package_code,
		package_type,
		package_discount,
		total_consultation,
		total_webinar,
		promo_code,
		promo_title,
		promo_discount,
		payment_method_code,
		payment_method,
		payment_method_type,
		payment_method_image,
		account_name,
		account_number,
		invoice_number,
		status_payment_code,
		status_payment,
		total_transfer,
		sender_bank,
		sender_name,
		image_proof,
		payment_type_code,
		payment_type,
		remarks,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date
	FROM 
		transaction.v_t_payment 
		%s
	`, filter)
	row := paymentsRepository.db.QueryRow(query)

	var isActive bool
	var id, totalTransfer int
	var createdDate time.Time
	var modifiedDate sql.NullTime
	var promoDiscount sql.NullFloat64
	var totalConsultation, totalWebinar, packageDiscount sql.NullInt64
	var promoCode, promoTitle, accountName, accountNumber, remarks, senderBank, senderName, imageProof, modifiedBy, classImage, paymentMethodImage sql.NullString
	var userName, classCode, className, classInitial, paymentMethodCode, paymentMethod, invoiceNumber, paymentTypeCode, paymentType, statusPaymentCode, statusPayment, packageCode, packageType, createdBy, paymentMethodType, code, userCode string

	sqlError := row.Scan(
		&id,
		&code,
		&userCode,
		&userName,
		&classCode,
		&className,
		&classImage,
		&classInitial,
		&packageCode,
		&packageType,
		&packageDiscount,
		&totalConsultation,
		&totalWebinar,
		&promoCode,
		&promoTitle,
		&promoDiscount,
		&paymentMethodCode,
		&paymentMethod,
		&paymentMethodType,
		&paymentMethodImage,
		&accountName,
		&accountNumber,
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
		utils.PushLogf("SQL error on GetPayment => ", sqlError.Error())
		return model.Payment{}, nil
	} else {
		paymentRow = model.Payment{
			ID:                 id,
			UserCode:           userCode,
			UserName:           userName,
			ClassCode:          classCode,
			ClassName:          className,
			ClassImage:         classImage,
			ClassInitial:       classInitial,
			PackageCode:        packageCode,
			PackageType:        packageType,
			PackageDiscount:    packageDiscount,
			TotalConsultation:  totalConsultation,
			TotalWebinar:       totalWebinar,
			PromoCode:          promoCode,
			PromoTitle:         promoTitle,
			PromoDiscount:      promoDiscount,
			PaymentMethodCode:  paymentMethodCode,
			PaymentMethod:      paymentMethod,
			PaymentMethodType:  paymentMethodType,
			PaymentMethodImage: paymentMethodImage,
			AccountName:        accountName,
			AccountNumber:      accountNumber,
			InvoiceNumber:      invoiceNumber,
			StatusPaymentCode:  statusPaymentCode,
			StatusPayment:      statusPayment,
			TotalTransfer:      totalTransfer,
			SenderBank:         senderBank,
			SenderName:         senderName,
			ImageProof:         imageProof,
			PaymentTypeCode:    paymentTypeCode,
			PaymentType:        paymentType,
			Remarks:            remarks,
			IsActive:           isActive,
			CreatedBy:          createdBy,
			CreatedDate:        createdDate,
			ModifiedBy:         modifiedBy,
			ModifiedDate:       modifiedDate,
		}
		return paymentRow, sqlError
	}
}

func (paymentsRepository *paymentsRepository) GetAllPayment(skip, take int, sort, search, filter, filterUser string) ([]model.Payment, error) {
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
		payment_method_type,
		payment_method_image,
		account_name,
		account_number,
		invoice_number,
		status_payment_code,
		status_payment,
		total_transfer,
		sender_bank,
		sender_name,
		image_proof,
		image_filename,
		payment_type_code,
		payment_type,
		remarks,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date
	FROM transaction.v_t_payment
		WHERE is_active = true
		%s %s %s %s
	OFFSET %d
	LIMIT %d
	`, filterUser, filter, search, sort, skip, take)

	rows, sqlError := paymentsRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPayment => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var isActive bool
			var id, totalTransfer int
			var createdDate time.Time
			var modifiedDate sql.NullTime
			var userName, classCode, className, classInitial, paymentMethodCode, paymentMethod, invoiceNumber, paymentTypeCode, paymentType, statusPaymentCode, statusPayment, packageCode, packageType, createdBy, paymentMethodType, code, userCode string
			var accounName, accountNumber, remarks, senderBank, senderName, imageProof, imageFilename, modifiedBy, paymentMethodImage sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&userCode,
				&userName,
				&classCode,
				&className,
				&classInitial,
				&packageCode,
				&packageType,
				&paymentMethodCode,
				&paymentMethod,
				&paymentMethodType,
				&paymentMethodImage,
				&accounName,
				&accountNumber,
				&invoiceNumber,
				&statusPaymentCode,
				&statusPayment,
				&totalTransfer,
				&senderBank,
				&senderName,
				&imageProof,
				&imageFilename,
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
				utils.PushLogf("SQL error on GetAllPayment => ", sqlError.Error())
			} else {
				paymentList = append(
					paymentList,
					model.Payment{
						ID:                 id,
						UserCode:           userCode,
						UserName:           userName,
						ClassCode:          classCode,
						ClassName:          className,
						ClassInitial:       classInitial,
						PackageCode:        packageCode,
						PackageType:        packageType,
						PaymentMethodCode:  paymentMethodCode,
						PaymentMethod:      paymentMethod,
						PaymentMethodType:  paymentMethodType,
						PaymentMethodImage: paymentMethodImage,
						AccountName:        accounName,
						AccountNumber:      accountNumber,
						InvoiceNumber:      invoiceNumber,
						StatusPaymentCode:  statusPaymentCode,
						StatusPayment:      statusPayment,
						TotalTransfer:      totalTransfer,
						SenderBank:         senderBank,
						SenderName:         senderName,
						ImageProof:         imageProof,
						ImageFilename:      imageFilename,
						PaymentTypeCode:    paymentTypeCode,
						PaymentType:        paymentType,
						Remarks:            remarks,
						IsActive:           isActive,
						CreatedBy:          createdBy,
						CreatedDate:        createdDate,
						ModifiedBy:         modifiedBy,
						ModifiedDate:       modifiedDate,
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
		transaction.v_t_payment 
	WHERE
		is_active = true 
		%s
	%s
	`, filterUser, filter)

	row := paymentsRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPaymentCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (paymentsRepository *paymentsRepository) InsertPayment(payment model.Payment) (model.Payment, bool, error) {
	var err error
	var result bool
	var payments model.Payment

	tx, errTx := paymentsRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in InsertPayment", errTx)
	} else {
		payments, err = insertPayment(tx, payment)
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

	return payments, result, err
}

func insertPayment(tx *sql.Tx, payment model.Payment) (model.Payment, error) {
	var id int
	var paymentRow model.Payment
	err := tx.QueryRow(`
	INSERT INTO transaction.transact_payment
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
		modified_date,
		promo_code
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
	) returning id
	`,
		payment.UserCode,
		payment.ClassCode,
		payment.PackageCode,
		payment.PaymentMethodCode,
		payment.InvoiceNumber,
		payment.StatusPaymentCode,
		payment.TotalTransfer,
		payment.PaymentTypeCode,
		payment.CreatedBy,
		payment.CreatedDate,
		payment.ModifiedBy.String,
		payment.ModifiedDate.Time,
		payment.PromoCode,
	).Scan(&id)

	if err != nil {
		utils.PushLogf("SQL error on Return insert payment => ", err)
		return model.Payment{}, nil
	} else {
		paymentRow = model.Payment{
			ID: id,
		}
		return paymentRow, err
	}
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
		transaction.transact_payment
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
		payment.PaymentMethodCode,
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
		transaction.transact_payment
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
		user_code,
		status_payment_code
	FROM transaction.v_t_payment
	WHERE  
		modified_date + interval '1440 minutes' <= now() AND 
		status_payment IN ('Waiting for Payment');
	`)

	if sqlError != nil {
		utils.PushLogf("SQL error on CheckAllPaymentExpired => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var statusPaymentCode, code, userCode string

			sqlError := rows.Scan(
				&id,
				&userCode,
				&statusPaymentCode,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on CheckAllPaymentExpired => ", sqlError.Error())
			} else {
				paymentList = append(paymentList, model.Payment{
					ID:                id,
					Code:              code,
					UserCode:          userCode,
					StatusPaymentCode: statusPaymentCode,
				})
			}
		}
	}
	return paymentList, sqlError
}

func (paymentsRepository *paymentsRepository) CheckAllPaymentBeforeExpired() ([]model.Payment, error) {
	var paymentList []model.Payment

	rows, sqlError := paymentsRepository.db.Query(`
	SELECT
		id,
		user_code,
		status_payment_code,
		modified_date
	FROM transaction.v_t_payment
	WHERE  
		DATE_PART('day', modified_date::timestamp - now()::timestamp) * 24 * 60 * 60 + 
		DATE_PART('hour', modified_date::timestamp - now()::timestamp) * 60 * 60 +
		DATE_PART('minute', modified_date::timestamp - now()::timestamp) * 60 +
		DATE_PART('second', modified_date::timestamp - now()::timestamp) + 86400 <= 7200 AND 
		status_payment IN ('Waiting for Payment');
	`)

	if sqlError != nil {
		utils.PushLogf("SQL error on CheckAllPaymentExpired => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var modifiedDate sql.NullTime
			var statusPaymentCode, code, userCode string

			sqlError := rows.Scan(
				&id,
				&userCode,
				&statusPaymentCode,
				&modifiedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on CheckAllPaymentExpired => ", sqlError.Error())
			} else {
				paymentList = append(paymentList, model.Payment{
					ID:                id,
					Code:              code,
					UserCode:          userCode,
					StatusPaymentCode: statusPaymentCode,
					ModifiedDate:      modifiedDate,
				})
			}
		}
	}
	return paymentList, sqlError
}
