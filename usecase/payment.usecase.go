package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type paymentUsecase struct {
	sytemConfig                *model.Config
	emailUsecase               EmailUsecase
	enumRepository             repository.EnumRepository
	packageRepository          repository.PackageRepository
	paymentsRepository         repository.PaymentsRepository
	scheduleRepository         repository.ScheduleRepository
	userClassRepository        repository.UserClassRepository
	approvalStatusRepository   repository.ApprovalStatusRepository
	userClassHistoryRepository repository.UserClassHistoryRepository
}

type PaymentUsecase interface {
	CheckAllPaymentExpired()
	CheckAllPayment2HourBeforeExpired()

	GetAllPayment(query model.Query) ([]shape.Payment, int, error)
	GetAllPaymentRejected(query model.Query) ([]shape.Payment, int, error)
	GetAllPaymentByUserID(query model.Query, userObj model.UserHeader) ([]shape.Payment, int, error)

	UploadPayment(payment shape.PaymentPost, email string) (bool, error)
	ConfirmPayment(payment shape.PaymentPost, email string) (bool, error)
	InsertPayment(payment shape.PaymentPost, email string) (shape.Payment, bool, error)

	UploadPaymentQuran(ctx *gin.Context, payment shape.PaymentPost) (bool, error)
	ConfirmPaymentQuran(ctx *gin.Context, payment shape.PaymentPost) (bool, error)
	InsertPaymentQuran(ctx *gin.Context, payment shape.PaymentPost) (shape.Payment, bool, error)
}

func InitPaymentUsecase(sytemConfig *model.Config, emailUsecase EmailUsecase, enumRepository repository.EnumRepository, packageRepository repository.PackageRepository, paymentsRepository repository.PaymentsRepository, scheduleRepository repository.ScheduleRepository, userClassRepository repository.UserClassRepository, approvalStatusRepository repository.ApprovalStatusRepository, userClassHistoryRepository repository.UserClassHistoryRepository) PaymentUsecase {
	return &paymentUsecase{
		sytemConfig,
		emailUsecase,
		enumRepository,
		packageRepository,
		paymentsRepository,
		scheduleRepository,
		userClassRepository,
		approvalStatusRepository,
		userClassHistoryRepository,
	}
}

func (paymentUsecase *paymentUsecase) GetAllPayment(query model.Query) ([]shape.Payment, int, error) {
	var payments []model.Payment
	var paymentResult []shape.Payment
	var filterQuery, filterUser, sorting, search string

	if len(query.Order) > 0 {
		sorting = strings.Replace(query.Order, "|", " ", 1)
		sorting = "ORDER BY " + sorting
	}

	if len(query.Search) > 0 {
		search = `AND LOWER(user_name) LIKE LOWER('%` + query.Search + `%') 
		OR LOWER(invoice_number) LIKE LOWER('%` + query.Search + `%')
		OR LOWER(created_by) LIKE LOWER('%` + query.Search + `%')`
	}

	filterUser = fmt.Sprintf(``)
	filterQuery = utils.GetFilterHandler(query.Filters)

	payments, err := paymentUsecase.paymentsRepository.GetAllPayment(query.Skip, query.Take, sorting, search, filterQuery, filterUser)
	count, errCount := paymentUsecase.paymentsRepository.GetAllPaymentCount(filterQuery, filterUser)

	if err == nil && errCount == nil {
		for _, value := range payments {
			paymentResult = append(paymentResult, shape.Payment{
				ID:                   value.ID,
				Code:                 value.Code,
				User_Code:            value.UserCode,
				User_Name:            value.UserName,
				Class_Code:           value.ClassCode,
				Class_Name:           value.ClassName,
				Class_Initial:        value.ClassInitial,
				Package_Code:         value.PackageCode,
				Package_Type:         value.PackageType,
				Payment_Method_Code:  value.PaymentMethodCode,
				Payment_Method:       value.PaymentMethod,
				Payment_Method_Type:  value.PaymentMethodType,
				Payment_Method_Image: value.PaymentMethodImage.String,
				Account_Name:         value.AccountName.String,
				Account_Number:       value.AccountNumber.String,
				Invoice_Number:       value.InvoiceNumber,
				Status_Payment_Code:  value.StatusPaymentCode,
				Status_Payment:       value.StatusPayment,
				Total_Transfer:       value.TotalTransfer,
				Sender_Bank:          value.SenderBank.String,
				Sender_Name:          value.SenderName.String,
				Image_Proof:          value.ImageProof.String,
				Payment_Type_Code:    value.PaymentTypeCode,
				Payment_Type:         value.PaymentType,
				Schedule_Code_1:      value.ScheduleCode1.String,
				Schedule_Code_2:      value.ScheduleCode2.String,
				Payment_Reference:    value.PaymentReference.String,
				Is_Active:            value.IsActive,
				Created_By:           value.CreatedBy,
				Created_Date:         value.CreatedDate,
				Modified_By:          value.ModifiedBy.String,
				Modified_Date:        value.ModifiedDate.Time,
			})
		}
	}
	paymentEmpty := make([]shape.Payment, 0)
	if len(paymentResult) == 0 {
		return paymentEmpty, count, err
	}
	return paymentResult, count, err
}

func (paymentUsecase *paymentUsecase) GetAllPaymentRejected(query model.Query) ([]shape.Payment, int, error) {
	var payments []model.Payment
	var paymentResult []shape.Payment
	var filterQuery, filterUser, sorting, search string

	paymentEmpty := make([]shape.Payment, 0)

	if len(query.Order) > 0 {
		sorting = strings.Replace(query.Order, "|", " ", 1)
		sorting = "ORDER BY " + sorting
	}
	if len(query.Search) > 0 {
		search = `AND LOWER(user_name) LIKE LOWER('%` + query.Search + `%') 
		OR LOWER(invoice_number) LIKE LOWER('%` + query.Search + `%')
		OR LOWER(created_by) LIKE LOWER('%` + query.Search + `%')`
	}

	filterUser = fmt.Sprintf(`AND status_payment in ('Failed', 'Canceled')`)
	filterQuery = utils.GetFilterHandler(query.Filters)

	payments, err := paymentUsecase.paymentsRepository.GetAllPayment(query.Skip, query.Take, sorting, search, filterQuery, filterUser)
	if err != nil {
		return paymentEmpty, 0, utils.WrapError(err, "paymentUsecase.paymentsRepository.GetAllPayment : ")
	}

	count, errCount := paymentUsecase.paymentsRepository.GetAllPaymentCount(filterQuery, filterUser)
	if errCount != nil {
		return paymentEmpty, 0, utils.WrapError(errCount, "paymentUsecase.paymentsRepository.GetAllPaymentCount : ")
	}

	if err == nil && errCount == nil {
		for _, value := range payments {
			paymentResult = append(paymentResult, shape.Payment{
				ID:                  value.ID,
				Code:                value.Code,
				User_Code:           value.UserCode,
				User_Name:           value.UserName,
				Class_Code:          value.ClassCode,
				Class_Name:          value.ClassName,
				Class_Initial:       value.ClassInitial,
				Package_Code:        value.PackageCode,
				Package_Type:        value.PackageType,
				Payment_Method_Code: value.PaymentMethodCode,
				Payment_Method:      value.PaymentMethod,
				Account_Name:        value.AccountName.String,
				Account_Number:      value.AccountNumber.String,
				Invoice_Number:      value.InvoiceNumber,
				Status_Payment_Code: value.StatusPaymentCode,
				Status_Payment:      value.StatusPayment,
				Total_Transfer:      value.TotalTransfer,
				Sender_Bank:         value.SenderBank.String,
				Sender_Name:         value.SenderName.String,
				Image_Proof:         value.ImageProof.String,
				Payment_Type_Code:   value.PaymentTypeCode,
				Payment_Type:        value.PaymentType,
				Schedule_Code_1:     value.ScheduleCode1.String,
				Schedule_Code_2:     value.ScheduleCode2.String,
				Payment_Reference:   value.PaymentReference.String,
				Is_Active:           value.IsActive,
				Created_By:          value.CreatedBy,
				Created_Date:        value.CreatedDate,
				Modified_By:         value.ModifiedBy.String,
				Modified_Date:       value.ModifiedDate.Time,
			})
		}
	}

	if len(paymentResult) == 0 {
		return paymentEmpty, count, err
	}

	return paymentResult, count, err
}

func (paymentUsecase *paymentUsecase) GetAllPaymentByUserID(query model.Query, userObj model.UserHeader) ([]shape.Payment, int, error) {
	var payments []model.Payment
	var paymentResult []shape.Payment
	var filterQuery, filterUser, sorting, search string

	paymentEmpty := make([]shape.Payment, 0)

	if len(query.Order) > 0 {
		sorting = strings.Replace(query.Order, "|", " ", 1)
		sorting = "ORDER BY " + sorting
	}
	if len(query.Search) > 0 {
		search = `AND (LOWER(user_name) LIKE LOWER('%` + query.Search + `%') 
		OR LOWER(invoice_number) LIKE LOWER('%` + query.Search + `%'))
		OR LOWER(email) LIKE LOWER('%` + query.Search + `%'))`
	}

	filterQuery = utils.GetFilterHandler(query.Filters)
	filterUser = fmt.Sprintf(`AND user_code='%s'`, userObj.Code)

	payments, err := paymentUsecase.paymentsRepository.GetAllPayment(query.Skip, query.Take, sorting, search, filterQuery, filterUser)
	if err != nil {
		return paymentEmpty, 0, utils.WrapError(err, "paymentUsecase.paymentsRepository.GetAllPayment : ")
	}

	count, errCount := paymentUsecase.paymentsRepository.GetAllPaymentCount(filterQuery, filterUser)
	if errCount != nil {
		return paymentEmpty, 0, utils.WrapError(errCount, "paymentUsecase.paymentsRepository.GetAllPaymentCount : ")
	}

	if err == nil && errCount == nil {
		for _, value := range payments {
			paymentResult = append(paymentResult, shape.Payment{
				ID:                   value.ID,
				Code:                 value.Code,
				User_Code:            value.UserCode,
				User_Name:            value.UserName,
				Class_Code:           value.ClassCode,
				Class_Name:           value.ClassName,
				Class_Initial:        value.ClassInitial,
				Package_Code:         value.PackageCode,
				Package_Type:         value.PackageType,
				Payment_Method_Code:  value.PaymentMethodCode,
				Payment_Method:       value.PaymentMethod,
				Payment_Method_Type:  value.PaymentMethodType,
				Payment_Method_Image: value.PaymentMethodImage.String,
				Payment_Type_Code:    value.PaymentTypeCode,
				Payment_Type:         value.PaymentType,
				Account_Name:         value.AccountName.String,
				Account_Number:       value.AccountNumber.String,
				Invoice_Number:       value.InvoiceNumber,
				Status_Payment_Code:  value.StatusPaymentCode,
				Status_Payment:       value.StatusPayment,
				Total_Transfer:       value.TotalTransfer,
				Sender_Bank:          value.SenderBank.String,
				Sender_Name:          value.SenderName.String,
				Image_Proof:          value.ImageProof.String,
				Schedule_Code_1:      value.ScheduleCode1.String,
				Schedule_Code_2:      value.ScheduleCode2.String,
				Payment_Reference:    value.PaymentReference.String,
				Is_Active:            value.IsActive,
				Created_By:           value.CreatedBy,
				Created_Date:         value.CreatedDate,
				Modified_By:          value.ModifiedBy.String,
				Modified_Date:        value.ModifiedDate.Time,
				Expired_Date:         utils.HandleAddDate(value.ModifiedDate.Time, value.StatusPayment),
			})
		}
	}

	if len(paymentResult) == 0 {
		return paymentEmpty, count, err
	}

	return paymentResult, count, err
}

func (paymentUsecase *paymentUsecase) InsertPayment(payment shape.PaymentPost, email string) (shape.Payment, bool, error) {
	var dataPayments shape.Payment
	var emailType string = "Waiting for Payment"
	var paymentType string = "WaitingForPayment|Waiting for Payment|Menunggu"

	enum, err := paymentUsecase.enumRepository.GetEnum(paymentType)
	if err != nil {
		return dataPayments, false, utils.WrapError(err, "paymentUsecase.enumRepository.GetEnum : ")
	}

	dataPayment := model.Payment{
		UserCode:  payment.User_Code,
		ClassCode: payment.Class_Code,
		PromoCode: sql.NullString{
			String: payment.Promo_Code,
		},
		PackageCode:       payment.Package_Code,
		PaymentMethodCode: payment.Payment_Method_Code,
		InvoiceNumber:     utils.GenerateInvoiceNumber(payment, "class general"),
		StatusPaymentCode: payment.Status_Payment_Code,
		TotalTransfer:     payment.Total_Transfer,
		PaymentTypeCode:   enum.Code,
		ScheduleCode1:     sql.NullString{String: payment.Schedule_Code_1},
		ScheduleCode2:     sql.NullString{String: payment.Schedule_Code_2},
		CreatedBy:         email,
		CreatedDate:       time.Now(),
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}

	data, result, err := paymentUsecase.paymentsRepository.InsertPayment(dataPayment)
	if err != nil {
		return dataPayments, false, utils.WrapError(err, "paymentUsecase.paymentsRepository.InsertPayment : ")
	}

	if data != (model.Payment{}) {
		filter := fmt.Sprintf(`WHERE code = '%s'`, data.Code)

		payments, err := paymentUsecase.paymentsRepository.GetPayment(filter)
		if err != nil {
			return dataPayments, false, utils.WrapError(err, "paymentUsecase.paymentsRepository.GetPayment : ")
		}

		dataPayments = shape.Payment{
			ID:                   payments.ID,
			Code:                 data.Code,
			Total_Transfer:       payments.TotalTransfer,
			Payment_Method:       payments.PaymentMethod,
			Payment_Method_Code:  payments.PaymentMethodCode,
			Payment_Method_Type:  payments.PaymentMethodType,
			Payment_Method_Image: payments.PaymentMethodImage.String,
			Status_Payment_Code:  payments.StatusPaymentCode,
			Account_Name:         payments.AccountName.String,
			Account_Number:       payments.AccountNumber.String,
			Expired_Date:         utils.HandleAddDate(payments.ModifiedDate.Time, payments.StatusPayment),
		}

		dataEmail := model.EmailBody{
			BodyTemp:          emailType,
			UserCode:          payments.UserCode,
			InvoiceNumber:     payments.InvoiceNumber,
			PaymentMethod:     payments.PaymentMethod,
			AccountName:       payments.AccountName.String,
			AccountNumber:     payments.AccountNumber.String,
			ClassName:         payments.ClassInitial,
			ClassPrice:        int(payments.PackageDiscount.Int64),
			ClassImage:        payments.ClassImage.String,
			PromoDiscount:     fmt.Sprintf("%d", int(payments.PromoDiscount.Float64)),
			TotalConsultation: int(payments.TotalConsultation.Int64),
			TotalWebinar:      int(payments.TotalWebinar.Int64),
			TotalTransfer:     payments.TotalTransfer,
		}

		paymentUsecase.emailUsecase.SendEmail(dataEmail)
	}

	return dataPayments, result, err
}

func (paymentUsecase *paymentUsecase) InsertPaymentQuran(ctx *gin.Context, payment shape.PaymentPost) (shape.Payment, bool, error) {
	var dataPayments shape.Payment
	var emailType string = "Waiting for Payment"
	var paymentType string = "WaitingForPayment|Waiting for Payment|Menunggu"

	email := ctx.Request.Header.Get("email")

	enum, err := paymentUsecase.enumRepository.GetEnum(paymentType)
	if err != nil {
		return dataPayments, false, utils.WrapError(err, "paymentUsecase.enumRepository.GetEnum : PaymentClassQuran ")
	}

	dataPayment := model.Payment{
		UserCode:  payment.User_Code,
		ClassCode: payment.Class_Code,
		PromoCode: sql.NullString{
			String: payment.Promo_Code,
		},
		PackageCode:       payment.Package_Code,
		PaymentMethodCode: payment.Payment_Method_Code,
		InvoiceNumber:     utils.GenerateInvoiceNumber(payment, "class quran"),
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

	data, result, err := paymentUsecase.paymentsRepository.InsertPaymentQuran(dataPayment)
	if err != nil {
		return dataPayments, false, utils.WrapError(err, "paymentUsecase.paymentsRepository.InsertPaymentQuran : ")
	}

	if data != (model.Payment{}) {
		filter := fmt.Sprintf(`WHERE code = '%s'`, data.Code)

		payments, err := paymentUsecase.paymentsRepository.GetPayment(filter)
		if err != nil {
			return dataPayments, false, utils.WrapError(err, "paymentUsecase.paymentsRepository.GetPaymentQuran : ")
		}

		dataPayments = shape.Payment{
			ID:                   payments.ID,
			Code:                 data.Code,
			Total_Transfer:       payments.TotalTransfer,
			Payment_Method:       payments.PaymentMethod,
			Payment_Method_Code:  payments.PaymentMethodCode,
			Payment_Method_Type:  payments.PaymentMethodType,
			Payment_Method_Image: payments.PaymentMethodImage.String,
			Status_Payment_Code:  payments.StatusPaymentCode,
			Account_Name:         payments.AccountName.String,
			Account_Number:       payments.AccountNumber.String,
			Expired_Date:         utils.HandleAddDate(payments.ModifiedDate.Time, payments.StatusPayment),
		}

		dataEmail := model.EmailBody{
			BodyTemp:          emailType,
			UserCode:          payments.UserCode,
			InvoiceNumber:     payments.InvoiceNumber,
			PaymentMethod:     payments.PaymentMethod,
			AccountName:       payments.AccountName.String,
			AccountNumber:     payments.AccountNumber.String,
			ClassName:         payments.ClassInitial,
			ClassPrice:        int(payments.PackageDiscount.Int64),
			ClassImage:        payments.ClassImage.String,
			PromoDiscount:     fmt.Sprintf("%d", int(payments.PromoDiscount.Float64)),
			TotalConsultation: int(payments.TotalConsultation.Int64),
			TotalWebinar:      int(payments.TotalWebinar.Int64),
			TotalTransfer:     payments.TotalTransfer,
		}

		paymentUsecase.emailUsecase.SendEmail(dataEmail)
	}

	return dataPayments, result, err
}

func (paymentUsecase *paymentUsecase) UploadPayment(payment shape.PaymentPost, email string) (bool, error) {
	var statusCode, emailType string

	status, err := paymentUsecase.approvalStatusRepository.GetApprovalStatus(payment.Status_Payment_Code)
	if err != nil {
		return false, utils.WrapError(err, "paymentUsecase.approvalStatusRepository.GetApprovalStatus : ")
	}

	switch strings.ToLower(payment.Action) {
	case "approved":
		statusCode = status.ApprovedStatus.String
		emailType = "Payment Upload"
	case "rejected":
		statusCode = status.RejectStatus.String
	default:
		statusCode = ""
		return false, utils.WrapError(errors.New("No action needed"), "paymentUsecase")
	}

	dataPayment := model.Payment{
		ID:                payment.ID,
		PaymentMethodCode: payment.Payment_Method_Code,
		StatusPaymentCode: statusCode,
		SenderBank: sql.NullString{
			String: payment.Sender_Bank,
		},
		SenderName: sql.NullString{
			String: payment.Sender_Name,
		},
		ImageProof: sql.NullString{
			String: payment.Image_Proof,
		},
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}

	result, err := paymentUsecase.paymentsRepository.UploadPayment(dataPayment)
	if err != nil {
		return false, utils.WrapError(err, "paymentUsecase.paymentsRepository.UploadPayment : ")
	}

	filters := fmt.Sprintf(`WHERE code = '%s'`, payment.Code)

	payments, err := paymentUsecase.paymentsRepository.GetPayment(filters)
	if err != nil {
		return false, utils.WrapError(err, "paymentUsecase.paymentsRepository.GetPayment : ")
	}

	dataEmail := model.EmailBody{
		BodyTemp:          emailType,
		UserCode:          payment.User_Code,
		InvoiceNumber:     payments.InvoiceNumber,
		PaymentMethod:     payments.PaymentMethod,
		AccountName:       payments.AccountName.String,
		AccountNumber:     payments.AccountNumber.String,
		ClassName:         payments.ClassInitial,
		ClassPrice:        int(payments.PackageDiscount.Int64),
		ClassImage:        payments.ClassImage.String,
		PromoDiscount:     fmt.Sprintf("%d", int(payments.PromoDiscount.Float64)),
		TotalConsultation: int(payments.TotalConsultation.Int64),
		TotalWebinar:      int(payments.TotalWebinar.Int64),
		TotalTransfer:     payments.TotalTransfer,
	}
	paymentUsecase.emailUsecase.SendEmail(dataEmail)

	return result, err
}

func (paymentUsecase *paymentUsecase) UploadPaymentQuran(ctx *gin.Context, payment shape.PaymentPost) (bool, error) {
	var statusCode, emailType string
	email := ctx.Request.Header.Get("email")

	status, err := paymentUsecase.approvalStatusRepository.GetApprovalStatus(payment.Status_Payment_Code)
	if err != nil {
		return false, utils.WrapError(err, "paymentUsecase.approvalStatusRepository.GetApprovalStatus : PaymentClassQuran ")
	}

	switch strings.ToLower(payment.Action) {
	case "approved":
		statusCode = status.ApprovedStatus.String
		emailType = "Payment Upload"
	case "rejected":
		statusCode = status.RejectStatus.String
	default:
		statusCode = ""
		return false, utils.WrapError(errors.New("No action needed"), "paymentUsecase")
	}

	dataPayment := model.Payment{
		ID:                payment.ID,
		PaymentMethodCode: payment.Payment_Method_Code,
		StatusPaymentCode: statusCode,
		SenderBank: sql.NullString{
			String: payment.Sender_Bank,
		},
		SenderName: sql.NullString{
			String: payment.Sender_Name,
		},
		ImageProof: sql.NullString{
			String: payment.Image_Proof,
		},
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}

	result, err := paymentUsecase.paymentsRepository.UploadPaymentQuran(dataPayment)
	if err != nil {
		return false, utils.WrapError(err, "paymentUsecase.paymentsRepository.UploadPaymentQuran : ")
	}

	filters := fmt.Sprintf(`WHERE code = '%s'`, payment.Code)

	payments, err := paymentUsecase.paymentsRepository.GetPayment(filters)
	if err != nil {
		return false, utils.WrapError(err, "paymentUsecase.paymentsRepository.GetPaymentQuran : ")
	}

	dataEmail := model.EmailBody{
		BodyTemp:          emailType,
		UserCode:          payment.User_Code,
		InvoiceNumber:     payments.InvoiceNumber,
		PaymentMethod:     payments.PaymentMethod,
		AccountName:       payments.AccountName.String,
		AccountNumber:     payments.AccountNumber.String,
		ClassName:         payments.ClassInitial,
		ClassPrice:        int(payments.PackageDiscount.Int64),
		ClassImage:        payments.ClassImage.String,
		PromoDiscount:     fmt.Sprintf("%d", int(payments.PromoDiscount.Float64)),
		TotalConsultation: int(payments.TotalConsultation.Int64),
		TotalWebinar:      int(payments.TotalWebinar.Int64),
		TotalTransfer:     payments.TotalTransfer,
	}
	paymentUsecase.emailUsecase.SendEmail(dataEmail)

	return result, err
}

func (paymentUsecase *paymentUsecase) ConfirmPayment(payment shape.PaymentPost, email string) (bool, error) {
	var err error
	var result bool
	var enum model.Enum
	var class model.UserClass
	var packages model.Package
	var schedules *[]model.Schedule
	var history model.UserClassHistory
	var userClassResult model.UserClass
	var statusCode, paymentType, emailType string

	var filter = fmt.Sprintf(`AND user_code='%s' AND class_code='%s'`,
		payment.User_Code,
		payment.Class_Code)

	var filterSchedule = fmt.Sprintf(`AND code IN ('%s', '%s') ORDER BY sequence ASC`,
		payment.Schedule_Code_1,
		payment.Schedule_Code_1)

	status, err := paymentUsecase.approvalStatusRepository.GetApprovalStatus(payment.Status_Payment_Code)
	if err != nil {
		return false, utils.WrapError(err, "paymentUsecase.approvalStatusRepository.GetApprovalStatus : ")
	}

	switch strings.ToLower(payment.Action) {
	case "approved":
		statusCode = status.ApprovedStatus.String
		paymentType = "Completed|Complete|Lunas"
		emailType = "Payment Success"

		packages, err = paymentUsecase.packageRepository.GetPackage(payment.Package_Code)
		if err != nil {
			return false, utils.WrapError(err, "paymentUsecase.packageRepository.GetPackage : ")
		}

		dataUserClass := model.UserClass{
			StatusCode:   statusCode,
			UserCode:     payment.User_Code,
			ClassCode:    payment.Class_Code,
			PackageCode:  payment.Package_Code,
			StartDate:    sql.NullTime{Time: time.Now()},
			PromoCode:    sql.NullString{String: payment.Promo_Code},
			CreatedBy:    email,
			CreatedDate:  time.Now(),
			ModifiedBy:   sql.NullString{String: email},
			ModifiedDate: sql.NullTime{Time: time.Now()},
		}

		class, err = paymentUsecase.userClassRepository.GetUserClass(filter)
		if err != nil {
			return false, utils.WrapError(err, "paymentUsecase.userClassRepository.GetUserClass : ")
		}

		if !payment.Is_Direct {
			if class == (model.UserClass{}) {
				userClassResult, result, err = paymentUsecase.userClassRepository.InsertUserClass(dataUserClass)
				if err != nil {
					return false, utils.WrapError(err, "paymentUsecase.userClassRepository.InsertUserClass : ")
				}

				dataHistory := model.UserClassHistory{
					UserClassCode:     userClassResult.Code,
					PackageCode:       dataUserClass.PackageCode,
					PromoCode:         dataUserClass.PromoCode.String,
					PaymentMethodCode: payment.Payment_Method_Code,
					Price:             payment.Total_Transfer,
					StartDate:         dataUserClass.StartDate,
					ExpiredDate:       dataUserClass.ExpiredDate,
					CreatedBy:         dataUserClass.CreatedBy,
					CreatedDate:       dataUserClass.CreatedDate,
					ModifiedBy:        dataUserClass.ModifiedBy,
					ModifiedDate:      dataUserClass.ModifiedDate,
				}

				history, result, err = paymentUsecase.userClassHistoryRepository.InsertUserClassHistory(dataHistory)
				if err != nil {
					return false, utils.WrapError(err, "paymentUsecase.userClassHistoryRepository.InsertUserClassHistory : ")
				}
			}
		} else {
			if class == (model.UserClass{}) {
				userClassResult, result, err = paymentUsecase.userClassRepository.InsertUserClass(dataUserClass)
				if err != nil {
					return false, utils.WrapError(err, "paymentUsecase.userClassRepository.InsertUserClass : ")
				}
			} else if class != (model.UserClass{}) {
				userClassResult, result, err = paymentUsecase.userClassRepository.InsertUserClass(dataUserClass)
				if err != nil {
					return false, utils.WrapError(err, "paymentUsecase.userClassRepository.InsertUserClass : ")
				}
			}

			schedules, err = paymentUsecase.scheduleRepository.GetAllMasterSchedule(filterSchedule)
			if err != nil {
				return false, utils.WrapError(err, "paymentUsecase.scheduleRepository.GetAllSchedule : ")
			}

			if len(*schedules) != 0 {
				for i := 0; i < int(packages.DurationFrequence.Int64)/len(*schedules); i++ {

					for index, data := range *schedules {
						data.User_Class_Code = history.Code
						data.Class_Code = payment.Class_Code
						data.User_Code = payment.User_Code
						data.Sequence = i + (index + i + 1)
						data.Created_By = email
						data.Modified_By = email
						data.Created_Date = dataUserClass.CreatedDate
						data.Modified_Date.NullTime = dataUserClass.ModifiedDate
						data.Description.NullString = sql.NullString{String: fmt.Sprintf(`Pertemuan %d`, i+(index+i+1))}

						result, err = paymentUsecase.scheduleRepository.InsertSchedule(data)
						if err != nil {
							return false, utils.WrapError(err, "paymentUsecase.scheduleRepository.GetAllSchedule : ")
						}
					}
				}

				dataHistory := model.UserClassHistory{
					UserClassCode:     userClassResult.Code,
					PackageCode:       dataUserClass.PackageCode,
					PromoCode:         dataUserClass.PromoCode.String,
					PaymentMethodCode: payment.Payment_Method_Code,
					Price:             payment.Total_Transfer,
					StartDate:         dataUserClass.StartDate,
					ExpiredDate:       dataUserClass.ExpiredDate,
					CreatedBy:         dataUserClass.CreatedBy,
					CreatedDate:       dataUserClass.CreatedDate,
					ModifiedBy:        dataUserClass.ModifiedBy,
					ModifiedDate:      dataUserClass.ModifiedDate,
				}

				history, result, err = paymentUsecase.userClassHistoryRepository.InsertUserClassHistory(dataHistory)
				if err != nil {
					return false, utils.WrapError(err, "paymentUsecase.userClassHistoryRepository.InsertUserClassHistory : ")
				}
			}
		}

	case "rejected":
		statusCode = status.RejectStatus.String
		paymentType = "Failed|Failed|Batal"
		emailType = "Payment Canceled"
	case "revised":
		statusCode = status.ReviseStatus.String
		paymentType = "WaitingForPayment|Waiting for Payment|Menunggu"
		emailType = "Payment Revised"

	default:
		statusCode = ""
		return false, utils.WrapError(errors.New("No action needed"), "paymentUsecase")
	}

	enum, err = paymentUsecase.enumRepository.GetEnum(paymentType)
	if err != nil {
		return false, utils.WrapError(err, "paymentUsecase.enumRepository.GetEnum : ")
	}

	dataPayment := model.Payment{
		ID:                payment.ID,
		StatusPaymentCode: statusCode,
		PaymentTypeCode:   enum.Code,
		Remarks: sql.NullString{
			String: payment.Remarks,
		},
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}

	result, err = paymentUsecase.paymentsRepository.ConfirmPayment(dataPayment)
	if err != nil {
		return false, utils.WrapError(err, "paymentUsecase.paymentsRepository.ConfirmPayment : ")
	}

	if status.CurrentStatusValue == "Has been Payment" && payment.Action != "Rejected" {
		filters := fmt.Sprintf(`WHERE code = '%s'`, payment.Code)
		payments, err := paymentUsecase.paymentsRepository.GetPayment(filters)
		if err != nil {
			return false, utils.WrapError(err, "paymentUsecase.paymentsRepository.GetPayment : ")
		}

		dataEmail := model.EmailBody{
			BodyTemp:          emailType,
			UserCode:          payments.UserCode,
			InvoiceNumber:     payments.InvoiceNumber,
			PaymentMethod:     payments.PaymentMethod,
			AccountName:       payments.AccountName.String,
			AccountNumber:     payments.AccountNumber.String,
			ClassName:         payments.ClassInitial,
			ClassPrice:        int(payments.PackageDiscount.Int64),
			ClassImage:        payments.ClassImage.String,
			PromoDiscount:     fmt.Sprintf("%d", int(payments.PromoDiscount.Float64)),
			TotalConsultation: int(payments.TotalConsultation.Int64),
			TotalWebinar:      int(payments.TotalWebinar.Int64),
			TotalTransfer:     payments.TotalTransfer,
		}

		paymentUsecase.emailUsecase.SendEmail(dataEmail)
	}

	return result, err
}

func (paymentUsecase *paymentUsecase) ConfirmPaymentQuran(ctx *gin.Context, payment shape.PaymentPost) (bool, error) {
	var err error
	var result bool
	var email string
	var enum model.Enum
	var class []shape.UserClass
	var statusCode, paymentType, emailType string

	var filter = fmt.Sprintf(`WHERE is_deleted=false AND user_code='%s' AND class_code='%s'`,
		payment.User_Code,
		payment.Class_Code)

	status, err := paymentUsecase.approvalStatusRepository.GetApprovalStatus(payment.Status_Payment_Code)
	if err != nil {
		return false, utils.WrapError(err, "paymentUsecase.approvalStatusRepository.GetApprovalStatus : PaymentClassQuran ")
	}

	switch strings.ToLower(payment.Action) {
	case "approved":
		email = ctx.Request.Header.Get("email")

		statusCode = status.ApprovedStatus.String
		paymentType = "Completed|Complete|Lunas"
		emailType = "Payment Success"

		dataUserClass := model.UserClass{
			StatusCode:   statusCode,
			UserCode:     payment.User_Code,
			ClassCode:    payment.Class_Code,
			PackageCode:  payment.Package_Code,
			StartDate:    sql.NullTime{Time: time.Now()},
			PromoCode:    sql.NullString{String: payment.Promo_Code},
			CreatedBy:    email,
			CreatedDate:  time.Now(),
			ModifiedBy:   sql.NullString{String: email},
			ModifiedDate: sql.NullTime{Time: time.Now()},
		}

		class, err = paymentUsecase.userClassRepository.GetAllUserClassQuran(filter)
		if err != nil {
			return false, utils.WrapError(err, "paymentUsecase.userClassRepository.GetUserClassQuran : ")
		}

		if len(class) == 0 {
			_, result, err = paymentUsecase.userClassRepository.InsertUserClassQuran(dataUserClass)
			if err != nil {
				return false, utils.WrapError(err, "paymentUsecase.userClassRepository.InsertUserClassQuran : PaymentClassQuran ")
			}
		}
	case "rejected":
		email = ctx.Request.Header.Get("email")
		statusCode = status.RejectStatus.String
		paymentType = "Failed|Failed|Batal"
		emailType = "Payment Canceled"
	case "revised":
		email = ctx.Request.Header.Get("email")
		statusCode = status.ReviseStatus.String
		paymentType = "WaitingForPayment|Waiting for Payment|Menunggu"
		emailType = "Payment Revised"
	case "system rejected":
		email = payment.Modified_By
		statusCode = status.RejectStatus.String
		paymentType = "Failed|Failed|Batal"
		emailType = "Payment Failed"

	default:
		statusCode = ""
		return false, utils.WrapError(errors.New("No action needed"), "paymentUsecase")
	}

	enum, err = paymentUsecase.enumRepository.GetEnum(paymentType)
	if err != nil {
		return false, utils.WrapError(err, "paymentUsecase.enumRepository.GetEnum : PaymentClassQuran ")
	}

	dataPayment := model.Payment{
		ID:                payment.ID,
		Code:              payment.Code,
		StatusPaymentCode: statusCode,
		PaymentTypeCode:   enum.Code,
		Remarks: sql.NullString{
			String: payment.Remarks,
		},
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}

	result, err = paymentUsecase.paymentsRepository.ConfirmPaymentQuran(dataPayment)
	if err != nil {
		return false, utils.WrapError(err, "paymentUsecase.paymentsRepository.ConfirmPayment : PaymentClassQuran ")
	}

	if (status.CurrentStatusValue == "Has been Payment" && payment.Action != "Rejected") || payment.Action == "System Rejected" {
		filters := fmt.Sprintf(`WHERE code = '%s'`, payment.Code)
		payments, err := paymentUsecase.paymentsRepository.GetPayment(filters)
		if err != nil {
			return false, utils.WrapError(err, "paymentUsecase.paymentsRepository.GetPayment : PaymentClassQuran ")
		}

		dataEmail := model.EmailBody{
			BodyTemp:      emailType,
			UserCode:      payments.UserCode,
			InvoiceNumber: payments.InvoiceNumber,
			PaymentMethod: payments.PaymentMethod,
			AccountName:   payments.AccountName.String,
			AccountNumber: payments.AccountNumber.String,
			ClassName:     payments.ClassInitial,
			ClassImage:    payments.ClassImage.String,
			ClassPrice:    int(payments.PackageDiscount.Int64),
			TotalTransfer: payments.TotalTransfer,
			PromoDiscount: fmt.Sprintf("%d", int(payments.PromoDiscount.Float64)),
		}
		paymentUsecase.emailUsecase.SendEmail(dataEmail)
	}

	return result, err
}

func (paymentUsecase *paymentUsecase) CheckAllPaymentExpired() {
	var err error
	var ctx *gin.Context
	var paymentList []model.Payment

	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}

		paymentList, err = paymentUsecase.paymentsRepository.CheckAllPaymentExpired()
		if err != nil {
			utils.PushLogf("paymentUsecase.paymentsRepository.CheckAllPaymentExpired : ", err.Error())
		}

		firstloop = false
		if err == nil {
			for _, value := range paymentList {
				dataPayment := shape.PaymentPost{
					ID:                  value.ID,
					Code:                value.Code,
					Action:              "System Rejected",
					Status_Payment_Code: value.StatusPaymentCode,
					Modified_By:         paymentUsecase.sytemConfig.System.EmailSystem,
				}

				_, err := paymentUsecase.ConfirmPaymentQuran(ctx, dataPayment)
				if err != nil {
					utils.PushLogf("paymentUsecase.CheckAllPaymentExpired.ConfirmPayment : ", err.Error())
				}

			}
		}
	}
}

func (paymentUsecase *paymentUsecase) CheckAllPayment2HourBeforeExpired() {
	var err error
	var paymentList []model.Payment
	var emailType string = "2 Hour Before Payment Expired"

	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}

		paymentList, err = paymentUsecase.paymentsRepository.CheckAllPaymentBeforeExpired()
		if err != nil {
			utils.PushLogf("paymentUsecase.paymentsRepository.CheckAllPaymentBeforeExpired : ", err.Error())
		}

		firstloop = false
		if err == nil {
			for _, value := range paymentList {
				filter := fmt.Sprintf(`WHERE code = '%s'`, value.Code)
				payments, _ := paymentUsecase.paymentsRepository.GetPayment(filter)

				dataEmail := model.EmailBody{
					BodyTemp:      emailType,
					UserCode:      payments.UserCode,
					InvoiceNumber: payments.InvoiceNumber,
					PaymentMethod: payments.PaymentMethod,
					AccountName:   payments.AccountName.String,
					AccountNumber: payments.AccountNumber.String,
					ClassName:     payments.ClassInitial,
					ClassPrice:    int(payments.PackageDiscount.Int64),
				}
				paymentUsecase.emailUsecase.SendEmail(dataEmail)

			}
		}
	}
}
