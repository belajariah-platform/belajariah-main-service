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

type classHandler struct {
	classUsecase usecase.ClassUsecase
}

type ClassHandler interface {
	GetAllClass(ctx *gin.Context)
}

func InitClassHandler(classUsecase usecase.ClassUsecase) ClassHandler {
	return &classHandler{
		classUsecase,
	}
}

func (classHandler *classHandler) GetAllClass(ctx *gin.Context) {
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
		classResult, count, err = classHandler.classUsecase.GetAllClass(query)
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
