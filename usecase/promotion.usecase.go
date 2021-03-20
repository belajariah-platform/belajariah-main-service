package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"
)

type promotionUsecase struct {
	promotionRepository repository.PromotionRepository
	userClassRepository repository.UserClassRepository
	paymentRepository   repository.PaymentsRepository
}

type PromotionUsecase interface {
	ClaimPromotion(promotion shape.PromotionClaim, userObj model.UserHeader) (shape.Promotion, string, error)
	GetAllPromotion(query model.Query) ([]shape.Promotion, int, error)
	GetPromotion(code string) (shape.Promotion, error)
	CheckAllPromotionExpired()
}

func InitPromotionUsecase(promotionRepository repository.PromotionRepository, userClassRepository repository.UserClassRepository, paymentRepository repository.PaymentsRepository) PromotionUsecase {
	return &promotionUsecase{
		promotionRepository,
		userClassRepository,
		paymentRepository,
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
				ID:              value.ID,
				Code:            value.Code,
				Class_Code:      value.ClassCode,
				Title:           value.Title,
				Description:     value.Description.String,
				Promo_Code:      value.PromoCode,
				Promo_Type_Code: value.PromoTypeCode.String,
				Promo_Type:      value.PromoType.String,
				Discount:        value.Discount.Float64,
				Banner_Image:    value.BannerImage.String,
				Header_Image:    value.HeaderImage.String,
				Expired_Date:    utils.HandleNullableDate(value.ExpiredDate.Time),
				Quota_User:      int(value.QuotaUser.Int64),
				Quota_Used:      int(value.QuotaUsed.Int64),
				Is_Active:       value.IsActive,
				Created_By:      value.CreatedBy,
				Created_Date:    value.CreatedDate,
				Modified_By:     value.ModifiedBy.String,
				Modified_Date:   value.ModifiedDate.Time,
				Deleted_By:      value.DeletedBy.String,
				Deleted_Date:    value.DeletedDate.Time,
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
		ID:              promotion.ID,
		Code:            promotion.Code,
		Class_Code:      promotion.ClassCode,
		Title:           promotion.Title,
		Description:     promotion.Description.String,
		Promo_Code:      promotion.PromoCode,
		Promo_Type:      promotion.PromoType.String,
		Promo_Type_Code: promotion.PromoTypeCode.String,
		Discount:        promotion.Discount.Float64,
		Banner_Image:    promotion.BannerImage.String,
		Header_Image:    promotion.HeaderImage.String,
		Expired_Date:    utils.HandleNullableDate(promotion.ExpiredDate.Time),
		Quota_User:      int(promotion.QuotaUser.Int64),
		Quota_Used:      int(promotion.QuotaUsed.Int64),
		Is_Active:       promotion.IsActive,
		Created_By:      promotion.CreatedBy,
		Created_Date:    promotion.CreatedDate,
		Modified_By:     promotion.ModifiedBy.String,
		Modified_Date:   promotion.ModifiedDate.Time,
		Deleted_By:      promotion.DeletedBy.String,
		Deleted_Date:    promotion.DeletedDate.Time,
	}
	return promotionResult, err
}

func (promotionUsecase *promotionUsecase) ClaimPromotion(promotions shape.PromotionClaim, userObj model.UserHeader) (shape.Promotion, string, error) {
	var message string
	var dataPromotion shape.Promotion
	filter := fmt.Sprintf(`AND user_code=%d AND class_code='%s'`,
		userObj.ID,
		promotions.Class_Code,
	)

	promotion, err := promotionUsecase.promotionRepository.GetPromotion(promotions.Promo_Code)
	if promotion == (model.Promotion{}) {
		message = "Mohon maaf kode promo tidak berlaku"
	} else {
		filters := fmt.Sprintf(`AND promo_code ='%s' AND status_payment in 
		('Waiting for Payment', 'Has been Payment', 'Completed')`, promotion.Code)
		count, _ := promotionUsecase.paymentRepository.GetAllPaymentCount(filter, filters)
		if count != 0 {
			message = "Mohon maaf kode promo sudah pernah digunakan"
		} else if promotion.QuotaUsed.Int64 >= promotion.QuotaUser.Int64 {
			message = "Mohon maaf kuota promo sudah penuh"
		} else if promotion.QuotaUsed.Int64 < promotion.QuotaUser.Int64 &&
			promotion.PromoType.String == "Extend Promo" {
			class, _ := promotionUsecase.userClassRepository.GetUserClass(filter)
			if class == (model.UserClass{}) {
				message = "Mohon maaf anda belum bisa menggunakan kode promo ini, Ayo berlangganan kelas dahulu"
			} else if utils.GetDuration(class.ExpiredDate, time.Now()) >= 10080 {
				message = "Mohon maaf anda belum bisa menggunakan kode promo ini"
			} else {
				message = ""
				dataPromotion.Discount = promotion.Discount.Float64
			}
		} else {
			message = ""
			dataPromotion.Discount = promotion.Discount.Float64
		}
	}

	return dataPromotion, message, err
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
