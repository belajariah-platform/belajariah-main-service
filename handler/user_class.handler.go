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

type userClassHandler struct {
	userClassUsecase usecase.UserClassUsecase
}

type UserClassHandler interface {
	GetAllUserClass(ctx *gin.Context)
	UpdateUserClassProgress(ctx *gin.Context)
}

func InitUserClassHandler(userClassUsecase usecase.UserClassUsecase) UserClassHandler {
	return &userClassHandler{
		userClassUsecase,
	}
}

func (userClassHandler *userClassHandler) GetAllUserClass(ctx *gin.Context) {
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

		var userClassResult []shape.UserClass
		userClassResult, count, err = userClassHandler.userClassUsecase.GetAllUserClass(query, userObj)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":      userClassResult,
				"dataCount": count,
				"error":     "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":      userClassResult,
				"dataCount": count,
				"error":     err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}

func (userClassHandler *userClassHandler) UpdateUserClassProgress(ctx *gin.Context) {
	var userClass shape.UserClassPost
	var email string
	for _, value := range ctx.Request.Header["Email"] {
		email = value
		break
	}
	if err := ctx.ShouldBindJSON(&userClass); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := userClassHandler.userClassUsecase.UpdateUserClassProgress(userClass, email)
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
