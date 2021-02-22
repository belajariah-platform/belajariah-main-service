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

type consultationHandler struct {
	consultationUsecase usecase.ConsultationUsecase
}

type ConsultationHandler interface {
	GetAllConsultation(ctx *gin.Context)
	GetAllConsultationUser(ctx *gin.Context)
	GetAllConsultationMentor(ctx *gin.Context)

	ReadConsultation(ctx *gin.Context)
	InsertConsultation(ctx *gin.Context)
	UpdateConsultation(ctx *gin.Context)
	ConfirmConsultation(ctx *gin.Context)
}

func InitConsultationHandler(consultationUsecase usecase.ConsultationUsecase) ConsultationHandler {
	return &consultationHandler{
		consultationUsecase,
	}
}

func (consultationHandler *consultationHandler) GetAllConsultation(ctx *gin.Context) {
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

		var consultationResult []shape.Consultation
		consultationResult, count, err := consultationHandler.consultationUsecase.GetAllConsultation(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":      consultationResult,
				"dataCount": count,
				"error":     "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":      consultationResult,
				"dataCount": count,
				"error":     err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}

func (consultationHandler *consultationHandler) GetAllConsultationUser(ctx *gin.Context) {
	var query model.Query
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

		var consultationResult []shape.Consultation
		consultationResult, count, err := consultationHandler.consultationUsecase.GetAllConsultationUser(query, userObj)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":      consultationResult,
				"dataCount": count,
				"error":     "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":      consultationResult,
				"dataCount": count,
				"error":     err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}

func (consultationHandler *consultationHandler) GetAllConsultationMentor(ctx *gin.Context) {
	var query model.Query
	err := ctx.BindQuery(&query)

	var userObj model.Mentor
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

		var consultationResult []shape.Consultation
		consultationResult, count, err := consultationHandler.consultationUsecase.GetAllConsultationMentor(query, userObj)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":      consultationResult,
				"dataCount": count,
				"error":     "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":      consultationResult,
				"dataCount": count,
				"error":     err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}

func (consultationHandler *consultationHandler) ReadConsultation(ctx *gin.Context) {
	var consultation shape.ConsultationPost
	var email string
	for _, value := range ctx.Request.Header["Email"] {
		email = value
		break
	}
	if err := ctx.ShouldBindJSON(&consultation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := consultationHandler.consultationUsecase.ReadConsultation(consultation, email)
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

func (consultationHandler *consultationHandler) InsertConsultation(ctx *gin.Context) {
	var consultation shape.ConsultationPost
	var email string
	for _, value := range ctx.Request.Header["Email"] {
		email = value
		break
	}
	if err := ctx.ShouldBindJSON(&consultation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := consultationHandler.consultationUsecase.InsertConsultation(consultation, email)
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

func (consultationHandler *consultationHandler) UpdateConsultation(ctx *gin.Context) {
	var consultation shape.ConsultationPost
	var email string
	for _, value := range ctx.Request.Header["Email"] {
		email = value
		break
	}
	if err := ctx.ShouldBindJSON(&consultation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := consultationHandler.consultationUsecase.UpdateConsultation(consultation, email)
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

func (consultationHandler *consultationHandler) ConfirmConsultation(ctx *gin.Context) {
	var consultation shape.ConsultationPost
	var email string
	for _, value := range ctx.Request.Header["Email"] {
		email = value
		break
	}
	if err := ctx.ShouldBindJSON(&consultation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := consultationHandler.consultationUsecase.ConfirmConsultation(consultation, email)
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
