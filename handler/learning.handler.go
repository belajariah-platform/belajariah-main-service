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

type learningHandler struct {
	learningUsecase usecase.LearningUsecase
}

type LearningHandler interface {
	Learning(ctx *gin.Context)
	GetAllLearning(ctx *gin.Context)
}

func InitLearningHandler(learningUsecase usecase.LearningUsecase) LearningHandler {
	return &learningHandler{
		learningUsecase,
	}
}

func (h *learningHandler) GetAllLearning(ctx *gin.Context) {
	var count int
	var query model.Query
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

		var learningResult []shape.Learning
		learningResult, count, err = h.learningUsecase.GetAllLearning(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":  learningResult,
				"count": count,
				"error": "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":  learningResult,
				"count": count,
				"error": err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}

func (h *learningHandler) Learning(ctx *gin.Context) {
	var request model.LearningQuranRequest
	if err := ctx.ShouldBindJSON(&request); err == nil {
		switch strings.ToUpper(request.Action) {
		case utils.GET_ALL_LEARNING_QURAN:
			h.getAllLearningQuran(ctx, request)
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

func (h *learningHandler) getAllLearningQuran(ctx *gin.Context, r model.LearningQuranRequest) {
	result, err := h.learningUsecase.GetAllLearningQuran(r)
	if err != nil {
		utils.PushLogStackTrace("", utils.UnwrapError(err))
		utils.Response(ctx, struct{}{}, 0, err)
		return
	}

	utils.Response(ctx, result, len(result), nil)
}
