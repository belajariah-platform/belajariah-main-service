package handler

import (
	"belajariah-main-service/model"
	"belajariah-main-service/usecase"
	"belajariah-main-service/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type coachingProgramHandler struct {
	coachingProgramUsecase usecase.CoachingProgramUsecase
}

type CoachingProgramHandler interface {
	CoachingProgram(ctx *gin.Context)
}

func InitCoachingProgramHandler(u usecase.CoachingProgramUsecase) CoachingProgramHandler {
	return &coachingProgramHandler{
		u,
	}
}

func (h *coachingProgramHandler) CoachingProgram(ctx *gin.Context) {
	var request model.CoachingProgramRequest
	if err := ctx.ShouldBindJSON(&request); err == nil {
		switch strings.ToUpper(request.Action) {
		case utils.GET_ALL_MASTER_COACHING_PROGRAM:
			h.getAllMasterCoachingProgram(ctx, request)
		case utils.GET_ALL_COACHING_PROGRAM:
			h.getAllCoachingProgram(ctx, request)
		case utils.INSERT_COACHING_PROGRAM:
			h.insertCoachingProgram(ctx, request)
		case utils.CONFIRM_COACHING_PROGRAM:
			h.confirmCoachingProgram(ctx, request)
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

func (h *coachingProgramHandler) getAllMasterCoachingProgram(ctx *gin.Context, r model.CoachingProgramRequest) {
	result, err := h.coachingProgramUsecase.GetAllMasterCoachingProgram(r)
	if err != nil {
		utils.PushLogStackTrace("", utils.UnwrapError(err))
		utils.Response(ctx, struct{}{}, 0, "", err)
		return
	}

	utils.Response(ctx, result, len(result), "", nil)
}

func (h *coachingProgramHandler) getAllCoachingProgram(ctx *gin.Context, r model.CoachingProgramRequest) {
	result, err := h.coachingProgramUsecase.GetAllCoachingProgram(r)
	if err != nil {
		utils.PushLogStackTrace("", utils.UnwrapError(err))
		utils.Response(ctx, struct{}{}, 0, "", err)
		return
	}

	utils.Response(ctx, result, len(result), "", nil)
}

func (h *coachingProgramHandler) confirmCoachingProgram(ctx *gin.Context, r model.CoachingProgramRequest) {
	result, err := h.coachingProgramUsecase.ConfirmCoachingProgram(ctx, r)
	utils.Response(ctx, result, 1, "", err)
}

func (h *coachingProgramHandler) insertCoachingProgram(ctx *gin.Context, r model.CoachingProgramRequest) {
	result, err := h.coachingProgramUsecase.InsertCoachingProgram(ctx, r)
	utils.Response(ctx, result, 1, "", err)
}
