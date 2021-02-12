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

type paymentMethodHandler struct {
	paymentMethodUsecase usecase.PaymentMethodUsecase
}

type PaymentMethodHandler interface {
	GetAllPaymentMethod(ctx *gin.Context)
}

func InitPaymentMethodHandler(paymentMethodUsecase usecase.PaymentMethodUsecase) PaymentMethodHandler {
	return &paymentMethodHandler{
		paymentMethodUsecase,
	}
}

func (paymentMethodHandler *paymentMethodHandler) GetAllPaymentMethod(ctx *gin.Context) {
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

		var paymentMethodResult []shape.PaymentMethod
		paymentMethodResult, count, err = paymentMethodHandler.paymentMethodUsecase.GetAllPaymentMethod(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":      paymentMethodResult,
				"dataCount": count,
				"error":     "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":      paymentMethodResult,
				"dataCount": count,
				"error":     err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}
