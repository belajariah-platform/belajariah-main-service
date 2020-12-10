package handler

import "belajariah-main-service/usecase"

type notificationHandler struct {
	notificationUsecase usecase.NotificationUsecase
}

type NotificationHandler interface {
}

func InitNotificationHandler(notificationUsecase usecase.NotificationUsecase) NotificationHandler {
	return &notificationHandler{
		notificationUsecase,
	}
}
