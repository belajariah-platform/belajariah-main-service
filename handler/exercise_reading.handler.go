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

type exerciseReadingHandler struct {
	exerciseReadingUsecase usecase.ExerciseReadingUsecase
}

type ExerciseReadingHandler interface {
	GetAllExerciseReading(ctx *gin.Context)
}

func InitExerciseReadingHandler(exerciseReadingUsecase usecase.ExerciseReadingUsecase) ExerciseReadingHandler {
	return &exerciseReadingHandler{
		exerciseReadingUsecase,
	}
}

func (exerciseReadingHandler *exerciseReadingHandler) GetAllExerciseReading(ctx *gin.Context) {
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

		var exerciseReadingResult []shape.ExerciseReading
		exerciseReadingResult, count, err = exerciseReadingHandler.exerciseReadingUsecase.GetAllExerciseReading(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":  exerciseReadingResult,
				"count": count,
				"error": "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":  exerciseReadingResult,
				"count": count,
				"error": err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}
