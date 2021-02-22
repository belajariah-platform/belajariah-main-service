package handler

import (
	"belajariah-main-service/shape"
	"belajariah-main-service/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userExerciseReadingHandler struct {
	userExerciseReadingUsecase usecase.UserExerciseReadingUsecase
}

type UserExerciseReadingHandler interface {
	InserteUserExerciseReading(ctx *gin.Context)
}

func InitUserExerciseReadingHandler(userExerciseReadingUsecase usecase.UserExerciseReadingUsecase) UserExerciseReadingHandler {
	return &userExerciseReadingHandler{
		userExerciseReadingUsecase,
	}
}

func (userExerciseReadingHandler *userExerciseReadingHandler) InserteUserExerciseReading(ctx *gin.Context) {
	var userExercise shape.UserExerciseReading
	var email string
	for _, value := range ctx.Request.Header["Email"] {
		email = value
		break
	}
	if err := ctx.ShouldBindJSON(&userExercise); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := userExerciseReadingHandler.userExerciseReadingUsecase.InserteUserExerciseReading(userExercise, email)
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
