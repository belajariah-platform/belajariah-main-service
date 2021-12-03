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

type ratingHandler struct {
	ratingUsecase usecase.RatingUsecase
}

type RatingHandler interface {
	Rating(ctx *gin.Context)
	GetAllRatingClass(ctx *gin.Context)
	GiveRatingClass(ctx *gin.Context)
	GiveRatingMentor(ctx *gin.Context)
}

func InitRatingHandler(ratingUsecase usecase.RatingUsecase) RatingHandler {
	return &ratingHandler{
		ratingUsecase,
	}
}

func (h *ratingHandler) Rating(ctx *gin.Context) {
	var request model.RatingQuranRequest
	if err := ctx.ShouldBindJSON(&request); err == nil {
		switch strings.ToUpper(request.Action) {
		case utils.GET_ALL_RATING_CLASS_QURAN:
			h.getAllRatingClassQuran(ctx, request)
		case utils.GIVE_RATING_CLASS_QURAN:
			h.giveRatingClassQuran(ctx, request)
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

func (h *ratingHandler) getAllRatingClassQuran(ctx *gin.Context, r model.RatingQuranRequest) {
	result, count, err := h.ratingUsecase.GetAllRatingClassQuran(r)
	if err != nil {
		utils.PushLogStackTrace("", utils.UnwrapError(err))
		utils.Response(ctx, struct{}{}, 0, "", err)
		return
	}

	utils.Response(ctx, result, count, "", nil)
}

func (h *ratingHandler) giveRatingClassQuran(ctx *gin.Context, r model.RatingQuranRequest) {
	result, err := h.ratingUsecase.GiveRatingClassQuran(ctx, r)
	utils.Response(ctx, result, 1, "", err)
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
				"data":  ratingResult,
				"count": count,
				"error": "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":  ratingResult,
				"count": count,
				"error": err.Error(),
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
