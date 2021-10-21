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

type quranHandler struct {
	quranUsecase usecase.QuranUsecase
}

type QuranHandler interface {
	GetAllQuran(ctx *gin.Context)
	GetAllAyatQuran(ctx *gin.Context)
	GetAllQuranView(ctx *gin.Context)
}

func InitQuranHandler(quranUsecase usecase.QuranUsecase) QuranHandler {
	return &quranHandler{
		quranUsecase,
	}
}

func (quranHandler *quranHandler) GetAllQuran(ctx *gin.Context) {
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

		var ratingResult []shape.Quran
		ratingResult, count, err = quranHandler.quranUsecase.GetAllQuran(query)
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

func (quranHandler *quranHandler) GetAllAyatQuran(ctx *gin.Context) {
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

		var ratingResult []shape.Quran
		ratingResult, count, err = quranHandler.quranUsecase.GetAllAyatQuran(query)
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

func (quranHandler *quranHandler) GetAllQuranView(ctx *gin.Context) {
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

		var ratingResult []shape.Quran
		ratingResult, count, err = quranHandler.quranUsecase.GetAllQuranView(query)
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
