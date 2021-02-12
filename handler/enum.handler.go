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

type enumHandler struct {
	enumUsecase usecase.EnumUsecase
}

type EnumHandler interface {
	GetAllEnum(ctx *gin.Context)
}

func InitEnumHandler(enumUsecase usecase.EnumUsecase) EnumHandler {
	return &enumHandler{
		enumUsecase,
	}
}

func (enumHandler *enumHandler) GetAllEnum(ctx *gin.Context) {
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

		var enumResult []shape.Enum
		enumResult, count, err = enumHandler.enumUsecase.GetAllEnum(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":      enumResult,
				"dataCount": count,
				"error":     "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":      enumResult,
				"dataCount": count,
				"error":     err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}
