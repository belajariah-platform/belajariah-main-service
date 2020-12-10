package repository

import "github.com/jmoiron/sqlx"

type notificationRepository struct {
	db *sqlx.DB
}

type NotificationRepository interface{}

func InitNotificationRepository(db *sqlx.DB) NotificationRepository {
	return &notificationRepository{
		db,
	}
}
