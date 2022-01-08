package handler

import (
	"belajariah-main-service/model"
	"belajariah-main-service/usecase"
	"belajariah-main-service/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type eventHandler struct {
	eventUsecase usecase.EventUsecase
}

type EventHandler interface {
	Event(ctx *gin.Context)
}

func InitEventHandler(eventUsecase usecase.EventUsecase) EventHandler {
	return &eventHandler{
		eventUsecase,
	}
}

func (h *eventHandler) Event(ctx *gin.Context) {
	var request model.EventRequest
	if err := ctx.ShouldBindJSON(&request); err == nil {
		switch strings.ToUpper(request.Action) {
		case utils.GET_ALL_EVENT:
			h.getAllEvent(ctx, request)
		case utils.INSERT_FORM_CLASS_INTENS:
			h.insertFormClassIntens(ctx, request)
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

func (h *eventHandler) getAllEvent(ctx *gin.Context, r model.EventRequest) {
	result, err := h.eventUsecase.GetAllEvent(r)
	if err != nil {
		utils.PushLogStackTrace("", utils.UnwrapError(err))
		utils.Response(ctx, struct{}{}, 0, "", err)
		return
	}

	utils.Response(ctx, result, len(result), "", nil)
}

func (h *eventHandler) insertFormClassIntens(ctx *gin.Context, r model.EventRequest) {
	result, err := h.eventUsecase.InsertFormClassIntens(ctx, r)
	utils.Response(ctx, result, 1, "", err)
}
