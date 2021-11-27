package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	_getPayment = `
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
			schedule_code_1,
			schedule_code_2,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date
		FROM 
			transaction.v_t_payment 
		%s
	`
	_getAllPayment = `
		SELECT
			id,
			code,
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
			payment_type_code,
			payment_type,
			remarks,
			schedule_code_1,
			schedule_code_2,
			payment_reference,
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
	`
	_getAllPaymentCount = `
		SELECT COUNT(*) FROM
			transaction.v_t_payment 
		WHERE
			is_active = true 
			%s
		%s
	`
	_confirmPayment = `
		UPDATE
			transaction.transact_payment
		SET
			status_payment_code=$1,
			payment_type=$2,
			remarks=$3,
			modified_by=$4,
			modified_date=$5
		WHERE
			id=$6
	`
	_uploadPayment = `
		UPDATE
			transaction.transact_payment
		SET
			payment_method_code=$1,
			status_payment_code=$2,
			sender_bank=$3,
			sender_name=$4,
			image_proof=$5,
			modified_by=$6,
			modified_date=$7
		WHERE
			id=$8
	`
	_insertPayment = `
		INSERT INTO transaction.transact_payment
		(
			user_code,
			class_code,
			package_code,
			payment_method_code,
			invoice_number,
			status_payment_code,
			total_transfer,
			payment_type,
			schedule_code_1,
			schedule_code_2,
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
			(SELECT code FROM master.master_status WHERE value='Waiting for Payment' AND type='payment' LIMIT 1),
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13,
			$14
		) returning code
	`
	_checkAllPaymentExpired = `
		SELECT
			id,
			code,
			user_code,
			status_payment_code
		FROM transaction.v_t_payment
		WHERE  
			modified_date + interval '1440 minutes' <= now() AND 
			status_payment IN ('Waiting for Payment');
	`
	_checkAllPaymentBeforeExpired = `
		SELECT
			id,
			code,
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
	`
	_confirmPaymentQuran = `
		UPDATE
			transaction.transact_payment_quran
		SET
			status_payment_code=$1,
			payment_type=$2,
			remarks=$3,
			modified_by=$4,
			modified_date=$5
		WHERE
			code=$6
		`
	_uploadPaymentQuran = `
		UPDATE
			transaction.transact_payment_quran
		SET
			payment_method_code=$1,
			status_payment_code=$2,
			sender_bank=$3,
			sender_name=$4,
			image_proof=$5,
			modified_by=$6,
			modified_date=$7
		WHERE
			id=$8
	`
	_insertPaymentQuran = `
		INSERT INTO transaction.transact_payment_quran
		(
			user_code,
			class_code,
			package_code,
			payment_method_code,
			invoice_number,
			status_payment_code,
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
			(SELECT code FROM master.master_status WHERE value='Waiting for Payment' AND type='payment' LIMIT 1),
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12
		) returning code
	`
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

	UploadPaymentQuran(payment model.Payment) (bool, error)
	ConfirmPaymentQuran(payment model.Payment) (bool, error)
	InsertPaymentQuran(payment model.Payment) (model.Payment, bool, error)

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
	query := fmt.Sprintf(_getPayment, filter)

	row := paymentsRepository.db.QueryRow(query)

	var isActive bool
	var id, totalTransfer int
	var createdDate time.Time
	var modifiedDate sql.NullTime
	var promoDiscount sql.NullFloat64
	var scheduleCode1, scheduleCode2 sql.NullString
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
		&scheduleCode1,
		&scheduleCode2,
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
			ScheduleCode1:      scheduleCode1,
			ScheduleCode2:      scheduleCode2,
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
	query := fmt.Sprintf(_getAllPayment, filterUser, filter, search, sort, skip, take)

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
			var scheduleCode1, scheduleCode2 sql.NullString
			var userName, classCode, className, classInitial, paymentMethodCode, paymentMethod, invoiceNumber, paymentTypeCode, paymentType, statusPaymentCode, statusPayment, packageCode, packageType, createdBy, paymentMethodType, code, userCode string
			var accounName, accountNumber, remarks, senderBank, senderName, imageProof, modifiedBy, paymentMethodImage, paymentReference sql.NullString

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
				&paymentTypeCode,
				&paymentType,
				&remarks,
				&scheduleCode1,
				&scheduleCode2,
				&paymentReference,
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
						Code:               code,
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
						PaymentTypeCode:    paymentTypeCode,
						PaymentType:        paymentType,
						Remarks:            remarks,
						ScheduleCode1:      scheduleCode1,
						ScheduleCode2:      scheduleCode2,
						PaymentReference:   paymentReference,
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
	query := fmt.Sprintf(_getAllPaymentCount, filterUser, filter)

	row := paymentsRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPaymentCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (r *paymentsRepository) InsertPayment(payment model.Payment) (model.Payment, bool, error) {
	var code string
	var payments model.Payment

	tx, err := r.db.Beginx()
	if err != nil {
		return payments, false, errors.New("paymentsRepository: InsertPayment: error begin transaction")
	}

	err = tx.QueryRow(_insertPayment,
		payment.UserCode,
		payment.ClassCode,
		payment.PackageCode,
		payment.PaymentMethodCode,
		payment.InvoiceNumber,
		payment.TotalTransfer,
		payment.PaymentTypeCode,
		payment.ScheduleCode1.String,
		payment.ScheduleCode2.String,
		payment.CreatedBy,
		payment.CreatedDate,
		payment.ModifiedBy.String,
		payment.ModifiedDate.Time,
		payment.PromoCode.String,
	).Scan(&code)

	if err != nil {
		tx.Rollback()
		return payments, false, utils.WrapError(err, "paymentsRepository: InsertPayment: error insert")
	}

	payments = model.Payment{Code: code}

	tx.Commit()
	return payments, err == nil, nil
}

func (r *paymentsRepository) UploadPayment(payment model.Payment) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("paymentsRepository: UploadPayment: error begin transaction")
	}

	_, err = tx.Exec(_uploadPayment,
		payment.PaymentMethodCode,
		payment.StatusPaymentCode,
		payment.SenderBank.String,
		payment.SenderName.String,
		payment.ImageProof.String,
		payment.ModifiedBy.String,
		payment.ModifiedDate.Time,
		payment.ID,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "paymentsRepository: UploadPayment: error update")
	}

	tx.Commit()
	return err == nil, nil
}

func (r *paymentsRepository) ConfirmPayment(payment model.Payment) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("paymentsRepository: ConfirmPayment: error begin transaction")
	}

	_, err = tx.Exec(_confirmPayment,
		payment.StatusPaymentCode,
		payment.PaymentTypeCode,
		payment.Remarks.String,
		payment.ModifiedBy.String,
		payment.ModifiedDate.Time,
		payment.ID,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "paymentsRepository: ConfirmPayment: error update")
	}

	tx.Commit()
	return err == nil, nil
}

func (r *paymentsRepository) InsertPaymentQuran(payment model.Payment) (model.Payment, bool, error) {
	var code string
	var payments model.Payment

	tx, err := r.db.Beginx()
	if err != nil {
		return payments, false, errors.New("paymentsRepository: InsertPaymentQuran: error begin transaction")
	}

	err = tx.QueryRow(_insertPaymentQuran,
		payment.UserCode,
		payment.ClassCode,
		payment.PackageCode,
		payment.PaymentMethodCode,
		payment.InvoiceNumber,
		payment.TotalTransfer,
		payment.PaymentTypeCode,
		payment.CreatedBy,
		payment.CreatedDate,
		payment.ModifiedBy.String,
		payment.ModifiedDate.Time,
		payment.PromoCode.String,
	).Scan(&code)

	if err != nil {
		tx.Rollback()
		return payments, false, utils.WrapError(err, "paymentsRepository: InsertPaymentQuran: error insert")
	}

	payments = model.Payment{Code: code}

	tx.Commit()
	return payments, err == nil, nil
}

func (r *paymentsRepository) UploadPaymentQuran(payment model.Payment) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("paymentsRepository: UploadPaymentQuran: error begin transaction")
	}

	_, err = tx.Exec(_uploadPaymentQuran,
		payment.PaymentMethodCode,
		payment.StatusPaymentCode,
		payment.SenderBank.String,
		payment.SenderName.String,
		payment.ImageProof.String,
		payment.ModifiedBy.String,
		payment.ModifiedDate.Time,
		payment.ID,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "paymentsRepository: UploadPaymentQuran: error update")
	}

	tx.Commit()
	return err == nil, nil
}

func (r *paymentsRepository) ConfirmPaymentQuran(payment model.Payment) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("paymentsRepository: ConfirmPaymentQuran: error begin transaction")
	}

	_, err = tx.Exec(_confirmPaymentQuran,
		payment.StatusPaymentCode,
		payment.PaymentTypeCode,
		payment.Remarks.String,
		payment.ModifiedBy.String,
		payment.ModifiedDate.Time,
		payment.Code,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "paymentsRepository: ConfirmPaymentQuran: error update")
	}

	tx.Commit()
	return err == nil, nil
}

func (paymentsRepository *paymentsRepository) CheckAllPaymentExpired() ([]model.Payment, error) {
	var paymentList []model.Payment

	rows, sqlError := paymentsRepository.db.Query(_checkAllPaymentExpired)

	if sqlError != nil {
		utils.PushLogf("SQL error on CheckAllPaymentExpired => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var statusPaymentCode, code, userCode string

			sqlError := rows.Scan(
				&id,
				&code,
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

	rows, sqlError := paymentsRepository.db.Query(_checkAllPaymentBeforeExpired)

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
				&code,
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
