package handler

import (
	"belajariah-main-service/model"
	"belajariah-main-service/usecase"
	"belajariah-main-service/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type countryCodeHandler struct {
	countryCodeUsecase usecase.CountryCodeUsecase
}

type CountryCodeHandler interface {
	CountryCode(ctx *gin.Context)
}

func InitCountryCodeHandler(u usecase.CountryCodeUsecase) CountryCodeHandler {
	return &countryCodeHandler{
		u,
	}
}

func (h *countryCodeHandler) CountryCode(ctx *gin.Context) {
	var request model.CountryCodeRequest
	if err := ctx.ShouldBindJSON(&request); err == nil {
		switch strings.ToUpper(request.Action) {
		case utils.GET_ALL_COUNTRY_CODE:
			h.getAllCountryCode(ctx, request)
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

func (h *countryCodeHandler) getAllCountryCode(ctx *gin.Context, r model.CountryCodeRequest) {
	result, err := h.countryCodeUsecase.GetAllCountryCode(r)
	if err != nil {
		utils.PushLogStackTrace("", utils.UnwrapError(err))
		utils.Response(ctx, struct{}{}, 0, "", err)
		return
	}

	utils.Response(ctx, result, len(result), "", nil)
}
