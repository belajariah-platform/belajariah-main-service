package usecase

import "belajariah-main-service/repository"

type notificationUsecase struct {
	notificationRepository repository.NotificationRepository
}

type NotificationUsecase interface{}

func InitNotificationUsecase(notificationRepository repository.NotificationRepository) NotificationUsecase {
	return &notificationUsecase{
		notificationRepository,
	}
}
