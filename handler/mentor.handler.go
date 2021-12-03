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

type mentorHandler struct {
	mentorUsecase usecase.MentorUsecase
}

type MentorHandler interface {
	GetMentor(ctx *gin.Context)
	GetAllMentor(ctx *gin.Context)

	Mentor(ctx *gin.Context)
}

func InitMentorHandler(mentorUsecase usecase.MentorUsecase) MentorHandler {
	return &mentorHandler{
		mentorUsecase,
	}
}

func (h *mentorHandler) Mentor(ctx *gin.Context) {
	var request model.MentorRequest
	if err := ctx.ShouldBindJSON(&request); err == nil {
		switch strings.ToUpper(request.Action) {
		case utils.REGISTER_MENTOR:
			h.registerMentor(ctx, request)
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

func (h *mentorHandler) registerMentor(ctx *gin.Context, r model.MentorRequest) {
	result, err := h.mentorUsecase.RegisterMentor(r.Data)
	utils.Response(ctx, result, 1, "", err)
}

func (mentorHandler *mentorHandler) GetMentor(ctx *gin.Context) {
	email := ctx.Param("email")

	result, err := mentorHandler.mentorUsecase.GetMentorInfo(email)
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

func (mentorHandler *mentorHandler) GetAllMentor(ctx *gin.Context) {
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

		var mentorResult []shape.MentorInfo
		mentorResult, count, err = mentorHandler.mentorUsecase.GetAllMentor(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":  mentorResult,
				"count": count,
				"error": "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":  mentorResult,
				"count": count,
				"error": err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}
