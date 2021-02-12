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

type ratingHandler struct {
	ratingUsecase usecase.RatingUsecase
}

type RatingHandler interface {
	GetAllRatingClass(ctx *gin.Context)
	GiveRatingClass(ctx *gin.Context)
	GiveRatingMentor(ctx *gin.Context)
}

func InitRatingHandler(ratingUsecase usecase.RatingUsecase) RatingHandler {
	return &ratingHandler{
		ratingUsecase,
	}
}

func (ratingHandler *ratingHandler) GetAllRatingClass(ctx *gin.Context) {
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

		var ratingResult []shape.Rating
		ratingResult, count, err = ratingHandler.ratingUsecase.GetAllRatingClass(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":      ratingResult,
				"dataCount": count,
				"error":     "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":      ratingResult,
				"dataCount": count,
				"error":     err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}

func (ratingHandler *ratingHandler) GiveRatingClass(ctx *gin.Context) {
	var rating shape.RatingPost
	var email string
	for _, value := range ctx.Request.Header["Email"] {
		email = value
		break
	}
	if err := ctx.ShouldBindJSON(&rating); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := ratingHandler.ratingUsecase.GiveRatingClass(rating, email)
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

func (ratingHandler *ratingHandler) GiveRatingMentor(ctx *gin.Context) {
	var rating shape.RatingPost
	var email string
	for _, value := range ctx.Request.Header["Email"] {
		email = value
		break
	}
	if err := ctx.ShouldBindJSON(&rating); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := ratingHandler.ratingUsecase.GiveRatingMentor(rating, email)
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
