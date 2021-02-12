package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
	"fmt"
)

type consultationUsecase struct {
	consultationRepository   repository.ConsultationRepository
	approvalStatusRepository repository.ApprovalStatusRepository
}

type ConsultationUsecase interface {
	GetAllConsultation(query model.Query) ([]shape.Consultation, int, error)
	GetAllConsultationUser(query model.Query, userObj model.UserInfo) ([]shape.Consultation, int, error)
	GetAllConsultationMentor(query model.Query, userObj model.Mentor) ([]shape.Consultation, int, error)
}

func InitConsultationUsecase(consultationRepository repository.ConsultationRepository, approvalStatusRepository repository.ApprovalStatusRepository) ConsultationUsecase {
	return &consultationUsecase{
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
