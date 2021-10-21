package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
)

type paymentMethodUsecase struct {
	paymentMethodRepository repository.PaymentMethodRepository
}

type PaymentMethodUsecase interface {
	GetAllPaymentMethod(query model.Query) ([]shape.PaymentMethod, int, error)
}

func InitPaymentMethodUsecase(paymentMethodRepository repository.PaymentMethodRepository) PaymentMethodUsecase {
	return &paymentMethodUsecase{
		paymentMethodRepository,
	}
}

func (paymentMethodUsecase *paymentMethodUsecase) GetAllPaymentMethod(query model.Query) ([]shape.PaymentMethod, int, error) {
	var filterQuery string
	var paymentMethods []model.PaymentMethod
	var paymentMethodResult []shape.PaymentMethod

	filterQuery = utils.GetFilterHandler(query.Filters)

	paymentMethods, err := paymentMethodUsecase.paymentMethodRepository.GetAllPaymentMethod(query.Skip, query.Take, filterQuery)
	count, errCount := paymentMethodUsecase.paymentMethodRepository.GetAllPaymentMethodCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range paymentMethods {
			paymentMethodResult = append(paymentMethodResult, shape.PaymentMethod{
				ID:             value.ID,
				Code:           value.Code,
				Type:           value.Type,
				Value:          value.Value,
				Account_Name:   value.AccountName.String,
				Account_Number: value.AccountNumber.String,
				Icon_Account:   value.IconAccount.String,
				Is_Active:      value.IsActive,
				Created_By:     value.CreatedBy,
				Created_Date:   value.CreatedDate,
				Modified_By:    value.ModifiedBy.String,
				Modified_Date:  value.ModifiedDate.Time,
				Is_Deleted:     value.IsDeleted,
			})
		}
	}
	paymentMethodEmpty := make([]shape.PaymentMethod, 0)
	if len(paymentMethodResult) == 0 {
		return paymentMethodEmpty, count, err
	}
	return paymentMethodResult, count, err
}
