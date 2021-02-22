package model

import (
	"database/sql"
	"time"
)

type Notification struct {
	ID                   int
	Code                 string
	UserClassCode        sql.NullInt64
	PaymentCode          sql.NullInt64
	TableRef             string
	NotificationType     string
	NotificationTypeText string
	UserCode             int
	UserName             string
	Sequence             int
	IsRead               bool
	IsActionTaken        bool
	ExpiredDate          sql.NullTime
	IsActive             bool
	CreatedBy            string
	CreatedDate          time.Time
	ModifiedBy           sql.NullString
	ModifiedDate         sql.NullTime
	DeletedBy            sql.NullString
	DeletedDate          sql.NullTime
}
