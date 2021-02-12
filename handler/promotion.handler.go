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

type promotionHandler struct {
	promotionUsecase usecase.PromotionUsecase
}

type PromotionHandler interface {
	GetAllPromotion(ctx *gin.Context)
	GetPromotion(ctx *gin.Context)
}

func InitPromotionHandler(promotionUsecase usecase.PromotionUsecase) PromotionHandler {
	return &promotionHandler{
		promotionUsecase,
	}
}

func (promotionHandler *promotionHandler) GetAllPromotion(ctx *gin.Context) {
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

		var promotionResult []shape.Promotion
		promotionResult, count, err = promotionHandler.promotionUsecase.GetAllPromotion(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":      promotionResult,
				"dataCount": count,
				"error":     "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":      promotionResult,
				"dataCount": count,
				"error":     err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}

func (promotionHandler *promotionHandler) GetPromotion(ctx *gin.Context) {
	code := ctx.Param("code")

	result, err := promotionHandler.promotionUsecase.GetPromotion(code)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": result,
			"error":  "",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": result,
			"error":  err.Error(),
		})
	}
}
