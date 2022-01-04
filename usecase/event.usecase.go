package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type eventUsecase struct {
	eventRepository repository.EventRepository
}

type EventUsecase interface {
	GetAllEvent(r model.EventRequest) ([]model.Event, error)
	InsertFormClassIntens(ctx *gin.Context, r model.EventRequest) (bool, error)
}

func InitEventUsecase(eventRepository repository.EventRepository) EventUsecase {
	return &eventUsecase{
		eventRepository,
	}
}

func (u *eventUsecase) GetAllEvent(r model.EventRequest) ([]model.Event, error) {
	var event []model.Event
	var orderDefault = "ORDER BY country asc"
	var filterDefault = "is_deleted = false and is_active = true"
	filterFinal := utils.GetFilterOrderHandler(filterDefault, orderDefault, r.Query)

	eventEmpty := make([]model.Event, 0)

	result, err := u.eventRepository.GetAllEvent(filterFinal)
	if err != nil {
		return eventEmpty, utils.WrapError(err, "eventRepository.GetAllEvent")
	}

	if len(*result) > 0 {
		for _, ev := range *result {
			filterFinal = fmt.Sprintf(`WHERE event_code='%s'`, ev.Code)
			resultEv, err := u.eventRepository.GetAllEventMappingForm(filterFinal)
			if err != nil {
				return nil, utils.WrapError(err, "eventRepository.GetAllEvent")
			}

			if len(*resultEv) > 0 {
				ev.EventFormDetail = *resultEv
			} else {
				ev.EventFormDetail = make([]model.EventMappingForm, 0)
			}

			event = append(event, ev)
		}
	} else {
		return eventEmpty, nil
	}

	return event, nil
}

func (u *eventUsecase) InsertFormClassIntens(ctx *gin.Context, r model.EventRequest) (bool, error) {
	var err error
	email := ctx.Request.Header.Get("email")

	for _, ev := range r.Data.EventFormDetail {
		ev.Modified_By.Valid = true
		ev.Modified_Date.Valid = true

		ev.Created_By = email
		ev.Modified_By.String = email

		ev.Created_Date = r.Data.Modified_Date.Time
		ev.Modified_Date.Time = r.Data.Modified_Date.Time
		ev.User_Code = r.Data.User_Code

		_, err = u.eventRepository.InsertFormClassIntens(ev)
		if err != nil {
			return false, utils.WrapError(err, "eventUsecase.InsertFormClassIntens")
		}
	}

	return err == nil, nil
}
