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

	"github.com/gin-gonic/gin"
)

type userClassUsecase struct {
	sytemConfig            *model.Config
	emailUsecase           EmailUsecase
	userRepository         repository.UserRepository
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

	GetUserClass(code string, userObj model.UserHeader) (shape.UserClass, error)
	GetAllUserClass(query model.Query, userObj model.UserHeader) ([]shape.UserClass, int, error)

	GetAllUserClassQuran(ctx *gin.Context, r model.UserClassRequest) ([]shape.UserClass, int, error)

	UpdateUserClassProgress(userClass shape.UserClassPost, email string) (bool, error)
}

func InitUserClassUsecase(sytemConfig *model.Config, emailUsecase EmailUsecase, userRepository repository.UserRepository, enumRepository repository.EnumRepository, promotionRepository repository.PromotionRepository, userClassRepository repository.UserClassRepository, notificationRepository repository.NotificationRepository) UserClassUsecase {
	return &userClassUsecase{
		sytemConfig,
		emailUsecase,
		userRepository,
		enumRepository,
		promotionRepository,
		userClassRepository,
		notificationRepository,
	}
}

func (userClassUsecase *userClassUsecase) GetUserClass(code string, userObj model.UserHeader) (shape.UserClass, error) {

	filter := fmt.Sprintf(`AND user_code='%s' AND class_code='%s'`, userObj.Code, code)
	value, err := userClassUsecase.userClassRepository.GetUserClass(filter)
	if err != nil {
		return shape.UserClass{}, utils.WrapError(err, "userClassUsecase.mentorRepuserClassRepositoryository.GetUserClass : ")
	}

	userClassResult := shape.UserClass{
		ID:                 value.ID,
		User_Code:          value.UserCode,
		Class_Code:         value.ClassCode,
		Total_User:         value.TotalUser,
		Type_Code:          value.TypeCode,
		Type:               value.Type,
		Status_Code:        value.StatusCode,
		Status:             value.Status,
		Package_Code:       value.PackageCode,
		Package_Type:       value.PackageType,
		Is_Expired:         value.IsExpired,
		Start_Date:         utils.HandleNullableDate(value.StartDate.Time),
		Expired_Date:       utils.HandleNullableDate(value.ExpiredDate.Time),
		Time_Duration:      value.TimeDuration,
		Progress:           value.Progress.Float64,
		Progress_Count:     int(value.ProgressCount.Int64),
		Progress_Index:     int(value.ProgressIndex.Int64),
		Progress_Subindex:  int(value.ProgressSubindex.Int64),
		Pre_Test_Scores:    value.PreTestScores.Float64,
		Post_Test_Scores:   value.PostTestScores.Float64,
		Post_Test_Date:     utils.HandleNullableDate(value.PostTestDate.Time),
		Pre_Test_Total:     int(value.PreTestTotal.Int64),
		Post_Test_Total:    int(value.PostTestTotal.Int64),
		Total_Consultation: int(value.TotalConsultation.Int64),
		Total_Webinar:      int(value.TotalWebinar.Int64),
	}
	return userClassResult, err
}

func (u *userClassUsecase) GetAllUserClass(query model.Query, userObj model.UserHeader) ([]shape.UserClass, int, error) {
	var userClass []model.UserClass
	var userClassResult []shape.UserClass
	var filterQuery, filterUser, sorting string
	classEmpty := make([]shape.UserClass, 0)

	if len(query.Order) > 0 {
		sorting = strings.Replace(query.Order, "|", " ", 1)
		sorting = "ORDER BY " + sorting
	}

	filterQuery = utils.GetFilterHandler(query.Filters)
	filterUser = fmt.Sprintf(`AND user_code='%s'`, userObj.Code)

	userClass, err := u.userClassRepository.GetAllUserClass(query.Skip, query.Take, sorting, filterQuery, filterUser)
	if err != nil {
		return classEmpty, 0, utils.WrapError(err, "userClassUsecase.mentorRepuserClassRepositoryository.GetAllUserClass : ")
	}

	count, errCount := u.userClassRepository.GetAllUserClassCount(filterQuery, filterUser)
	if err != nil {
		return classEmpty, 0, utils.WrapError(err, "userClassUsecase.mentorRepuserClassRepositoryository.GetAllUserClassCount : ")
	}

	if err == nil && errCount == nil {
		for _, value := range userClass {
			userClassResult = append(userClassResult, shape.UserClass{
				ID:                 value.ID,
				User_Code:          value.UserCode,
				Class_Code:         value.ClassCode,
				Class_Name:         value.ClassName,
				Class_Initial:      value.ClassInitial.String,
				Class_Category:     value.ClassCategory,
				Class_Description:  value.ClassDescription.String,
				Class_Image:        value.ClassImage.String,
				Class_Rating:       value.ClassRating,
				Total_User:         value.TotalUser,
				Type_Code:          value.TypeCode,
				Type:               value.Type,
				Status_Code:        value.StatusCode,
				Status:             value.Status,
				Package_Code:       value.PackageCode,
				Package_Type:       value.PackageType,
				Is_Expired:         value.IsExpired,
				Start_Date:         utils.HandleNullableDate(value.StartDate.Time),
				Expired_Date:       utils.HandleNullableDate(value.ExpiredDate.Time),
				Time_Duration:      value.TimeDuration,
				Progress:           value.Progress.Float64,
				Progress_Count:     int(value.ProgressCount.Int64),
				Progress_Index:     int(value.ProgressIndex.Int64),
				Progress_Subindex:  int(value.ProgressSubindex.Int64),
				Pre_Test_Scores:    value.PreTestScores.Float64,
				Post_Test_Scores:   value.PostTestScores.Float64,
				Post_Test_Date:     utils.HandleNullableDate(value.PostTestDate.Time),
				Pre_Test_Total:     int(value.PreTestTotal.Int64),
				Post_Test_Total:    int(value.PostTestTotal.Int64),
				Total_Consultation: int(value.TotalConsultation.Int64),
				Total_Webinar:      int(value.TotalWebinar.Int64),
				Is_Active:          value.IsActive,
				Created_By:         value.CreatedBy,
				Created_Date:       value.CreatedDate,
				Modified_By:        value.ModifiedBy.String,
				Modified_Date:      value.ModifiedDate.Time,
				Is_Deleted:         value.IsDeleted,
			})
		}
	}

	if len(userClassResult) == 0 {
		return classEmpty, count, err
	}

	return userClassResult, count, err
}

func (u *userClassUsecase) GetAllUserClassQuran(ctx *gin.Context, r model.UserClassRequest) ([]shape.UserClass, int, error) {
	email := ctx.Request.Header.Get("email")

	users, err := u.userRepository.GetUserInfo(email)
	if err != nil {
		return nil, 0, utils.WrapError(err, "userClassUsecase.GetUserInfo")
	}

	var orderDefault = "ORDER BY code asc"
	var filterDefault = fmt.Sprintf("is_deleted = false and is_active = true and user_code='%s'", users.Code)

	filterFinal := utils.GetFilterOrderHandler(filterDefault, orderDefault, r.Query)

	result, err := u.userClassRepository.GetAllUserClassQuran(filterFinal)
	if err != nil {
		return nil, 0, utils.WrapError(err, "userClassUsecase.GetAllUserClassQuran")
	}

	userClassEmpty := make([]shape.UserClass, 0)
	if len(result) == 0 {
		return userClassEmpty, 0, err
	}

	return result, len(result), nil
}

func (userClassUsecase *userClassUsecase) UpdateUserClassProgress(userClass shape.UserClassPost, email string) (bool, error) {

	dataUserClass := model.UserClass{
		ID:       userClass.ID,
		UserCode: userClass.User_Code,
		Status:   userClass.Status,
		Progress: sql.NullFloat64{
			Float64: userClass.Progress,
		},
		ProgressCount: sql.NullInt64{
			Int64: userClass.Progress_Count,
		},
		ProgressIndex: sql.NullInt64{
			Int64: userClass.Progress_Index,
		},
		ProgressSubindex: sql.NullInt64{
			Int64: userClass.Progress_Subindex,
		},
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}

	result, err := userClassUsecase.userClassRepository.UpdateUserClassProgress(dataUserClass)
	if err != nil {
		return false, utils.WrapError(err, "userClassUsecase.userClassRepository.UpdateUserClassProgress : ")
	}

	return result, err
}

func (userClassUsecase *userClassUsecase) CheckAllUserClassExpired() {
	var promoCode = "BLJEXPD"
	var types string = "TodayClassExp"
	var table string = "TRANSACT USER CLASS"
	var emailType string = "Class Has Been Expired"

	var dataEmail model.EmailBody
	var email = userClassUsecase.sytemConfig.System.EmailSystem

	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}

		userClassList, err := userClassUsecase.userClassRepository.CheckAllUserClassExpired()
		if err != nil {
			utils.PushLogf("userClassUsecase.userClassRepository.CheckAllUserClassExpired : ", err.Error())
		}

		enum, err := userClassUsecase.enumRepository.GetEnumSplit(types)
		if err != nil {
			utils.PushLogf("userClassUsecase.enumRepository.GetEnumSplit : ", err.Error())
		}

		firstloop = false
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

			_, err := userClassUsecase.userClassRepository.UpdateUserClassExpired(dataUserClass)
			if err != nil {
				utils.PushLogf("userClassUsecase.userClassRepository.UpdateUserClassExpired : ", err.Error())
			}

			filterDefault := fmt.Sprintf(`is_deleted=false AND is_active=true AND class_code='%s' 
				AND (code = '%s' OR promo_code = '%s') LIMIT 1`, dataUserClass.ClassCode, "", promoCode)

			if err == nil {
				promotions, err := userClassUsecase.promotionRepository.GetAllPromotions(filterDefault)
				for _, promotion := range *promotions {
					dataEmail = model.EmailBody{
						BodyTemp:      emailType,
						UserCode:      value.UserCode,
						PromoDiscount: fmt.Sprintf(`%d`, int(promotion.Discount)),
					}
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
						Time: value.ExpiredDate.Time,
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

				_, err = userClassUsecase.notificationRepository.InsertNotification(dataNotification)
				if err != nil {
					utils.PushLogf("userClassUsecase.notificationRepository.InsertNotification : ", err.Error())
				}

			} else {
				utils.PushLogf("userClassUsecase.notificationRepository.CheckAllUserClassExpired : ", err.Error())
			}
		}
	}
}

func (userClassUsecase *userClassUsecase) CheckAllUserClass7DaysBeforeExpired() {
	var minutes float64 = 11520
	var promoCode = "BLJEXPD"
	var table string = "TRANSACT USER CLASS"
	var types string = "7DaysBeforeClassExp"
	var emailType string = "7 Days Before Class Expired"

	var dataEmail model.EmailBody
	var email = userClassUsecase.sytemConfig.System.EmailSystem

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
		if err != nil {
			utils.PushLogf("userClassUsecase.userClassRepository.UpdateUserClassExpired : ", err.Error())
		}

		enum, err := userClassUsecase.enumRepository.GetEnumSplit(types)
		if err != nil {
			utils.PushLogf("userClassUsecase.enumRepository.GetEnumSplit : ", err.Error())
		}

		firstloop = false
		for _, value := range userClassList {
			filterNotif := fmt.Sprintf(`AND user_class_code = %d`, value.ID)

			notification, err := userClassUsecase.notificationRepository.GetNotification(filterNotif, types)
			if err != nil {
				utils.PushLogf("userClassUsecase.notificationRepository.GetNotification : ", err.Error())
			}

			if utils.GetDuration(notification.ExpiredDate.Time, value.ExpiredDate.Time) > minutes {
				promotions, err := userClassUsecase.promotionRepository.GetAllPromotions(promoCode)
				for _, promotion := range *promotions {
					dataEmail = model.EmailBody{
						BodyTemp:      emailType,
						UserCode:      value.UserCode,
						PromoDiscount: fmt.Sprintf(`%d`, int(promotion.Discount)),
					}
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
						Time: value.ExpiredDate.Time,
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
				_, err = userClassUsecase.notificationRepository.InsertNotification(dataNotification)
				if err != nil {
					utils.PushLogf("userClassUsecase.notificationRepository.InsertNotification : ", err.Error())
				}
			}
		}
	}
}

func (userClassUsecase *userClassUsecase) CheckAllUserClass5DaysBeforeExpired() {
	var minutes float64 = 8640
	var promoCode = "BLJEXPD"
	var table string = "TRANSACT USER CLASS"
	var types string = "5DaysBeforeClassExp"
	var emailType string = "5 Days Before Class Expired"

	var dataEmail model.EmailBody
	var email = userClassUsecase.sytemConfig.System.EmailSystem

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
		if err != nil {
			utils.PushLogf("userClassUsecase.userClassRepository.CheckAllUserClassBeforeExpired : ", err.Error())
		}

		enum, err := userClassUsecase.enumRepository.GetEnumSplit(types)
		if err != nil {
			utils.PushLogf("userClassUsecase.enumRepository.GetEnumSplit : ", err.Error())
		}

		firstloop = false
		for _, value := range userClassList {
			filter := fmt.Sprintf(`AND user_class_code = %d`, value.ID)
			notification, err := userClassUsecase.notificationRepository.GetNotification(filter, types)
			if err != nil {
				utils.PushLogf("userClassUsecase.notificationRepository.GetNotification : ", err.Error())
			}

			if utils.GetDuration(notification.ExpiredDate.Time, value.ExpiredDate.Time) > minutes {
				promotions, err := userClassUsecase.promotionRepository.GetAllPromotions(promoCode)
				for _, promotion := range *promotions {
					dataEmail = model.EmailBody{
						BodyTemp:      emailType,
						UserCode:      value.UserCode,
						PromoDiscount: fmt.Sprintf(`%d`, int(promotion.Discount)),
					}
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
						Time: value.ExpiredDate.Time,
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

				_, err = userClassUsecase.notificationRepository.InsertNotification(dataNotification)
				if err != nil {
					utils.PushLogf("userClassUsecase.notificationRepository.InsertNotification : ", err.Error())
				}
			}
		}
	}
}

func (userClassUsecase *userClassUsecase) CheckAllUserClass2DaysBeforeExpired() {
	var minutes float64 = 4320
	var promoCode = "BLJEXPD"
	var table string = "TRANSACT USER CLASS"
	var types string = "2DaysBeforeClassExp"
	var emailType string = "2 Days Before Class Expired"

	var dataEmail model.EmailBody
	var email = userClassUsecase.sytemConfig.System.EmailSystem

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
		if err != nil {
			utils.PushLogf("userClassUsecase.userClassRepository.CheckAllUserClassBeforeExpired : ", err.Error())
		}

		enum, err := userClassUsecase.enumRepository.GetEnumSplit(types)
		if err != nil {
			utils.PushLogf("userClassUsecase.enumRepository.GetEnumSplit : ", err.Error())
		}

		firstloop = false
		for _, value := range userClassList {
			filter := fmt.Sprintf(`AND user_class_code = %d`, value.ID)
			notification, err := userClassUsecase.notificationRepository.GetNotification(filter, types)
			if err != nil {
				utils.PushLogf("userClassUsecase.notificationRepository.GetNotification : ", err.Error())
			}

			if utils.GetDuration(notification.ExpiredDate.Time, value.ExpiredDate.Time) > minutes {
				promotions, err := userClassUsecase.promotionRepository.GetAllPromotions(promoCode)
				for _, promotion := range *promotions {
					dataEmail = model.EmailBody{
						BodyTemp:      emailType,
						UserCode:      value.UserCode,
						PromoDiscount: fmt.Sprintf(`%d`, int(promotion.Discount)),
					}
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
						Time: value.ExpiredDate.Time,
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

				_, err = userClassUsecase.notificationRepository.InsertNotification(dataNotification)
				if err != nil {
					utils.PushLogf("userClassUsecase.notificationRepository.InsertNotification : ", err.Error())
				}
			}
		}
	}
}

func (userClassUsecase *userClassUsecase) CheckAllUserClass1DaysBeforeExpired() {
	var minutes float64 = 2880
	var promoCode = "BLJEXPD"
	var table string = "TRANSACT USER CLASS"
	var types string = "1DaysBeforeClassExp"
	var emailType string = "1 Days Before Class Expired"

	var dataEmail model.EmailBody
	var email = userClassUsecase.sytemConfig.System.EmailSystem

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
		if err != nil {
			utils.PushLogf("userClassUsecase.userClassRepository.CheckAllUserClassBeforeExpired : ", err.Error())
		}

		enum, err := userClassUsecase.enumRepository.GetEnumSplit(types)
		if err != nil {
			utils.PushLogf("userClassUsecase.enumRepository.GetEnumSplit : ", err.Error())
		}

		firstloop = false
		for _, value := range userClassList {
			filter := fmt.Sprintf(`AND user_class_code = %d`, value.ID)
			notification, err := userClassUsecase.notificationRepository.GetNotification(filter, types)
			if err != nil {
				utils.PushLogf("userClassUsecase.notificationRepository.GetNotification : ", err.Error())
			}

			if utils.GetDuration(notification.ExpiredDate.Time, value.ExpiredDate.Time) > minutes {
				promotions, err := userClassUsecase.promotionRepository.GetAllPromotions(promoCode)
				for _, promotion := range *promotions {
					dataEmail = model.EmailBody{
						BodyTemp:      emailType,
						UserCode:      value.UserCode,
						PromoDiscount: fmt.Sprintf(`%d`, int(promotion.Discount)),
					}
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
						Time: value.ExpiredDate.Time,
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

				_, err = userClassUsecase.notificationRepository.InsertNotification(dataNotification)
				if err != nil {
					utils.PushLogf("userClassUsecase.notificationRepository.InsertNotification : ", err.Error())
				}
			}
		}
	}
}
