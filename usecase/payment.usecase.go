package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type paymentUsecase struct {
	enumRepository           repository.EnumRepository
	paymentsRepository       repository.PaymentsRepository
	approvalStatusRepository repository.ApprovalStatusRepository
}

type PaymentUsecase interface {
	CheckAllPaymentExpired()
	GetAllPayment(query model.Query) ([]shape.Payment, int, error)
	GetAllPaymentByUserID(query model.Query, userObj model.UserInfo) ([]shape.Payment, int, error)
	InsertPayment(payment shape.PaymentPost, email string) (bool, error)
	UploadPayment(payment shape.PaymentPost, email string) (bool, error)
	ConfirmPayment(payment shape.PaymentPost, email string) (bool, error)
}

func InitPaymentUsecase(enumRepository repository.EnumRepository, paymentsRepository repository.PaymentsRepository, approvalStatusRepository repository.ApprovalStatusRepository) PaymentUsecase {
	return &paymentUsecase{
		enumRepository,
		paymentsRepository,
		approvalStatusRepository,
	}
}

func (paymentUsecase *paymentUsecase) GetAllPayment(query model.Query) ([]shape.Payment, int, error) {
	var filterQuery, filterUser string
	var payments []model.Payment
	var paymentResult []shape.Payment

	filterUser = fmt.Sprintf(``)
	filterQuery = utils.GetFilterHandler(query.Filters)

	payments, err := paymentUsecase.paymentsRepository.GetAllPayment(query.Skip, query.Take, filterQuery, filterUser)
	count, errCount := paymentUsecase.paymentsRepository.GetAllPaymentCount(filterQuery, filterUser)

	if err == nil && errCount == nil {
		for _, value := range payments {
			paymentResult = append(paymentResult, shape.Payment{
				ID:                  value.ID,
				User_Code:           value.UserCode,
				User_Name:           value.UserName,
				Class_Code:          value.ClassCode,
				Class_Name:          value.ClassName,
				Class_Initial:       value.ClassInitial,
				Package_Code:        value.PackageCode,
				Package_Type:        value.PackageType,
				Payment_Method_Code: value.PaymentMethodCode,
				Payment_Method:      value.PaymentMethod,
				Invoice_Number:      value.InvoiceNumber,
				Status_Payment_Code: value.StatusPaymentCode,
				Status_Payment:      value.StatusPayment,
				Total_Transfer:      value.TotalTransfer,
				Sender_Bank:         value.SenderBank.String,
				Sender_Name:         value.SenderName.String,
				Image_Proof:         value.ImageProof.String,
				Payment_Type_Code:   value.PaymentTypeCode,
				Payment_Type:        value.PaymentType,
				Is_Active:           value.IsActive,
				Created_By:          value.CreatedBy,
				Created_Date:        value.CreatedDate,
				Modified_By:         value.ModifiedBy.String,
				Modified_Date:       value.ModifiedDate.Time,
			})
		}
	}
	paymentEmpty := make([]shape.Payment, 0)
	if len(paymentResult) == 0 {
		return paymentEmpty, count, err
	}
	return paymentResult, count, err
}

func (paymentUsecase *paymentUsecase) GetAllPaymentByUserID(query model.Query, userObj model.UserInfo) ([]shape.Payment, int, error) {
	var filterQuery, filterUser string
	var payments []model.Payment
	var paymentResult []shape.Payment

	filterQuery = utils.GetFilterHandler(query.Filters)
	filterUser = fmt.Sprintf(`WHERE user_code=%d`, userObj.ID)

	payments, err := paymentUsecase.paymentsRepository.GetAllPayment(query.Skip, query.Take, filterQuery, filterUser)
	count, errCount := paymentUsecase.paymentsRepository.GetAllPaymentCount(filterQuery, filterUser)

	if err == nil && errCount == nil {
		for _, value := range payments {
			paymentResult = append(paymentResult, shape.Payment{
				ID:                  value.ID,
				User_Code:           value.UserCode,
				User_Name:           value.UserName,
				Class_Code:          value.ClassCode,
				Class_Name:          value.ClassName,
				Class_Initial:       value.ClassInitial,
				Package_Code:        value.PackageCode,
				Package_Type:        value.PackageType,
				Payment_Method_Code: value.PaymentMethodCode,
				Payment_Method:      value.PaymentMethod,
				Invoice_Number:      value.InvoiceNumber,
				Status_Payment_Code: value.StatusPaymentCode,
				Status_Payment:      value.StatusPayment,
				Total_Transfer:      value.TotalTransfer,
				Sender_Bank:         value.SenderBank.String,
				Sender_Name:         value.SenderName.String,
				Image_Proof:         value.ImageProof.String,
				Is_Active:           value.IsActive,
				Created_By:          value.CreatedBy,
				Created_Date:        value.CreatedDate,
				Modified_By:         value.ModifiedBy.String,
				Modified_Date:       value.ModifiedDate.Time,
			})
		}
	}
	paymentEmpty := make([]shape.Payment, 0)
	if len(paymentResult) == 0 {
		return paymentEmpty, count, err
	}
	return paymentResult, count, err
}

func (paymentUsecase *paymentUsecase) InsertPayment(payment shape.PaymentPost, email string) (bool, error) {
	var paymentType = "WaitingForPayment|Waiting for Payment|Menunggu"

	enum, err := paymentUsecase.enumRepository.GetEnum(paymentType)
	dataPayment := model.Payment{
		UserCode:          payment.User_Code,
		ClassCode:         payment.Class_Code,
		PackageCode:       payment.Package_Code,
		PaymentMethodCode: payment.Payment_Method_Code,
		InvoiceNumber:     payment.Invoice_Number,
		StatusPaymentCode: payment.Status_Payment_Code,
		TotalTransfer:     payment.Total_Transfer,
		PaymentTypeCode:   enum.Code,
		CreatedBy:         email,
		CreatedDate:       time.Now(),
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}
	result, err := paymentUsecase.paymentsRepository.InsertPayment(dataPayment)
	return result, err
}

func (paymentUsecase *paymentUsecase) UploadPayment(payment shape.PaymentPost, email string) (bool, error) {
	var srStatusCode string

	status, err := paymentUsecase.approvalStatusRepository.GetApprovalStatus(payment.Status_Payment_Code)
	switch strings.ToLower(payment.Action) {
	case "approved":
		srStatusCode = status.ApprovedStatus.String
	case "rejected":
	default:
		srStatusCode = ""
	}

	dataPayment := model.Payment{
		ID:                payment.ID,
		PaymentMethodCode: payment.Payment_Method_Code,
		StatusPaymentCode: srStatusCode,
		SenderBank: sql.NullString{
			String: payment.Sender_Bank,
		},
		SenderName: sql.NullString{
			String: payment.Sender_Name,
		},
		ImageCode: sql.NullInt64{
			Int64: payment.Image_Code,
		},
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}
	result, err := paymentUsecase.paymentsRepository.UploadPayment(dataPayment)
	return result, err
}

func (paymentUsecase *paymentUsecase) ConfirmPayment(payment shape.PaymentPost, email string) (bool, error) {
	var srStatusCode, paymentType string

	status, err := paymentUsecase.approvalStatusRepository.GetApprovalStatus(payment.Status_Payment_Code)
	switch strings.ToLower(payment.Action) {
	case "approved":
		srStatusCode = status.ApprovedStatus.String
		paymentType = "Completed|Complete|Lunas"
	case "rejected":
		srStatusCode = status.RejectStatus.String
		paymentType = "Failed|Failed|Batal"
	default:
		srStatusCode = ""
	}

	enum, err := paymentUsecase.enumRepository.GetEnum(paymentType)
	dataPayment := model.Payment{
		ID:                payment.ID,
		StatusPaymentCode: srStatusCode,
		PaymentTypeCode:   enum.Code,
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}
	result, err := paymentUsecase.paymentsRepository.ConfirmPayment(dataPayment)
	return result, err
}

func (paymentUsecase *paymentUsecase) CheckAllPaymentExpired() {
	var err error
	var paymentList []model.Payment
	var email = "belajariah20@gmail.com"

	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}
		paymentList, err = paymentUsecase.paymentsRepository.CheckAllPaymentExpired()
		firstloop = false
		if err == nil {
			for _, value := range paymentList {
				dataPayment := shape.PaymentPost{
					ID:                  value.ID,
					Action:              "Rejected",
					Status_Payment_Code: value.StatusPaymentCode,
				}
				result, err := paymentUsecase.ConfirmPayment(dataPayment, email)
				if err != nil {
					utils.PushLogf("ERROR : ", err, result)
				}
			}
		}
	}
}
