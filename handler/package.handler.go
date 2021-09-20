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

type packageHandler struct {
	packageUsecase usecase.PackageUsecase
}

type PackageHandler interface {
	GetAllPackage(ctx *gin.Context)
	GetAllBenefit(ctx *gin.Context)
}

func InitPackageHandler(packageUsecase usecase.PackageUsecase) PackageHandler {
	return &packageHandler{
		packageUsecase,
	}
}

func (packageHandler *packageHandler) GetAllPackage(ctx *gin.Context) {
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

		var packageResult []shape.Package
		packageResult, count, err = packageHandler.packageUsecase.GetAllPackage(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":      packageResult,
				"dataCount": count,
				"error":     "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":      packageResult,
				"dataCount": count,
				"error":     err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}

func (packageHandler *packageHandler) GetAllBenefit(ctx *gin.Context) {
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

		var benefitResult []shape.Benefit
		benefitResult, err = packageHandler.packageUsecase.GetAllBenefit(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":  benefitResult,
				"error": "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":  benefitResult,
				"error": err.Error(),
			})
		}
	} else {
		utils.PushLogf("err", err)
	}
}
