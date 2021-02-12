package handler

import (
	"belajariah-main-service/model"
	"belajariah-main-service/shape"
	"belajariah-main-service/usecase"
	"belajariah-main-service/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type learningHandler struct {
	learningUsecase usecase.LearningUsecase
}

type LearningHandler interface {
	GetAllLearning(ctx *gin.Context)
}

func InitLearningHandler(learningUsecase usecase.LearningUsecase) LearningHandler {
	return &learningHandler{
		learningUsecase,
	}
}

func (learningHandler *learningHandler) GetAllLearning(ctx *gin.Context) {
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

		var learningResult []shape.Learning
		learningResult, count, err = learningHandler.learningUsecase.GetAllLearning(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":      learningResult,
				"dataCount": count,
				"error":     "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":      learningResult,
				"dataCount": count,
				"error":     err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}
