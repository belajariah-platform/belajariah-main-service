package handler

import (
	"belajariah-main-service/model"
	"belajariah-main-service/usecase"
	"belajariah-main-service/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type promotionHandler struct {
	promotionUsecase usecase.PromotionUsecase
}

type PromotionHandler interface {
	Promotion(ctx *gin.Context)
}

func InitPromotionHandler(promotionUsecase usecase.PromotionUsecase) PromotionHandler {
	return &promotionHandler{
		promotionUsecase,
	}
}

func (h *promotionHandler) Promotion(ctx *gin.Context) {
	var request model.PromotionRequest
	if err := ctx.ShouldBindJSON(&request); err == nil {
		switch strings.ToUpper(request.Action) {
		case utils.GET_ALL_PROMOTION:
			h.getAllPromotions(ctx, request)
		case utils.GET_ALL_PROMOTION_HEADER:
			h.getAllPromotionHeader(ctx, request)
		case utils.CLAIM_PROMOTION:
			h.claimPromotion(ctx, request)
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

func (h *promotionHandler) getAllPromotions(ctx *gin.Context, r model.PromotionRequest) {
	result, err := h.promotionUsecase.GetAllPromotion(r)
	if err != nil {
		utils.PushLogStackTrace("", utils.UnwrapError(err))
		utils.Response(ctx, struct{}{}, 0, "", err)
		return
	}

	utils.Response(ctx, result, len(result), "", nil)
}

func (h *promotionHandler) getAllPromotionHeader(ctx *gin.Context, r model.PromotionRequest) {
	result, err := h.promotionUsecase.GetAllPromotionHeader(r)
	if err != nil {
		utils.PushLogStackTrace("", utils.UnwrapError(err))
		utils.Response(ctx, struct{}{}, 0, "", err)
		return
	}

	utils.Response(ctx, result, len(result), "", nil)
}
func (h *promotionHandler) claimPromotion(ctx *gin.Context, r model.PromotionRequest) {
	result, message, err := h.promotionUsecase.ClaimPromotion(ctx, r)
	utils.Response(ctx, result, 1, message, err)
}
