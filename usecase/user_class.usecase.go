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

type userClassUsecase struct {
	emailUsecase           EmailUsecase
	enumRepository         repository.EnumRepository
	promotionRepository    repository.PromotionRepository
	userClassRepository    repository.UserClassRepository
	notificationRepository repository.NotificationRepository
}

type UserClassUsecase interface {
	CheckAllUserClassExpired()
	CheckAllUserClass1DaysBeforeExpired()
	CheckAllUserClass2DaysBeforeExpired()
	CheckAllUserClass5DaysBeforeExpired()
	CheckAllUserClass7DaysBeforeExpired()
	GetAllUserClass(query model.Query, userObj model.UserInfo) ([]shape.UserClass, int, error)
	UpdateUserClassProgress(userClass shape.UserClassPost, email string) (bool, error)
}

func InitUserClassUsecase(emailUsecase EmailUsecase, enumRepository repository.EnumRepository, promotionRepository repository.PromotionRepository, userClassRepository repository.UserClassRepository, notificationRepository repository.NotificationRepository) UserClassUsecase {
	return &userClassUsecase{
		emailUsecase,
		enumRepository,
		promotionRepository,
		userClassRepository,
		notificationRepository,
	}
}

func (userClassUsecase *userClassUsecase) GetAllUserClass(query model.Query, userObj model.UserInfo) ([]shape.UserClass, int, error) {
	var filterQuery, filterUser string
	var userClass []model.UserClass
	var userClassResult []shape.UserClass

	filterQuery = utils.GetFilterHandler(query.Filters)
	filterUser = fmt.Sprintf(`AND user_code=%d`, userObj.ID)

	userClass, err := userClassUsecase.userClassRepository.GetAllUserClass(query.Skip, query.Take, filterQuery, filterUser)
	count, errCount := userClassUsecase.userClassRepository.GetAllUserClassCount(filterQuery, filterUser)

	if err == nil && errCount == nil {
		for _, value := range userClass {
			userClassResult = append(userClassResult, shape.UserClass{
				ID:                    value.ID,
				User_Code:             value.UserCode,
				Class_Code:            value.ClassCode,
				Class_Name:            value.ClassName,
				Class_Initial:         value.ClassInitial.String,
				Class_Category:        value.ClassCategory,
				Class_Description:     value.ClassDescription.String,
				Class_Image:           value.ClassImage.String,
				Class_Rating:          value.ClassRating,
				Total_User:            value.TotalUser,
				Type_Code:             value.TypeCode,
				Type:                  value.Type,
				Status_Code:           value.StatusCode,
				Status:                value.Status,
				Package_Code:          value.PackageCode,
				Package_Type:          value.PackageType,
				Is_Expired:            value.IsExpired,
				Start_Date:            utils.HandleNullableDate(value.StartDate),
				Expired_Date:          utils.HandleNullableDate(value.ExpiredDate),
				Time_Duration:         value.TimeDuration,
				Progress:              value.Progress.Float64,
				Progress_Index:        int(value.ProgressIndex.Int64),
				Progress_Cur_Index:    int(value.ProgressCurIndex.Int64),
				Progress_Cur_Subindex: int(value.ProgressCurSubindex.Int64),
				Pre_Test_Scores:       value.PreTestScores.Float64,
				Post_Test_Scores:      value.PostTestScores.Float64,
				Post_Test_Date:        utils.HandleNullableDate(value.PostTestDate.Time),
				Total_Consultation:    int(value.TotalConsultation.Int64),
				Total_Webinar:         int(value.TotalWebinar.Int64),
				Is_Active:             value.IsActive,
				Created_By:            value.CreatedBy,
				Created_Date:          value.CreatedDate,
				Modified_By:           value.ModifiedBy.String,
				Modified_Date:         value.ModifiedDate.Time,
				Deleted_By:            value.DeletedBy.String,
				Deleted_Date:          value.DeletedDate.Time,
			})
		}
	}
	paymentEmpty := make([]shape.UserClass, 0)
	if len(userClassResult) == 0 {
		return paymentEmpty, count, err
	}
	return userClassResult, count, err
}

func (userClassUsecase *userClassUsecase) UpdateUserClassProgress(userClass shape.UserClassPost, email string) (bool, error) {

	dataUserClass := model.UserClass{
		ID:       userClass.ID,
		UserCode: userClass.User_Code,
		Progress: sql.NullFloat64{
			Float64: userClass.Progress,
		},
		ProgressIndex: sql.NullInt64{
			Int64: userClass.Progress_Index,
		},
		ProgressCurIndex: sql.NullInt64{
			Int64: userClass.Progress_Cur_Index,
		},
		ProgressCurSubindex: sql.NullInt64{
			Int64: userClass.Progress_Cur_Subindex,
		},
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}
	result, err := userClassUsecase.userClassRepository.UpdateUserClassProgress(dataUserClass)
	return result, err
}

func (userClassUsecase *userClassUsecase) CheckAllUserClassExpired() {
	var promoCode = "BLJEXPD"
	var types string = "TodayClassExp"
	var table string = "Transact user class"
	var email string = "belajariah20@gmail.com"
	var emailType string = "Class Has Been Expired"

	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}
		userClassList, err := userClassUsecase.userClassRepository.CheckAllUserClassExpired()
		enum, err := userClassUsecase.enumRepository.GetEnumSplit(types)
		firstloop = false
		if err == nil {
			for _, value := range userClassList {
				dataUserClass := model.UserClass{
					ID:       value.ID,
					UserCode: value.UserCode,
					ModifiedBy: sql.NullString{
						String: email,
					},
					ModifiedDate: sql.NullTime{
						Time: time.Now(),
					},
				}
				result, err := userClassUsecase.userClassRepository.UpdateUserClassExpired(dataUserClass)
				if err == nil {
					promotion, err := userClassUsecase.promotionRepository.GetPromotion(promoCode)
					dataEmail := model.EmailBody{
						BodyTemp:      emailType,
						UserCode:      value.UserCode,
						PromoDiscount: fmt.Sprintf(`%d`, int(promotion.Discount.Float64)),
					}
					userClassUsecase.emailUsecase.SendEmail(dataEmail)

					dataNotification := model.Notification{
						TableRef: table,
						UserClassCode: sql.NullInt64{
							Int64: int64(value.ID),
						},
						NotificationType: enum.Code,
						UserCode:         value.UserCode,
						Sequence:         1,
						ExpiredDate: sql.NullTime{
							Time: value.ExpiredDate,
						},
						CreatedBy:   email,
						CreatedDate: time.Now(),
						ModifiedBy: sql.NullString{
							String: email,
						},
						ModifiedDate: sql.NullTime{
							Time: time.Now(),
						},
					}
					result, err := userClassUsecase.notificationRepository.InsertNotification(dataNotification)
					if err != nil {
						utils.PushLogf("ERROR : ", err, result)
					}
				} else {
					utils.PushLogf("ERROR : ", err, result)
				}
			}
		}
	}
}

func (userClassUsecase *userClassUsecase) CheckAllUserClass7DaysBeforeExpired() {
	var minutes float64 = 11520
	var promoCode = "BLJEXPD"
	var table string = "Transact user class"
	var types string = "7DaysBeforeClassExp"
	var email string = "belajariah20@gmail.com"
	var emailType string = "7 Days Before Class Expired"

	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}
		dataInterval := model.TimeInterval{
			Interval1: 604800,
			Interval2: 432000,
		}
		userClassList, err := userClassUsecase.userClassRepository.CheckAllUserClassBeforeExpired(dataInterval)
		enum, err := userClassUsecase.enumRepository.GetEnumSplit(types)
		firstloop = false
		if err == nil {
			for _, value := range userClassList {
				filterNotif := fmt.Sprintf(`AND user_class_code = %d`, value.ID)
				notification, err := userClassUsecase.notificationRepository.GetNotification(filterNotif, types)
				if err == nil && utils.GetDuration(notification.ExpiredDate.Time, value.ExpiredDate) > minutes {
					promotion, err := userClassUsecase.promotionRepository.GetPromotion(promoCode)
					dataEmail := model.EmailBody{
						BodyTemp:      emailType,
						UserCode:      value.UserCode,
						ExpiredDate:   value.ExpiredDate,
						PromoDiscount: fmt.Sprintf(`%d`, int(promotion.Discount.Float64)),
					}
					userClassUsecase.emailUsecase.SendEmail(dataEmail)

					dataNotification := model.Notification{
						TableRef: table,
						UserClassCode: sql.NullInt64{
							Int64: int64(value.ID),
						},
						NotificationType: enum.Code,
						UserCode:         value.UserCode,
						Sequence:         1,
						ExpiredDate: sql.NullTime{
							Time: value.ExpiredDate,
						},
						CreatedBy:   email,
						CreatedDate: time.Now(),
						ModifiedBy: sql.NullString{
							String: email,
						},
						ModifiedDate: sql.NullTime{
							Time: time.Now(),
						},
					}
					result, err := userClassUsecase.notificationRepository.InsertNotification(dataNotification)
					if err != nil {
						utils.PushLogf("ERROR : ", err, result)
					}
				}
			}
		}
	}
}

func (userClassUsecase *userClassUsecase) CheckAllUserClass5DaysBeforeExpired() {
	var minutes float64 = 8640
	var promoCode = "BLJEXPD"
	var table string = "Transact user class"
	var types string = "5DaysBeforeClassExp"
	var email string = "belajariah20@gmail.com"
	var emailType string = "5 Days Before Class Expired"

	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}
		dataInterval := model.TimeInterval{
			Interval1: 432000,
			Interval2: 172800,
		}
		userClassList, err := userClassUsecase.userClassRepository.CheckAllUserClassBeforeExpired(dataInterval)
		enum, err := userClassUsecase.enumRepository.GetEnumSplit(types)
		firstloop = false
		if err == nil {
			for _, value := range userClassList {
				filter := fmt.Sprintf(`AND user_class_code = %d`, value.ID)
				notification, err := userClassUsecase.notificationRepository.GetNotification(filter, types)
				if err == nil && utils.GetDuration(notification.ExpiredDate.Time, value.ExpiredDate) > minutes {
					promotion, err := userClassUsecase.promotionRepository.GetPromotion(promoCode)
					dataEmail := model.EmailBody{
						BodyTemp:      emailType,
						UserCode:      value.UserCode,
						ExpiredDate:   value.ExpiredDate,
						PromoDiscount: fmt.Sprintf(`%d`, int(promotion.Discount.Float64)),
					}
					userClassUsecase.emailUsecase.SendEmail(dataEmail)

					dataNotification := model.Notification{
						TableRef: table,
						UserClassCode: sql.NullInt64{
							Int64: int64(value.ID),
						},
						NotificationType: enum.Code,
						UserCode:         value.UserCode,
						Sequence:         1,
						ExpiredDate: sql.NullTime{
							Time: value.ExpiredDate,
						},
						CreatedBy:   email,
						CreatedDate: time.Now(),
						ModifiedBy: sql.NullString{
							String: email,
						},
						ModifiedDate: sql.NullTime{
							Time: time.Now(),
						},
					}
					result, err := userClassUsecase.notificationRepository.InsertNotification(dataNotification)
					if err != nil {
						utils.PushLogf("ERROR : ", err, result)
					}
				}
			}
		}
	}
}

func (userClassUsecase *userClassUsecase) CheckAllUserClass2DaysBeforeExpired() {
	var minutes float64 = 4320
	var promoCode = "BLJEXPD"
	var table string = "Transact user class"
	var types string = "2DaysBeforeClassExp"
	var email string = "belajariah20@gmail.com"
	var emailType string = "2 Days Before Class Expired"

	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}
		dataInterval := model.TimeInterval{
			Interval1: 172800,
			Interval2: 86400,
		}
		userClassList, err := userClassUsecase.userClassRepository.CheckAllUserClassBeforeExpired(dataInterval)
		enum, err := userClassUsecase.enumRepository.GetEnumSplit(types)
		firstloop = false
		if err == nil {
			for _, value := range userClassList {
				filter := fmt.Sprintf(`AND user_class_code = %d`, value.ID)
				notification, err := userClassUsecase.notificationRepository.GetNotification(filter, types)
				if err == nil && utils.GetDuration(notification.ExpiredDate.Time, value.ExpiredDate) > minutes {
					promotion, err := userClassUsecase.promotionRepository.GetPromotion(promoCode)
					dataEmail := model.EmailBody{
						BodyTemp:      emailType,
						UserCode:      value.UserCode,
						ExpiredDate:   value.ExpiredDate,
						PromoDiscount: fmt.Sprintf(`%d`, int(promotion.Discount.Float64)),
					}
					userClassUsecase.emailUsecase.SendEmail(dataEmail)

					dataNotification := model.Notification{
						TableRef: table,
						UserClassCode: sql.NullInt64{
							Int64: int64(value.ID),
						},
						NotificationType: enum.Code,
						UserCode:         value.UserCode,
						Sequence:         1,
						ExpiredDate: sql.NullTime{
							Time: value.ExpiredDate,
						},
						CreatedBy:   email,
						CreatedDate: time.Now(),
						ModifiedBy: sql.NullString{
							String: email,
						},
						ModifiedDate: sql.NullTime{
							Time: time.Now(),
						},
					}
					result, err := userClassUsecase.notificationRepository.InsertNotification(dataNotification)
					if err != nil {
						utils.PushLogf("ERROR : ", err, result)
					}
				}
			}
		}
	}
}

func (userClassUsecase *userClassUsecase) CheckAllUserClass1DaysBeforeExpired() {
	var minutes float64 = 2880
	var promoCode = "BLJEXPD"
	var table string = "Transact user class"
	var types string = "1DaysBeforeClassExp"
	var email string = "belajariah20@gmail.com"
	var emailType string = "1 Days Before Class Expired"

	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}
		dataInterval := model.TimeInterval{
			Interval1: 86400,
			Interval2: 0,
		}
		userClassList, err := userClassUsecase.userClassRepository.CheckAllUserClassBeforeExpired(dataInterval)
		enum, err := userClassUsecase.enumRepository.GetEnumSplit(types)
		firstloop = false
		if err == nil {
			for _, value := range userClassList {
				filter := fmt.Sprintf(`AND user_class_code = %d`, value.ID)
				notification, err := userClassUsecase.notificationRepository.GetNotification(filter, types)
				if err == nil && utils.GetDuration(notification.ExpiredDate.Time, value.ExpiredDate) > minutes {
					promotion, err := userClassUsecase.promotionRepository.GetPromotion(promoCode)
					dataEmail := model.EmailBody{
						BodyTemp:      emailType,
						UserCode:      value.UserCode,
						ExpiredDate:   value.ExpiredDate,
						PromoDiscount: fmt.Sprintf(`%d`, int(promotion.Discount.Float64)),
					}
					userClassUsecase.emailUsecase.SendEmail(dataEmail)

					dataNotification := model.Notification{
						TableRef: table,
						UserClassCode: sql.NullInt64{
							Int64: int64(value.ID),
						},
						NotificationType: enum.Code,
						UserCode:         value.UserCode,
						Sequence:         1,
						ExpiredDate: sql.NullTime{
							Time: value.ExpiredDate,
						},
						CreatedBy:   email,
						CreatedDate: time.Now(),
						ModifiedBy: sql.NullString{
							String: email,
						},
						ModifiedDate: sql.NullTime{
							Time: time.Now(),
						},
					}
					result, err := userClassUsecase.notificationRepository.InsertNotification(dataNotification)
					if err != nil {
						utils.PushLogf("ERROR : ", err, result)
					}
				}
			}
		}
	}
}
