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

type consultationUsecase struct {
	userRepository           repository.UserRepository
	enumRepository           repository.EnumRepository
	consultationRepository   repository.ConsultationRepository
	approvalStatusRepository repository.ApprovalStatusRepository
}

type ConsultationUsecase interface {
	GetAllConsultation(query model.Query) ([]shape.Consultation, int, error)
	GetAllConsultationLimit(query model.Query, userObj model.UserHeader) ([]shape.Consultation, error)
	GetAllConsultationUser(query model.Query, userObj model.UserHeader) ([]shape.Consultation, int, error)
	GetAllConsultationMentor(query model.Query, userObj model.UserHeader) ([]shape.Consultation, int, error)

	ReadConsultation(consultation shape.ConsultationPost, email string) (bool, error)
	InsertConsultation(consultation shape.ConsultationPost, email string) (bool, error)
	UpdateConsultation(consultation shape.ConsultationPost, email string) (bool, error)
	ConfirmConsultation(consultation shape.ConsultationPost, email string) (bool, error)

	CheckAllConsultationExpired()
	CheckConsultationSpamUser(userObj model.UserInfo) (int, error)
	CheckConsultationSpamMentor(userObj model.Mentor) (int, error)
}

func InitConsultationUsecase(userRepository repository.UserRepository, enumRepository repository.EnumRepository, consultationRepository repository.ConsultationRepository, approvalStatusRepository repository.ApprovalStatusRepository) ConsultationUsecase {
	return &consultationUsecase{
		userRepository,
		enumRepository,
		consultationRepository,
		approvalStatusRepository,
	}
}

func (consultationUsecase *consultationUsecase) GetAllConsultation(query model.Query) ([]shape.Consultation, int, error) {
	var consultations []model.Consultation
	var consultationsResult []shape.Consultation
	var filterQuery, filterUser, sorting, search string

	if len(query.Order) > 0 {
		sorting = strings.Replace(query.Order, "|", " ", 1)
		sorting = "ORDER BY " + sorting
	}
	if len(query.Search) > 0 {
		search = `AND (LOWER(user_name) LIKE LOWER('%` + query.Search + `%') 
		OR LOWER(created_by) LIKE LOWER('%` + query.Search + `%'))`
	}

	filterUser = fmt.Sprintf(``)
	filterQuery = utils.GetFilterHandler(query.Filters)

	consultations, err := consultationUsecase.consultationRepository.GetAllConsultation(query.Skip, query.Take, sorting, search, filterQuery, filterUser)
	count, errCount := consultationUsecase.consultationRepository.GetAllConsultationCount(filterQuery, filterUser)

	if err == nil && errCount == nil {
		for _, value := range consultations {
			consultationsResult = append(consultationsResult, shape.Consultation{
				ID:                 value.ID,
				User_Code:          value.UserCode,
				User_Name:          value.UserName,
				Class_Code:         value.ClassCode,
				Class_Name:         value.ClassName,
				Recording_Code:     int(value.RecordingCode.Int64),
				Recording_Path:     value.RecordingPath.String,
				Recording_Name:     value.RecordingName.String,
				Recording_Duration: int(value.RecordingDuration.Int64),
				Status_Code:        value.StatusCode,
				Status:             value.Status,
				Description:        value.Description.String,
				Taken_Code:         int(value.TakenCode.Int64),
				Taken_Name:         value.TakenName.String,
				Is_Play:            value.IsPlay.Bool,
				Is_Read:            value.IsPlay.Bool,
				Is_Action_Taken:    value.IsActionTaken.Bool,
				Is_Active:          value.IsActive,
				Created_By:         value.CreatedBy,
				Created_Date:       value.CreatedDate,
				Modified_By:        value.ModifiedBy.String,
				Modified_Date:      value.ModifiedDate.Time,
				Deleted_By:         value.DeletedBy.String,
				Deleted_Date:       value.DeletedDate.Time,
			})
		}
	}
	consultationEmpty := make([]shape.Consultation, 0)
	if len(consultationsResult) == 0 {
		return consultationEmpty, count, err
	}
	return consultationsResult, count, err
}

func (consultationUsecase *consultationUsecase) GetAllConsultationLimit(query model.Query, userObj model.UserHeader) ([]shape.Consultation, error) {
	var filterQuery, filterUser string
	var consultations []model.Consultation
	var consultationsResult []shape.Consultation

	filterQuery = utils.GetFilterHandler(query.Filters)
	filterUser = fmt.Sprintf(`AND taken_code=%d`, userObj.ID)

	consultations, err := consultationUsecase.consultationRepository.GetAllConsultationLimit(query.Skip, query.Take, filterQuery, filterUser)

	if err == nil {
		for _, value := range consultations {
			consultationsResult = append(consultationsResult, shape.Consultation{
				ID:                 value.ID,
				User_Code:          value.UserCode,
				User_Name:          value.UserName,
				Class_Code:         value.ClassCode,
				Class_Name:         value.ClassName,
				Recording_Code:     int(value.RecordingCode.Int64),
				Recording_Path:     value.RecordingPath.String,
				Recording_Name:     value.RecordingName.String,
				Recording_Duration: int(value.RecordingDuration.Int64),
				Status_Code:        value.StatusCode,
				Status:             value.Status,
				Description:        value.Description.String,
				Taken_Code:         int(value.TakenCode.Int64),
				Taken_Name:         value.TakenName.String,
				Is_Play:            value.IsPlay.Bool,
				Is_Read:            value.IsPlay.Bool,
				Is_Action_Taken:    value.IsActionTaken.Bool,
				Is_Active:          value.IsActive,
				Created_By:         value.CreatedBy,
				Created_Date:       value.CreatedDate,
				Modified_By:        value.ModifiedBy.String,
				Modified_Date:      value.ModifiedDate.Time,
				Deleted_By:         value.DeletedBy.String,
				Deleted_Date:       value.DeletedDate.Time,
			})
		}
	}
	consultationEmpty := make([]shape.Consultation, 0)
	if len(consultationsResult) == 0 {
		return consultationEmpty, err
	}
	return consultationsResult, err
}

func (consultationUsecase *consultationUsecase) GetAllConsultationUser(query model.Query, userObj model.UserHeader) ([]shape.Consultation, int, error) {
	var consultations []model.Consultation
	var consultationsResult []shape.Consultation
	var filterQuery, filterUser, sorting, search string

	if len(query.Order) > 0 {
		sorting = strings.Replace(query.Order, "|", " ", 1)
		sorting = "ORDER BY " + sorting
	}
	if len(query.Search) > 0 {
		search = `AND (LOWER(user_name) LIKE LOWER('%` + query.Search + `%') 
		OR LOWER(email) LIKE LOWER('%` + query.Search + `%'))`
	}

	filterQuery = utils.GetFilterHandler(query.Filters)
	filterUser = fmt.Sprintf(`AND (user_code=%d OR taken_code=%d)`, userObj.ID, userObj.ID)

	consultations, err := consultationUsecase.consultationRepository.GetAllConsultation(query.Skip, query.Take, sorting, search, filterQuery, filterUser)
	count, errCount := consultationUsecase.consultationRepository.GetAllConsultationCount(filterQuery, filterUser)

	if err == nil && errCount == nil {
		for _, value := range consultations {
			consultationsResult = append(consultationsResult, shape.Consultation{
				ID:                 value.ID,
				User_Code:          value.UserCode,
				User_Name:          value.UserName,
				Class_Code:         value.ClassCode,
				Class_Name:         value.ClassName,
				Recording_Code:     int(value.RecordingCode.Int64),
				Recording_Path:     value.RecordingPath.String,
				Recording_Name:     value.RecordingName.String,
				Recording_Duration: int(value.RecordingDuration.Int64),
				Status_Code:        value.StatusCode,
				Status:             value.Status,
				Description:        value.Description.String,
				Taken_Code:         int(value.TakenCode.Int64),
				Taken_Name:         value.TakenName.String,
				Is_Play:            value.IsPlay.Bool,
				Is_Read:            value.IsPlay.Bool,
				Is_Action_Taken:    value.IsActionTaken.Bool,
				Is_Active:          value.IsActive,
				Created_By:         value.CreatedBy,
				Created_Date:       value.CreatedDate,
				Modified_By:        value.ModifiedBy.String,
				Modified_Date:      value.ModifiedDate.Time,
				Deleted_By:         value.DeletedBy.String,
				Deleted_Date:       value.DeletedDate.Time,
			})
		}
	}
	consultationEmpty := make([]shape.Consultation, 0)
	if len(consultationsResult) == 0 {
		return consultationEmpty, count, err
	}
	return consultationsResult, count, err
}

func (consultationUsecase *consultationUsecase) GetAllConsultationMentor(query model.Query, userObj model.UserHeader) ([]shape.Consultation, int, error) {
	var consultations []model.Consultation
	var consultationsResult []shape.Consultation
	var filterQuery, filterUser, sorting, search string

	if len(query.Order) > 0 {
		sorting = strings.Replace(query.Order, "|", " ", 1)
		sorting = "ORDER BY " + sorting
	}
	if len(query.Search) > 0 {
		search = `AND (LOWER(user_name) LIKE LOWER('%` + query.Search + `%') 
		OR LOWER(email) LIKE LOWER('%` + query.Search + `%'))`
	}
	filterQuery = utils.GetFilterHandler(query.Filters)
	filterUser = fmt.Sprintf(`AND (user_code=%d OR taken_code=%d)`, userObj.ID, userObj.ID)

	consultations, err := consultationUsecase.consultationRepository.GetAllConsultation(query.Skip, query.Take, sorting, search, filterQuery, filterUser)
	count, errCount := consultationUsecase.consultationRepository.GetAllConsultationCount(filterQuery, filterUser)

	if err == nil && errCount == nil {
		for _, value := range consultations {
			consultationsResult = append(consultationsResult, shape.Consultation{
				ID:                 value.ID,
				User_Code:          value.UserCode,
				User_Name:          value.UserName,
				Class_Code:         value.ClassCode,
				Class_Name:         value.ClassName,
				Recording_Code:     int(value.RecordingCode.Int64),
				Recording_Path:     value.RecordingPath.String,
				Recording_Name:     value.RecordingName.String,
				Recording_Duration: int(value.RecordingDuration.Int64),
				Status_Code:        value.StatusCode,
				Status:             value.Status,
				Description:        value.Description.String,
				Taken_Code:         int(value.TakenCode.Int64),
				Taken_Name:         value.TakenName.String,
				Is_Play:            value.IsPlay.Bool,
				Is_Read:            value.IsPlay.Bool,
				Is_Action_Taken:    value.IsActionTaken.Bool,
				Is_Active:          value.IsActive,
				Created_By:         value.CreatedBy,
				Created_Date:       value.CreatedDate,
				Modified_By:        value.ModifiedBy.String,
				Modified_Date:      value.ModifiedDate.Time,
				Deleted_By:         value.DeletedBy.String,
				Deleted_Date:       value.DeletedDate.Time,
			})
		}
	}
	consultationEmpty := make([]shape.Consultation, 0)
	if len(consultationsResult) == 0 {
		return consultationEmpty, count, err
	}
	return consultationsResult, count, err
}

func (consultationUsecase *consultationUsecase) InsertConsultation(consultation shape.ConsultationPost, email string) (bool, error) {
	var enum model.Enum
	var user model.UserInfo
	var consultations model.Consultation

	var err error
	var takenCode int
	var statusCode string
	var result, isActionTaken bool
	var approved string = "Approved"
	var completed string = "Completed"
	var waitingForResponse string = "Waiting for Response"

	status, err := consultationUsecase.approvalStatusRepository.GetApprovalStatus(consultation.Status_Code)
	switch strings.ToLower(consultation.Action) {
	case "approved":
		filterUser := fmt.Sprintf(`AND user_code=%d AND class_code='%s' AND expired_date='%s'`,
			consultation.User_Code,
			consultation.Class_Code,
			utils.CurrentDateStringCustom(consultation.Expired_Date),
		)
		filterMentor := fmt.Sprintf(`AND user_code=%d AND class_code='%s' AND expired_date='%s'`,
			consultation.Taken_Code,
			consultation.Class_Code,
			utils.CurrentDateStringCustom(consultation.Expired_Date),
		)

		user, err = consultationUsecase.userRepository.GetUserInfo(email)
		if user.Role == "Mentor" {
			consultations, err = consultationUsecase.consultationRepository.GetConsultation(filterMentor)
			statusCode = status.ApprovedStatus.String
			takenCode = int(consultations.UserCode)
			isActionTaken = consultations.IsActionTaken.Bool
			if consultations.Status == approved {
				return result, err
			}
		} else {
			consultations, err = consultationUsecase.consultationRepository.GetConsultation(filterUser)
			if consultations == (model.Consultation{}) {
				statusCode = consultation.Status_Code
			} else if consultations != (model.Consultation{}) && consultations.Status == waitingForResponse {
				enum, err = consultationUsecase.enumRepository.GetEnum(waitingForResponse)
				statusCode = enum.Code
				takenCode = int(consultations.TakenCode.Int64)
				isActionTaken = consultations.IsActionTaken.Bool
			} else if consultations != (model.Consultation{}) && consultations.Status == completed {
				enum, err = consultationUsecase.enumRepository.GetEnum(waitingForResponse)
				statusCode = enum.Code
				takenCode = int(consultations.TakenCode.Int64)
				isActionTaken = consultations.IsActionTaken.Bool
			} else if consultations != (model.Consultation{}) && consultations.Status == approved {
				enum, err = consultationUsecase.enumRepository.GetEnum(approved)
				statusCode = enum.Code
			} else {
				statusCode = consultation.Status_Code
			}
		}

	case "rejected":
		statusCode = status.RejectStatus.String
	default:
		statusCode = ""
	}

	if err == nil {
		dataConsultation := model.Consultation{
			UserCode:  consultation.User_Code,
			ClassCode: consultation.Class_Code,
			RecordingCode: sql.NullInt64{
				Int64: consultation.Recording_Code,
			},
			RecordingDuration: sql.NullInt64{
				Int64: consultation.Recording_Duration,
			},
			StatusCode: statusCode,
			Description: sql.NullString{
				String: consultation.Description,
			},
			TakenCode: sql.NullInt64{
				Int64: int64(takenCode),
			},
			IsActionTaken: sql.NullBool{
				Bool: isActionTaken,
			},
			ExpiredDate: sql.NullTime{
				Time: consultation.Expired_Date,
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
		if user.Role == "Mentor" {
			dataConsultationMentor := model.Consultation{
				UserCode:   int(consultation.Taken_Code),
				ClassCode:  consultation.Class_Code,
				StatusCode: statusCode,
				TakenCode: sql.NullInt64{
					Int64: int64(consultation.User_Code),
				},
				ExpiredDate: sql.NullTime{
					Time: consultation.Expired_Date,
				},
				IsActionTaken: sql.NullBool{
					Bool: true,
				},
				ModifiedBy: sql.NullString{
					String: email,
				},
				ModifiedDate: sql.NullTime{
					Time: time.Now(),
				},
			}
			result, err = consultationUsecase.consultationRepository.UpdateConsultation(dataConsultationMentor, consultation.Status_Code)
		}
		result, err = consultationUsecase.consultationRepository.InsertConsultation(dataConsultation)
	}

	return result, err
}

func (consultationUsecase *consultationUsecase) ReadConsultation(consultation shape.ConsultationPost, email string) (bool, error) {

	dataConsultation := model.Consultation{
		ID: consultation.ID,
	}
	result, err := consultationUsecase.consultationRepository.ReadConsultation(dataConsultation)
	return result, err
}

func (consultationUsecase *consultationUsecase) UpdateConsultation(consultation shape.ConsultationPost, email string) (bool, error) {
	var statusCode string

	status, err := consultationUsecase.approvalStatusRepository.GetApprovalStatus(consultation.Status_Code)
	switch strings.ToLower(consultation.Action) {
	case "approved":
		statusCode = status.ApprovedStatus.String
	case "rejected":
		statusCode = status.RejectStatus.String
	case "revised":
		statusCode = status.ReviseStatus.String
	default:
		statusCode = ""
	}
	dataConsultation := model.Consultation{
		ID:         consultation.ID,
		UserCode:   consultation.User_Code,
		ClassCode:  consultation.Class_Code,
		StatusCode: statusCode,
		TakenCode: sql.NullInt64{
			Int64: consultation.Taken_Code,
		},
		ExpiredDate: sql.NullTime{
			Time: consultation.Expired_Date,
		},
		IsActionTaken: sql.NullBool{
			Bool: consultation.Is_Action_Taken,
		},
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}
	result, err := consultationUsecase.consultationRepository.UpdateConsultation(dataConsultation, consultation.Status_Code)
	return result, err
}

func (consultationUsecase *consultationUsecase) ConfirmConsultation(consultation shape.ConsultationPost, email string) (bool, error) {
	var statusCode string

	status, err := consultationUsecase.approvalStatusRepository.GetApprovalStatus(consultation.Status_Code)
	switch strings.ToLower(consultation.Action) {
	case "approved":
		statusCode = status.ApprovedStatus.String
	case "rejected":
		statusCode = status.RejectStatus.String
	default:
		statusCode = ""
	}
	dataConsultation := model.Consultation{
		ID:         consultation.ID,
		UserCode:   consultation.User_Code,
		ClassCode:  consultation.Class_Code,
		StatusCode: statusCode,
		ExpiredDate: sql.NullTime{
			Time: consultation.Expired_Date,
		},
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}
	result, err := consultationUsecase.consultationRepository.ConfirmConsultation(dataConsultation, consultation.Status_Code)
	return result, err
}

func (consultationUsecase *consultationUsecase) CheckAllConsultationExpired() {
	var err error
	var consultationList []model.Consultation
	var email = "belajariah20@gmail.com"

	firstloop := true
	for {
		if !firstloop {
			time.Sleep(time.Minute)
		}
		consultationList, err = consultationUsecase.consultationRepository.CheckAllConsultationExpired()
		firstloop = false
		if err == nil {
			for _, value := range consultationList {
				dataConsultation := shape.ConsultationPost{
					ID:           value.ID,
					Action:       "Revised",
					User_Code:    value.UserCode,
					Class_Code:   value.ClassCode,
					Expired_Date: value.ExpiredDate.Time,
					Status_Code:  value.StatusCode,
				}
				result, err := consultationUsecase.UpdateConsultation(dataConsultation, email)
				if err != nil {
					utils.PushLogf("ERROR : ", err, result)
				}
			}
		}
	}
}

func (consultationUsecase *consultationUsecase) CheckConsultationSpamUser(userObj model.UserInfo) (int, error) {
	var err error
	var count int
	var consultationList []model.Consultation

	consultationList, err = consultationUsecase.consultationRepository.CheckConsultationSpam()
	if err == nil {
		for _, value := range consultationList {
			if value.UserCode == userObj.ID {
				count = count + 1
			}
		}
	}
	return count, err
}

func (consultationUsecase *consultationUsecase) CheckConsultationSpamMentor(userObj model.Mentor) (int, error) {
	var err error
	var count int
	var consultationList []model.Consultation

	consultationList, err = consultationUsecase.consultationRepository.CheckConsultationSpam()
	if err == nil {
		for _, value := range consultationList {
			if value.UserCode == userObj.ID {
				count = count + 1
			}
		}
	}
	return count, err
}
