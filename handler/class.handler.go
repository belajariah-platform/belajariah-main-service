package handler

import (
	"belajariah-main-service/model"
	"belajariah-main-service/shape"
	"belajariah-main-service/usecase"
	"belajariah-main-service/utils"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type classHandler struct {
	classUsecase usecase.ClassUsecase
}

type ClassHandler interface {
	Class(ctx *gin.Context)
	GetAllClass(ctx *gin.Context)
}

func InitClassHandler(classUsecase usecase.ClassUsecase) ClassHandler {
	return &classHandler{
		classUsecase,
	}
}

func (h *classHandler) GetAllClass(ctx *gin.Context) {
	var query model.Query
	var count int
	err := ctx.BindQuery(&query)

	if err == nil {
		var array []map[string]interface{}
		if err := json.Unmarshal([]byte(query.Filter), &array); err != nil {
			panic(err)
		}
		for _, arr := range array {
			query.Filters = append(query.Filters, model.Filter{
				Type:  arr["type"].(string),
				Field: arr["field"].(string),
				Value: arr["value"].(string),
			})
		}

		var classResult []shape.Class
		classResult, count, err = h.classUsecase.GetAllClass(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":  classResult,
				"count": count,
				"error": "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":  classResult,
				"count": count,
				"error": err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}

func (h *classHandler) Class(ctx *gin.Context) {
	var request model.ClassQuranRequest
	if err := ctx.ShouldBindJSON(&request); err == nil {
		switch strings.ToUpper(request.Action) {
		case utils.GET_ALL_CLASS_QURAN:
			h.getAllClassQuran(ctx, request)
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

func (h *classHandler) getAllClassQuran(ctx *gin.Context, r model.ClassQuranRequest) {
	result, count, err := h.classUsecase.GetAllClassQuran(r)
	if err != nil {
		utils.PushLogStackTrace("", utils.UnwrapError(err))
		utils.Response(ctx, struct{}{}, 0, err)
		return
	}

	utils.Response(ctx, result, count, nil)
}
