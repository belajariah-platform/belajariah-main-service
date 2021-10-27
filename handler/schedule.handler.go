package handler

import (
	"belajariah-main-service/model"
	"belajariah-main-service/usecase"
	"belajariah-main-service/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type scheduleHandler struct {
	scheduleUsecase usecase.ScheduleUsecase
}

type ScheduleHandler interface {
	Schedule(ctx *gin.Context)
}

func InitScheduleHandler(u usecase.ScheduleUsecase) ScheduleHandler {
	return &scheduleHandler{
		u,
	}
}

func (h *scheduleHandler) Schedule(ctx *gin.Context) {
	var request model.ScheduleRequest
	if err := ctx.ShouldBindJSON(&request); err == nil {
		switch strings.ToUpper(request.Action) {
		case utils.GET_ALL_SCHEDULE:
			h.getAllSchedule(ctx, request)
		case utils.UPDATE_USER_SCHEDULE:
			h.updateScheuleUser(ctx, request)
		case utils.UPDATE_MENTOR_SCHEDULE:
			h.updateScheuleMentor(ctx, request)
		default:
			utils.NotFoundActionResponse(ctx, request.Action)
		}
	} else {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: model.RequestResponse{
				Count:  0,
				Data:   nil,
				Error:  err.Error(),
				Result: false,
			},
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
	}
}

func (h *scheduleHandler) getAllSchedule(ctx *gin.Context, r model.ScheduleRequest) {
	result, err := h.scheduleUsecase.GetAllSchedule(r)
	if err != nil {
		utils.PushLogStackTrace("", utils.UnwrapError(err))
		utils.Response(ctx, struct{}{}, 0, err)
		return
	}

	utils.Response(ctx, result, len(result), nil)
}

func (h *scheduleHandler) updateScheuleUser(ctx *gin.Context, r model.ScheduleRequest) {
	result, err := h.scheduleUsecase.UpdateScheduleUser(ctx, r)
	utils.Response(ctx, result, 1, err)
}

func (h *scheduleHandler) updateScheuleMentor(ctx *gin.Context, r model.ScheduleRequest) {
	result, err := h.scheduleUsecase.UpdateScheduleMentor(ctx, r)
	utils.Response(ctx, result, 1, err)
}
