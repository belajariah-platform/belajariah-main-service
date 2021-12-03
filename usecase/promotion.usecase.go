package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type promotionUsecase struct {
	sytemConfig         *model.Config
	userRepository      repository.UserRepository
	promotionRepository repository.PromotionRepository
	userClassRepository repository.UserClassRepository
	paymentRepository   repository.PaymentsRepository
}

type PromotionUsecase interface {
	ClaimPromotion(ctx *gin.Context, r model.PromotionRequest) (model.Promotion, string, error)
	GetAllPromotionHeader(r model.PromotionRequest) ([]model.Promotion, error)
	GetAllPromotion(r model.PromotionRequest) ([]model.Promotion, error)
	CheckAllPromotionExpired()
}

func InitPromotionUsecase(sytemConfig *model.Config, userRepository repository.UserRepository, promotionRepository repository.PromotionRepository, userClassRepository repository.UserClassRepository, paymentRepository repository.PaymentsRepository) PromotionUsecase {
	return &promotionUsecase{
		sytemConfig,
		userRepository,
		promotionRepository,
		userClassRepository,
		paymentRepository,
	}
}

func (u *promotionUsecase) GetAllPromotion(r model.PromotionRequest) ([]model.Promotion, error) {
	var orderDefault = "ORDER BY code asc"
	var filterDefault = fmt.Sprintf(`is_deleted=false AND is_active=true AND class_code='%s' 
		AND (code = '%s' OR promo_code = '%s')`, r.Data.ClassCode, r.Data.Code, r.Data.PromoCode)

	promoEmpty := make([]model.Promotion, 0)

	filterFinal := utils.GetFilterOrderHandler(filterDefault, orderDefault, model.Query{})

	result, err := u.promotionRepository.GetAllPromotions(filterFinal)
	if err != nil {
		return nil, utils.WrapError(err, "promotionUsecase.GetAllPromotions")
	}

	if len(*result) == 0 {
		return promoEmpty, err
	}

	return *result, nil
}

func (u *promotionUsecase) GetAllPromotionHeader(r model.PromotionRequest) ([]model.Promotion, error) {
	var orderDefault = "ORDER BY code asc"
	var filterDefault = "is_deleted = false and is_active = true"
	promoEmpty := make([]model.Promotion, 0)

	filterFinal := utils.GetFilterOrderHandler(filterDefault, orderDefault, r.Query)

	result, err := u.promotionRepository.GetAllPromotionHeader(filterFinal)
	if err != nil {
		return promoEmpty, utils.WrapError(err, "promotionUsecase.GetAllPromotionHeader")
	}

	if len(*result) == 0 {
		return promoEmpty, err
	}

	return *result, nil
}

func (u *promotionUsecase) ClaimPromotion(ctx *gin.Context, r model.PromotionRequest) (model.Promotion, string, error) {
	var message string
	var result model.Promotion

	email := ctx.Request.Header.Get("email")

	users, err := u.userRepository.GetUserInfo(email)
	if err != nil {
		return result, message, utils.WrapError(err, "userClassUsecase.GetUserInfo")
	}

	var filter = fmt.Sprintf(`AND user_code='%s' AND class_code='%s'`,
		users.Code,
		r.Data.ClassCode,
	)

	var filterDefault = fmt.Sprintf(`WHERE is_deleted=false AND is_active=true 
		AND class_code='%s' AND (code = '%s' OR promo_code = '%s') LIMIT 1`,
		r.Data.ClassCode,
		r.Data.Code,
		r.Data.PromoCode,
	)

	var filterPayment = fmt.Sprintf(`AND promo_code ='%s' AND status_payment in 
		('Waiting for Payment', 'Has been Payment', 'Completed')`,
		r.Data.PromoCode,
	)

	promotion, err := u.promotionRepository.GetAllPromotions(filterDefault)
	if err != nil {
		return result, message, utils.WrapError(err, "promotionUsecase.promotionRepository.GetAllPromotions : ")
	}

	if len(*promotion) == 0 {
		message = "Mohon maaf kode promo tidak berlaku"
	} else {
		count, _ := u.paymentRepository.GetAllPaymentCount(filter, filterPayment)

		for _, promo := range *promotion {
			if count != 0 {
				message = "Mohon maaf kode promo sudah pernah digunakan"
			} else if promo.QuotaUsed >= promo.QuotaUser {
				message = "Mohon maaf kuota promo sudah penuh"
			} else if promo.QuotaUsed < promo.QuotaUser && promo.PromoType == "Extend Promo" {
				class, _ := u.userClassRepository.GetUserClass(filter)
				if class == (model.UserClass{}) {
					message = "Mohon maaf anda belum bisa menggunakan kode promo ini, Ayo berlangganan kelas dahulu"
				} else if utils.GetDuration(class.ExpiredDate.Time, time.Now()) >= 10080 {
					message = "Mohon maaf anda belum bisa menggunakan kode promo ini"
				} else {
					message = ""
					result.Discount = promo.Discount
				}
			} else {
				message = ""
				result.Discount = promo.Discount
			}
		}
	}

	return result, message, err
}

func (promotionUsecase *promotionUsecase) CheckAllPromotionExpired() {
	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}
		promotionList, err := promotionUsecase.promotionRepository.CheckAllPromotionExpired()
		if err != nil {
			utils.PushLogf("promotionUsecase.promotionRepository.CheckAllPromotionExpired : ", err.Error())
		}

		firstloop = false
		if err == nil {
			for _, value := range promotionList {
				dataPromotion := model.Promotion{
					Code:         value.Code,
					ModifiedBy:   promotionUsecase.sytemConfig.System.EmailSystem,
					ModifiedDate: time.Now(),
				}
				_, err := promotionUsecase.promotionRepository.UpdatePromotionActivated(dataPromotion)
				if err != nil {
					utils.PushLogf("promotionUsecase.promotionRepository.UpdatePromotionActivated : ", err.Error())
				}
			}
		}
	}
}
