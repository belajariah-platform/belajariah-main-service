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
	_getNotification = `
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
			transaction.v_t_notifications
		WHERE 
			is_deleted=false AND
			split_part(notification_type_text, '|', 1)='%s' 
			%s
	`
	_insertNotification = `
		INSERT INTO transaction.transact_notification
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
			)
	`
)

type notificationRepository struct {
	db *sqlx.DB
}

type NotificationRepository interface {
	GetNotification(filter, types string) (model.Notification, error)
	InsertNotification(notification model.Notification) (bool, error)
}

func InitNotificationRepository(db *sqlx.DB) NotificationRepository {
	return &notificationRepository{
		db,
	}
}

func (notificationRepository *notificationRepository) GetNotification(filter, types string) (model.Notification, error) {
	var notificationRow model.Notification
	query := fmt.Sprintf(_getNotification, types, filter)

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

func (r *notificationRepository) InsertNotification(notification model.Notification) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("notificationRepository: InsertNotification: error begin transaction")
	}

	_, err = tx.Exec(_insertNotification,
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

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "notificationRepository: InsertNotification: error insert")
	}

	tx.Commit()
	return err == nil, nil
}
