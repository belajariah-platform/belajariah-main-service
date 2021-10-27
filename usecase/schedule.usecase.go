package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/utils"

	"github.com/gin-gonic/gin"
)

type scheduleUsecase struct {
	scheduleRepository repository.ScheduleRepository
}
type ScheduleUsecase interface {
	GetAllSchedule(r model.ScheduleRequest) ([]model.Schedule, error)
	InsertSchedule(ctx *gin.Context, r model.ScheduleRequest) (bool, error)
	UpdateScheduleUser(ctx *gin.Context, r model.ScheduleRequest) (bool, error)
	UpdateScheduleMentor(ctx *gin.Context, r model.ScheduleRequest) (bool, error)
}

func InitScheduleUsecase(scheduleRepository repository.ScheduleRepository) ScheduleUsecase {
	return &scheduleUsecase{
		scheduleRepository,
	}
}

func (u *scheduleUsecase) GetAllSchedule(r model.ScheduleRequest) ([]model.Schedule, error) {
	var orderDefault = "ORDER BY code asc"
	var filterDefault = "is_deleted = false and is_active = true"
	filterFinal := utils.GetFilterOrderHandler(filterDefault, orderDefault, r.Query)

	result, err := u.scheduleRepository.GetAllSchedule(filterFinal)
	if err != nil {
		return nil, utils.WrapError(err, "scheduleUsecase.GetAllSchedule")
	}

	return *result, nil
}

func (u *scheduleUsecase) InsertSchedule(ctx *gin.Context, r model.ScheduleRequest) (bool, error) {

	result, err := u.scheduleRepository.InsertSchedule(r.Data)
	if err != nil {
		return false, utils.WrapError(err, "scheduleUsecase.InsertSchedule")
	}

	return result, nil
}

func (u *scheduleUsecase) UpdateScheduleUser(ctx *gin.Context, r model.ScheduleRequest) (bool, error) {
	email := ctx.Request.Header.Get("email")

	r.Data.Modified_By = email

	result, err := u.scheduleRepository.UpdateScheduleUser(r.Data)
	if err != nil {
		return false, utils.WrapError(err, "scheduleUsecase.UpdateScheduleUser")
	}

	return result, nil
}

func (u *scheduleUsecase) UpdateScheduleMentor(ctx *gin.Context, r model.ScheduleRequest) (bool, error) {
	email := ctx.Request.Header.Get("email")

	r.Data.Modified_By = email

	result, err := u.scheduleRepository.UpdateScheduleMentor(r.Data)
	if err != nil {
		return false, utils.WrapError(err, "scheduleUsecase.UpdateScheduleMentor")
	}

	return result, nil
}
