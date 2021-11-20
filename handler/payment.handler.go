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

type paymentHandler struct {
	paymentUsecase usecase.PaymentUsecase
}

type PaymentHandler interface {
	Payment(ctx *gin.Context)
	GetAllPayment(ctx *gin.Context)
	GetAllPaymentRejected(ctx *gin.Context)
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

func (h *paymentHandler) Payment(ctx *gin.Context) {
	var request shape.PaymentnRequest
	if err := ctx.ShouldBindJSON(&request); err == nil {
		switch strings.ToUpper(request.Action) {
		case utils.INSERT_PAYMENT_QURAN:
			h.insertPaymentQuran(ctx, request)
		case utils.UPLOAD_PAYMENT_QURAN:
			h.uploadPaymentQuran(ctx, request)
		case utils.CONFIRM_PAYMENT_QURAN:
			h.confirmPaymentQuran(ctx, request)
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

func (h *paymentHandler) insertPaymentQuran(ctx *gin.Context, r shape.PaymentnRequest) {
	result, _, err := h.paymentUsecase.InsertPaymentQuran(ctx, r.Data)
	utils.Response(ctx, result, 1, err)
}

func (h *paymentHandler) uploadPaymentQuran(ctx *gin.Context, r shape.PaymentnRequest) {
	result, err := h.paymentUsecase.UploadPaymentQuran(ctx, r.Data)
	utils.Response(ctx, result, 1, err)
}

func (h *paymentHandler) confirmPaymentQuran(ctx *gin.Context, r shape.PaymentnRequest) {
	result, err := h.paymentUsecase.ConfirmPaymentQuran(ctx, r.Data)
	utils.Response(ctx, result, 1, err)
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
				"data":  paymentResult,
				"count": count,
				"error": "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":  paymentResult,
				"count": count,
				"error": err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}

func (paymentHandler *paymentHandler) GetAllPaymentRejected(ctx *gin.Context) {
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
		paymentResult, count, err = paymentHandler.paymentUsecase.GetAllPaymentRejected(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":  paymentResult,
				"count": count,
				"error": "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":  paymentResult,
				"count": count,
				"error": err.Error(),
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

	var userObj model.UserHeader
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
				"data":  paymentResult,
				"count": count,
				"error": "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":  paymentResult,
				"count": count,
				"error": err.Error(),
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

	payments, result, err := paymentHandler.paymentUsecase.InsertPayment(payment, email)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": result,
			"data":   payments,
			"error":  "",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": result,
			"data":   payments,
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
