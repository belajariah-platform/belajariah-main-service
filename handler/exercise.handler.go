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

type exerciseHandler struct {
	exerciseUsecase usecase.ExerciseUsecase
}

type ExerciseHandler interface {
	GetAllExercise(ctx *gin.Context)
}

func InitExerciseHandler(exerciseUsecase usecase.ExerciseUsecase) ExerciseHandler {
	return &exerciseHandler{
		exerciseUsecase,
	}
}

func (exerciseHandler *exerciseHandler) GetAllExercise(ctx *gin.Context) {
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

		var exerciseResult []shape.Exercise
		exerciseResult, count, err = exerciseHandler.exerciseUsecase.GetAllExercise(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":      exerciseResult,
				"dataCount": count,
				"error":     "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":      exerciseResult,
				"dataCount": count,
				"error":     err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}
