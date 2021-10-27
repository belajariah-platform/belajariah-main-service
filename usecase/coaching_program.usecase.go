package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/utils"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type coachingProgramUsecase struct {
	coachingProgramRepository repository.CoachingProgramRepository
}

type CoachingProgramUsecase interface {
	GetAllMasterCoachingProgram(r model.CoachingProgramRequest) ([]model.MasterCoachingProgram, error)
	GetAllCoachingProgram(r model.CoachingProgramRequest) ([]model.CoachingProgram, error)
	InsertCoachingProgram(ctx *gin.Context, r model.CoachingProgramRequest) (bool, error)
	ConfirmCoachingProgram(ctx *gin.Context, r model.CoachingProgramRequest) (bool, error)
}

func InitCoachingProgramUsecase(cp repository.CoachingProgramRepository) CoachingProgramUsecase {
	return &coachingProgramUsecase{
		cp,
	}
}

func (u *coachingProgramUsecase) GetAllMasterCoachingProgram(r model.CoachingProgramRequest) ([]model.MasterCoachingProgram, error) {
	var orderDefault = "ORDER BY code asc"
	var filterDefault = "is_deleted = false and is_active = true"
	filterFinal := utils.GetFilterOrderHandler(filterDefault, orderDefault, r.Query)

	result, err := u.coachingProgramRepository.GetAllMasterCoachingProgram(filterFinal)
	if err != nil {
		return nil, utils.WrapError(err, "coachingProgramUsecase.GetAllMasterCoachingProgram")
	}

	return *result, nil
}

func (u *coachingProgramUsecase) GetAllCoachingProgram(r model.CoachingProgramRequest) ([]model.CoachingProgram, error) {
	var orderDefault = "ORDER BY code asc"
	var filterDefault = "is_deleted = false and is_active = true"
	filterFinal := utils.GetFilterOrderHandler(filterDefault, orderDefault, r.Query)

	result, err := u.coachingProgramRepository.GetAllCoachingProgram(filterFinal)
	if err != nil {
		return nil, utils.WrapError(err, "coachingProgramUsecase.GetAllMasterCoachingProgram")
	}

	return *result, nil
}

func (u *coachingProgramUsecase) InsertCoachingProgram(ctx *gin.Context, r model.CoachingProgramRequest) (bool, error) {
	email := ctx.Request.Header.Get("email")

	r.Data_TCP.Created_By = email
	r.Data_TCP.Modified_By = email
	r.Data_TCP.Created_Date = r.Data_TCP.Modified_Date.Time

	filter := fmt.Sprintf(`WHERE user_code='%s'`, r.Data_TCP.User_Code)
	data, err := u.coachingProgramRepository.GetAllCoachingProgram(filter)
	if err != nil {
		return false, utils.WrapError(err, "coachingProgramUsecase.GetAllCoachingProgram")
	}

	if len(*data) != 0 {
		return false, errors.New("coachingProgramUsecase.GetAllCoachingProgram: error account already used")
	}

	result, err := u.coachingProgramRepository.InsertCoachingProgram(r.Data_TCP)
	if err != nil {
		return false, utils.WrapError(err, "coachingProgramUsecase.InsertCoachingProgram")
	}

	return result, nil
}

func (u *coachingProgramUsecase) ConfirmCoachingProgram(ctx *gin.Context, r model.CoachingProgramRequest) (bool, error) {
	email := ctx.Request.Header.Get("email")

	r.Data_TCP.Is_Confirmed = true
	r.Data_TCP.Modified_By = email

	result, err := u.coachingProgramRepository.ConfirmCoachingProgram(r.Data_TCP)
	if err != nil {
		return false, utils.WrapError(err, "coachingProgramUsecase.ConfirmCoachingProgram")
	}

	return result, nil
}
