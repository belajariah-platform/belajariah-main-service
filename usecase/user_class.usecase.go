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
	enumRepository         repository.EnumRepository
	userClassRepository    repository.UserClassRepository
	notificationRepository repository.NotificationRepository
}

type UserClassUsecase interface {
	CheckAllUserClassExpired()
	CheckAllUserClass3DaysExpired()
	CheckAllUserClassOneDaysExpired()
	GetAllUserClass(query model.Query, userObj model.UserInfo) ([]shape.UserClass, int, error)
	UpdateUserClassProgress(userClass shape.UserClassPost, email string) (bool, error)
}

func InitUserClassUsecase(enumRepository repository.EnumRepository, userClassRepository repository.UserClassRepository, notificationRepository repository.NotificationRepository) UserClassUsecase {
	return &userClassUsecase{
		enumRepository,
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
	var types string = "TodayClassExp"
	var email = "belajariah20@gmail.com"
	var table string = "Transact user class"

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
					dataNotification := model.Notification{
						TableRef:         table,
						NotificationType: enum.Code,
						UserCode:         value.UserCode,
						Sequence:         1,
						CreatedBy:        email,
						CreatedDate:      time.Now(),
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

func (userClassUsecase *userClassUsecase) CheckAllUserClass3DaysExpired() {
	var times int = 259200
	var minutes float64 = 5760
	var table string = "Transact user class"
	var types string = "3DaysBeforeClassExp"
	var email string = "belajariah20@gmail.com"

	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}
		userClassList, err := userClassUsecase.userClassRepository.CheckAllUserClassBeforeExpired(times)
		enum, err := userClassUsecase.enumRepository.GetEnumSplit(types)
		firstloop = false
		if err == nil {
			for _, value := range userClassList {
				filter := fmt.Sprintf(`AND user_class_code = %d`, value.ID)
				notification, err := userClassUsecase.notificationRepository.GetNotification(filter, types)
				if err == nil && utils.GetDuration(notification.ExpiredDate.Time, value.ExpiredDate) > minutes {
					dataNotification := model.Notification{
						TableRef:         table,
						NotificationType: enum.Code,
						UserCode:         value.UserCode,
						Sequence:         1,
						CreatedBy:        email,
						CreatedDate:      time.Now(),
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

func (userClassUsecase *userClassUsecase) CheckAllUserClassOneDaysExpired() {
	var times int = 86400
	var minutes float64 = 2880
	var table string = "Transact user class"
	var types string = "1DaysBeforeClassExp"
	var email string = "belajariah20@gmail.com"

	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}
		userClassList, err := userClassUsecase.userClassRepository.CheckAllUserClassBeforeExpired(times)
		enum, err := userClassUsecase.enumRepository.GetEnumSplit(types)
		firstloop = false
		if err == nil {
			for _, value := range userClassList {
				filter := fmt.Sprintf(`AND user_class_code = %d`, value.ID)
				notification, err := userClassUsecase.notificationRepository.GetNotification(filter, types)
				if err == nil && utils.GetDuration(notification.ExpiredDate.Time, value.ExpiredDate) > minutes {
					dataNotification := model.Notification{
						TableRef:         table,
						NotificationType: enum.Code,
						UserCode:         value.UserCode,
						Sequence:         1,
						CreatedBy:        email,
						CreatedDate:      time.Now(),
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
