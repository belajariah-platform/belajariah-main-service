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
	GetAllConsultationUser(query model.Query, userObj model.UserInfo) ([]shape.Consultation, int, error)
	GetAllConsultationMentor(query model.Query, userObj model.Mentor) ([]shape.Consultation, int, error)

	ReadConsultation(consultation shape.ConsultationPost, email string) (bool, error)
	InsertConsultation(consultation shape.ConsultationPost, email string) (bool, error)
	UpdateConsultation(consultation shape.ConsultationPost, email string) (bool, error)
	ConfirmConsultation(consultation shape.ConsultationPost, email string) (bool, error)

	CheckAllConsultationExpired()
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
	var filterQuery, filterUser string
	var consultations []model.Consultation
	var consultationsResult []shape.Consultation

	filterUser = fmt.Sprintf(``)
	filterQuery = utils.GetFilterHandler(query.Filters)

	consultations, err := consultationUsecase.consultationRepository.GetAllConsultation(query.Skip, query.Take, filterQuery, filterUser)
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

func (consultationUsecase *consultationUsecase) GetAllConsultationUser(query model.Query, userObj model.UserInfo) ([]shape.Consultation, int, error) {
	var filterQuery, filterUser string
	var consultations []model.Consultation
	var consultationsResult []shape.Consultation

	filterQuery = utils.GetFilterHandler(query.Filters)
	filterUser = fmt.Sprintf(`AND (user_code=%d OR taken_code=%d)`, userObj.ID, userObj.ID)

	consultations, err := consultationUsecase.consultationRepository.GetAllConsultation(query.Skip, query.Take, filterQuery, filterUser)
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

func (consultationUsecase *consultationUsecase) GetAllConsultationMentor(query model.Query, userObj model.Mentor) ([]shape.Consultation, int, error) {
	var filterQuery, filterUser string
	var consultations []model.Consultation
	var consultationsResult []shape.Consultation

	filterQuery = utils.GetFilterHandler(query.Filters)
	filterUser = fmt.Sprintf(`AND (user_code=%d OR taken_code=%d)`, userObj.ID, userObj.ID)

	consultations, err := consultationUsecase.consultationRepository.GetAllConsultation(query.Skip, query.Take, filterQuery, filterUser)
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
	var err error
	var result bool
	var statusCode string

	var enum model.Enum
	var user model.UserInfo
	var approved string = "Approved"
	var completed string = "Completed"
	var waitingForResponse string = "Waiting for Response"

	status, err := consultationUsecase.approvalStatusRepository.GetApprovalStatus(consultation.Status_Code)
	switch strings.ToLower(consultation.Action) {
	case "approved":
		filter := fmt.Sprintf(`AND user_code=%d AND class_code='%s' AND expired_date='%s'`,
			consultation.User_Code,
			consultation.Class_Code,
			consultation.Expired_Date,
		)
		user, err = consultationUsecase.userRepository.GetUserInfo(email)
		if user.Role == "Mentor" {
			statusCode = status.ApprovedStatus.String
		} else {
			consultations, _ := consultationUsecase.consultationRepository.GetConsultation(filter)
			if consultations == (model.Consultation{}) {
				statusCode = status.ApprovedStatus.String
			} else if consultations != (model.Consultation{}) && consultations.Status == waitingForResponse {
				enum, err = consultationUsecase.enumRepository.GetEnum(waitingForResponse)
				statusCode = enum.Code
			} else if consultations != (model.Consultation{}) && consultations.Status == completed {
				enum, err = consultationUsecase.enumRepository.GetEnum(waitingForResponse)
				statusCode = enum.Code
			} else if consultations != (model.Consultation{}) && consultations.Status == approved {
				enum, err = consultationUsecase.enumRepository.GetEnum(approved)
				statusCode = enum.Code
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
				Int64: consultation.Taken_Code,
			},
			ExpiredDate: sql.NullTime{
				Time: time.Now(),
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
			dataConsultation.IsActionTaken.Bool = true
			result, err = consultationUsecase.consultationRepository.UpdateConsultation(dataConsultation, consultation.Status_Code)
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
