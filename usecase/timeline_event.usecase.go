package usecase

import "belajariah-main-service/repository"

type eventUsecase struct {
	eventRepository repository.EventRepository
}

type EventUsecase interface{}

func InitEventUsecase(eventRepository repository.EventRepository) EventUsecase {
	return &eventUsecase{
		eventRepository,
	}
}
