package utils

import (
	"belajariah-main-service/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, data interface{}, count int, err error) {
	if err == nil {
		ctx.JSON(http.StatusOK, model.Response{
			Message: model.RequestResponse{
				Data:   data,
				Count:  count,
				Result: true,
			},
			Status: http.StatusOK,
		})
	} else {
		PushLogStackTrace("", UnwrapError(err))
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

func NotFoundActionResponse(ctx *gin.Context, action string) {
	ctx.JSON(http.StatusNotFound, model.Response{
		Message: model.RequestResponse{
			Count:  0,
			Data:   nil,
			Error:  fmt.Sprintf("No action `%s` is available", action),
			Result: false,
		},
		Status: http.StatusBadRequest,
		Error:  fmt.Sprintf("No action `%s` is available", action),
	})
}
