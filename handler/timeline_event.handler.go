package handler

import "belajariah-main-service/usecase"

type eventHandler struct {
	eventUsecase usecase.EventUsecase
}

type EventHandler interface {
}

func InitEventHandler(eventUsecase usecase.EventUsecase) EventHandler {
	return &eventHandler{
		eventUsecase,
	}
}
