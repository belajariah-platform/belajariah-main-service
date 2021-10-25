package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type notificationRepository struct {
	db *sqlx.DB
}

type NotificationRepository interface {
	GetNotificationCount(code, types string) (int, error)
	GetNotification(filter, types string) (model.Notification, error)
	InsertNotification(notification model.Notification) (bool, error)
}

func InitNotificationRepository(db *sqlx.DB) NotificationRepository {
	return &notificationRepository{
		db,
	}
}

func (notificationRepository *notificationRepository) GetNotificationCount(code, types string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) 
	FROM 
		v_t_notifications 
	WHERE 
		deleted_by IS NULL AND
		code_ref = '%s' AND 
		notification_type_text = '%s'
	`, code, types)

	row := notificationRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetNotificationCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}

func (notificationRepository *notificationRepository) GetNotification(filter, types string) (model.Notification, error) {
	var notificationRow model.Notification
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		user_class_code,
		payment_code,
		table_ref,
		notification_type,
		notification_type_text,
		user_code,
		user_name,
		sequence,
		is_read,
		is_action_taken,
		expired_date
	FROM
		v_t_notifications
	WHERE 
		deleted_by IS NULL AND
		split_part(notification_type_text, '|', 1)='%s' 
		%s
	`, types, filter)

	row := notificationRepository.db.QueryRow(query)
	var expiredDate sql.NullTime
	var id, sequence int
	var isRead, isActionTaken bool
	var userClassCode, paymentCode sql.NullInt64
	var code, tableRef, notificationType, notificationTypeText, userName, userCode string

	sqlError := row.Scan(
		&id,
		&code,
		&userClassCode,
		&paymentCode,
		&tableRef,
		&notificationType,
		&notificationTypeText,
		&userCode,
		&userName,
		&sequence,
		&isRead,
		&isActionTaken,
		&expiredDate,
	)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetNotification => ", sqlError)
		return model.Notification{}, nil
	} else {
		notificationRow = model.Notification{
			ID:                   id,
			Code:                 code,
			UserClassCode:        userClassCode,
			PaymentCode:          paymentCode,
			TableRef:             tableRef,
			NotificationType:     notificationType,
			NotificationTypeText: notificationTypeText,
			UserCode:             userCode,
			UserName:             userName,
			Sequence:             sequence,
			IsRead:               isRead,
			IsActionTaken:        isActionTaken,
			ExpiredDate:          expiredDate,
		}
		return notificationRow, sqlError
	}
}

func (notificationRepository *notificationRepository) InsertNotification(notification model.Notification) (bool, error) {
	var err error
	var result bool

	tx, errTx := notificationRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in insertNotification", errTx)
	} else {
		err = insertNotification(tx, notification)
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to insertNotification")
	}

	return result, err
}

func insertNotification(tx *sql.Tx, notification model.Notification) error {
	sqlQuery := `
	INSERT INTO transact_notification
	(
		table_ref,
		user_class_code,
		payment_code,
		notification_type,
		user_code,
		sequence,
		created_by,
		created_date,
		modified_by,
		modified_date,
		expired_date
	)
	VALUES (
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
		);
`
	_, err := tx.Exec(sqlQuery,
		notification.TableRef,
		notification.UserClassCode.Int64,
		notification.PaymentCode,
		notification.NotificationType,
		notification.UserCode,
		notification.Sequence,
		notification.CreatedBy,
		notification.CreatedDate,
		notification.ModifiedBy.String,
		notification.ModifiedDate.Time,
		notification.ExpiredDate.Time,
	)
	return err
}
