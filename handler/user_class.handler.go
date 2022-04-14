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

type userClassHandler struct {
	userClassUsecase usecase.UserClassUsecase
}

type UserClassHandler interface {
	UserClass(ctx *gin.Context)
	GetUserClass(ctx *gin.Context)
	GetAllUserClass(ctx *gin.Context)
}

func InitUserClassHandler(userClassUsecase usecase.UserClassUsecase) UserClassHandler {
	return &userClassHandler{
		userClassUsecase,
	}
}

func (h *userClassHandler) UserClass(ctx *gin.Context) {
	var request model.UserClassQuranRequest
	if err := ctx.ShouldBindJSON(&request); err == nil {
		switch strings.ToUpper(request.Action) {
		case utils.GET_ALL_USER_CLASS_QURAN:
			h.getAllUserClassQuran(ctx, request)
		case utils.GET_ALL_USER_CLASS_DETAIL_QURAN:
			h.getAllUserClassDetailQuran(ctx, request)
		case utils.GET_ALL_USER_CLASS_SCHEDULE_QURAN:
			h.getAllUserClassScheduleQuran(ctx, request)
		case utils.UPDATE_USER_CLASS_QURAN_PROGRESS:
			h.updateUserClassQuranProgress(ctx, request)
		case utils.INSERT_USER_CLASS_QURAN_SCHEDULE:
			h.insertUserClassQuranSchedule(ctx, request)
		case utils.UPDATE_USER_CLASS_QURAN_SCHEDULE:
			h.updateUserClassQuranSchedule(ctx, request)
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

func (h *userClassHandler) getAllUserClassQuran(ctx *gin.Context, r model.UserClassQuranRequest) {
	result, count, err := h.userClassUsecase.GetAllUserClassQuran(ctx, r)
	if err != nil {
		utils.PushLogStackTrace("", utils.UnwrapError(err))
		utils.Response(ctx, struct{}{}, 0, "", err)
		return
	}

	utils.Response(ctx, result, count, "", nil)
}

func (h *userClassHandler) getAllUserClassDetailQuran(ctx *gin.Context, r model.UserClassQuranRequest) {
	result, count, err := h.userClassUsecase.GetAllUserClassQuranDetail(ctx, r)
	if err != nil {
		utils.PushLogStackTrace("", utils.UnwrapError(err))
		utils.Response(ctx, struct{}{}, 0, "", err)
		return
	}

	utils.Response(ctx, result, count, "", nil)
}

func (h *userClassHandler) getAllUserClassScheduleQuran(ctx *gin.Context, r model.UserClassQuranRequest) {
	result, count, err := h.userClassUsecase.GetAllUserClassQuranSchedule(ctx, r)
	if err != nil {
		utils.PushLogStackTrace("", utils.UnwrapError(err))
		utils.Response(ctx, struct{}{}, 0, "", err)
		return
	}

	utils.Response(ctx, result, count, "", nil)
}

func (h *userClassHandler) updateUserClassQuranProgress(ctx *gin.Context, r model.UserClassQuranRequest) {
	result, err := h.userClassUsecase.UpdateUserClassQuranProgress(ctx, r)
	utils.Response(ctx, result, 1, "", err)
}

func (h *userClassHandler) insertUserClassQuranSchedule(ctx *gin.Context, r model.UserClassQuranRequest) {
	result, err := h.userClassUsecase.InsertUserClassQuranSchedule(ctx, r)
	utils.Response(ctx, result, 1, "", err)
}

func (h *userClassHandler) updateUserClassQuranSchedule(ctx *gin.Context, r model.UserClassQuranRequest) {
	result, err := h.userClassUsecase.UpdateUserClassQuranSchedule(ctx, r)
	utils.Response(ctx, result, 1, "", err)
}

func (userClassHandler *userClassHandler) GetUserClass(ctx *gin.Context) {
	code := ctx.Param("code")

	var userObj model.UserHeader
	for _, valueUser := range ctx.Request.Header["User"] {
		itemInfoBytes := []byte(valueUser)
		er := json.Unmarshal(itemInfoBytes, &userObj)
		if er != nil {
			utils.PushLogf("[Error Unmarshal] :", er)
		}
	}

	result, err := userClassHandler.userClassUsecase.GetUserClass(code, userObj)
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

func (userClassHandler *userClassHandler) GetAllUserClass(ctx *gin.Context) {
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

		var userClassResult []shape.UserClass
		userClassResult, count, err = userClassHandler.userClassUsecase.GetAllUserClass(query, userObj)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":  userClassResult,
				"count": count,
				"error": "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":  userClassResult,
				"count": count,
				"error": err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}
