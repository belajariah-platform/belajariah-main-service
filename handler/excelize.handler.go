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

type excelizeHandler struct {
	excelizeUsecase usecase.ExcelizeUsecase
}

type ExcelizeHandler interface {
	GetAllExcelize(ctx *gin.Context)
}

func InitExcelizeHandler(excelizeUsecase usecase.ExcelizeUsecase) ExcelizeHandler {
	return &excelizeHandler{
		excelizeUsecase,
	}
}

func (excelizeHandler *excelizeHandler) GetAllExcelize(ctx *gin.Context) {
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

		var excelizeResult []shape.UserInfo
		excelizeResult, err = excelizeHandler.excelizeUsecase.GetAllExcelize(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":  excelizeResult,
				"error": "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":  excelizeResult,
				"error": err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}
