package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
	"database/sql"
	"time"
)

type promotionUsecase struct {
	promotionRepository repository.PromotionRepository
}

type PromotionUsecase interface {
	CheckAllPromotionExpired()
	GetAllPromotion(query model.Query) ([]shape.Promotion, int, error)
	GetPromotion(code string) (shape.Promotion, error)
}

func InitPromotionUsecase(promotionRepository repository.PromotionRepository) PromotionUsecase {
	return &promotionUsecase{
		promotionRepository,
	}
}

func (promotionUsecase *promotionUsecase) GetAllPromotion(query model.Query) ([]shape.Promotion, int, error) {
	var filterQuery string
	var promotions []model.Promotion
	var promotionResult []shape.Promotion

	filterQuery = utils.GetFilterHandler(query.Filters)

	promotions, err := promotionUsecase.promotionRepository.GetAllPromotion(query.Skip, query.Take, filterQuery)
	count, errCount := promotionUsecase.promotionRepository.GetAllPromotionCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range promotions {
			promotionResult = append(promotionResult, shape.Promotion{
				ID:            value.ID,
				Code:          value.Code,
				Class_Code:    value.ClassCode,
				Title:         value.Title,
				Description:   value.Description.String,
				Promo_Code:    value.PromoCode,
				Discount:      value.Discount.Float64,
				Banner_Image:  value.BannerImage.String,
				Header_Image:  value.HeaderImage.String,
				Expired_Date:  utils.HandleNullableDate(value.ExpiredDate.Time),
				Quota_User:    int(value.QuotaUser.Int64),
				Quota_Used:    int(value.QuotaUsed.Int64),
				Is_Active:     value.IsActive,
				Created_By:    value.CreatedBy,
				Created_Date:  value.CreatedDate,
				Modified_By:   value.ModifiedBy.String,
				Modified_Date: value.ModifiedDate.Time,
				Deleted_By:    value.DeletedBy.String,
				Deleted_Date:  value.DeletedDate.Time,
			})
		}
	}
	promotionEmpty := make([]shape.Promotion, 0)
	if len(promotionResult) == 0 {
		return promotionEmpty, count, err
	}
	return promotionResult, count, err
}

func (promotionUsecase *promotionUsecase) GetPromotion(code string) (shape.Promotion, error) {
	promotion, err := promotionUsecase.promotionRepository.GetPromotion(code)
	if promotion == (model.Promotion{}) {
		return shape.Promotion{}, nil
	}
	promotionResult := shape.Promotion{
		ID:            promotion.ID,
		Code:          promotion.Code,
		Class_Code:    promotion.ClassCode,
		Title:         promotion.Title,
		Description:   promotion.Description.String,
		Promo_Code:    promotion.PromoCode,
		Discount:      promotion.Discount.Float64,
		Banner_Image:  promotion.BannerImage.String,
		Header_Image:  promotion.HeaderImage.String,
		Expired_Date:  utils.HandleNullableDate(promotion.ExpiredDate.Time),
		Quota_User:    int(promotion.QuotaUser.Int64),
		Quota_Used:    int(promotion.QuotaUsed.Int64),
		Is_Active:     promotion.IsActive,
		Created_By:    promotion.CreatedBy,
		Created_Date:  promotion.CreatedDate,
		Modified_By:   promotion.ModifiedBy.String,
		Modified_Date: promotion.ModifiedDate.Time,
		Deleted_By:    promotion.DeletedBy.String,
		Deleted_Date:  promotion.DeletedDate.Time,
	}
	return promotionResult, err
}

func (promotionUsecase *promotionUsecase) CheckAllPromotionExpired() {
	var email = "belajariah20@gmail.com"

	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}
		promotionList, err := promotionUsecase.promotionRepository.CheckAllPromotionExpired()
		firstloop = false
		if err == nil {
			for _, value := range promotionList {
				dataPromotion := model.Promotion{
					ID: value.ID,
					ModifiedBy: sql.NullString{
						String: email,
					},
					ModifiedDate: sql.NullTime{
						Time: time.Now(),
					},
				}
				result, err := promotionUsecase.promotionRepository.UpdatePromotionActivated(dataPromotion)
				if err != nil {
					utils.PushLogf("ERROR : ", err, result)
				}
			}
		}
	}
}
