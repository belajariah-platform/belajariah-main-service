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

type paymentHandler struct {
	paymentUsecase usecase.PaymentUsecase
}

type PaymentHandler interface {
	GetAllPayment(ctx *gin.Context)
	GetAllPaymentByUserID(ctx *gin.Context)
	InsertPayment(ctx *gin.Context)
	UploadPayment(ctx *gin.Context)
	ConfirmPayment(ctx *gin.Context)
}

func InitPaymentHandler(paymentUsecase usecase.PaymentUsecase) PaymentHandler {
	return &paymentHandler{
		paymentUsecase,
	}
}

func (paymentHandler *paymentHandler) GetAllPayment(ctx *gin.Context) {
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

		var paymentResult []shape.Payment
		paymentResult, count, err = paymentHandler.paymentUsecase.GetAllPayment(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":      paymentResult,
				"dataCount": count,
				"error":     "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":      paymentResult,
				"dataCount": count,
				"error":     err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}

func (paymentHandler *paymentHandler) GetAllPaymentByUserID(ctx *gin.Context) {
	var query model.Query
	var count int
	err := ctx.BindQuery(&query)

	var userObj model.UserInfo
	for _, valueUser := range ctx.Request.Header["User"] {
		itemInfoBytes := []byte(valueUser)

		er := json.Unmarshal(itemInfoBytes, &userObj)
		if er != nil {
			utils.PushLogf("[Error Unmarshal] :", er)
		}
	}

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

		var paymentResult []shape.Payment
		paymentResult, count, err = paymentHandler.paymentUsecase.GetAllPaymentByUserID(query, userObj)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":      paymentResult,
				"dataCount": count,
				"error":     "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":      paymentResult,
				"dataCount": count,
				"error":     err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}

func (paymentHandler *paymentHandler) InsertPayment(ctx *gin.Context) {
	var payment shape.PaymentPost
	var email string
	for _, value := range ctx.Request.Header["Email"] {
		email = value
		break
	}
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := paymentHandler.paymentUsecase.InsertPayment(payment, email)
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

func (paymentHandler *paymentHandler) UploadPayment(ctx *gin.Context) {
	var payment shape.PaymentPost
	var email string
	for _, value := range ctx.Request.Header["Email"] {
		email = value
		break
	}
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := paymentHandler.paymentUsecase.UploadPayment(payment, email)
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

func (paymentHandler *paymentHandler) ConfirmPayment(ctx *gin.Context) {
	var payment shape.PaymentPost
	var email string
	for _, value := range ctx.Request.Header["Email"] {
		email = value
		break
	}
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := paymentHandler.paymentUsecase.ConfirmPayment(payment, email)
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
