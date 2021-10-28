package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
	"strings"
)

type mentorUsecase struct {
	mentorRepository repository.MentorRepository
}

type MentorUsecase interface {
	GetMentorInfo(email string) (shape.MentorInfo, error)
	GetAllMentor(query model.Query) ([]shape.MentorInfo, int, error)

	RegisterMentor(r model.Mentors) (bool, error)
}

func InitMentorUsecase(mentorRepository repository.MentorRepository) MentorUsecase {
	return &mentorUsecase{
		mentorRepository,
	}
}

func (mentorUsecase *mentorUsecase) GetMentorInfo(email string) (shape.MentorInfo, error) {
	mentor, err := mentorUsecase.mentorRepository.GetMentorInfo(email)
	if err != nil {
		return shape.MentorInfo{}, utils.WrapError(err, "mentorUsecase.mentorRepository.GetMentorInfo : ")
	}

	mentorResult := shape.MentorInfo{
		ID:                   mentor.ID,
		Code:                 mentor.Code,
		Role_Code:            mentor.RoleCode,
		Role:                 mentor.Role,
		Mentor_Code:          mentor.MentorCode,
		Email:                mentor.Email,
		Full_Name:            mentor.FullName.String,
		Phone:                int(mentor.Phone.Int64),
		Profession:           mentor.Profession.String,
		Gender:               mentor.Gender.String,
		Age:                  int(mentor.Age.Int64),
		Province:             mentor.Province.String,
		City:                 mentor.City.String,
		Address:              mentor.Address.String,
		Birth:                mentor.Birth.Time,
		Description:          mentor.Description.String,
		ImageProfile:         mentor.ImageProfile.String,
		Account_Name:         mentor.AccountName.String,
		Account_Owner:        mentor.AccountNumber.String,
		Account_Number:       mentor.AccountNumber.String,
		Learning_Method:      mentor.LearningMethod.String,
		Learning_Method_Text: mentor.LearningMethodText.String,
		Rating:               mentor.Rating,
		Is_Active:            mentor.IsActive,
		Created_By:           mentor.CreatedBy,
		Created_Date:         mentor.CreatedDate,
		Modified_By:          mentor.ModifiedBy.String,
		Modified_Date:        mentor.ModifiedDate.Time,
	}
	return mentorResult, err
}

func (mentorUsecase *mentorUsecase) GetAllMentor(query model.Query) ([]shape.MentorInfo, int, error) {
	var mentors []model.MentorInfo
	var mentorResult []shape.MentorInfo
	var filterQuery, sorting, search string

	if len(query.Order) > 0 {
		sorting = strings.Replace(query.Order, "|", " ", 1)
		sorting = "ORDER BY " + sorting
	}
	if len(query.Search) > 0 {
		search = `AND (LOWER(full_name) LIKE LOWER('%` + query.Search + `%') 
		OR LOWER(email) LIKE LOWER('%` + query.Search + `%'))`
	}

	filterQuery = utils.GetFilterHandler(query.Filters)

	mentorEmpty := make([]shape.MentorInfo, 0)
	mentorScheduleEmpty := make([]shape.MentorSchedule, 0)
	mentorExperienceEmpty := make([]shape.MentorExperience, 0)

	mentors, err := mentorUsecase.mentorRepository.GetAllMentor(query.Skip, query.Take, sorting, search, filterQuery)
	if err != nil {
		return mentorEmpty, 0, utils.WrapError(err, "mentorUsecase.mentorRepository.GetAllMentor : ")
	}

	count, errCount := mentorUsecase.mentorRepository.GetAllMentorCount(filterQuery)
	if errCount != nil {
		return mentorEmpty, 0, utils.WrapError(errCount, "mentorUsecase.mentorRepository.GetAllMentorCount : ")
	}

	if err == nil && errCount == nil {
		for _, value := range mentors {
			var mentorSchedule []model.MentorSchedule
			var mentorScheduleResult []shape.MentorSchedule

			mentorSchedule, err := mentorUsecase.mentorRepository.GetAllMentorSchedule(value.Code)
			if err != nil {
				return mentorEmpty, 0, utils.WrapError(err, "mentorUsecase.mentorRepository.GetAllMentorSchedule : ")
			}

			if err == nil {
				for _, schedule := range mentorSchedule {
					mentorScheduleResult = append(mentorScheduleResult, shape.MentorSchedule{
						ID:            schedule.ID,
						Code:          schedule.Code,
						Mentor_Code:   schedule.MentorCode,
						Time_Zone:     schedule.TimeZone,
						Shift_Name:    schedule.ShiftName,
						Start_Date:    schedule.StartDate,
						End_Date:      schedule.EndDate,
						Sequence:      schedule.Sequence,
						Is_Active:     schedule.IsActive,
						Created_By:    schedule.CreatedBy,
						Created_Date:  schedule.CreatedDate,
						Modified_By:   schedule.ModifiedBy.String,
						Modified_Date: schedule.ModifiedDate.Time,
						Is_Deleted:    schedule.IsDeleted,
					})
				}
			}

			if len(mentorScheduleResult) == 0 {
				mentorScheduleResult = mentorScheduleEmpty
			}

			var mentorExperience []model.MentorExperience
			var mentorExperienceResult []shape.MentorExperience

			mentorExperience, errs := mentorUsecase.mentorRepository.GetAllMentorExperience(value.Code)
			if errs != nil {
				return mentorEmpty, 0, utils.WrapError(errs, "mentorUsecase.mentorRepository.GetAllMentorExperience : ")
			}

			if errs == nil {
				for _, experience := range mentorExperience {
					mentorExperienceResult = append(mentorExperienceResult, shape.MentorExperience{
						ID:            experience.ID,
						Code:          experience.Code,
						Mentor_Code:   experience.MentorCode,
						Experience:    experience.Experience,
						Is_Active:     experience.IsActive,
						Created_By:    experience.CreatedBy,
						Created_Date:  experience.CreatedDate,
						Modified_By:   experience.ModifiedBy.String,
						Modified_Date: experience.ModifiedDate.Time,
						Is_Deleted:    experience.IsDeleted,
					})
				}
			}

			if len(mentorExperienceResult) == 0 {
				mentorExperienceResult = mentorExperienceEmpty
			}

			mentorResult = append(mentorResult, shape.MentorInfo{
				ID:                   value.ID,
				Code:                 value.Code,
				Role_Code:            value.RoleCode,
				Role:                 value.Role,
				Mentor_Code:          value.MentorCode,
				Email:                value.Email,
				Full_Name:            value.FullName.String,
				Phone:                int(value.Phone.Int64),
				Profession:           value.Profession.String,
				Gender:               value.Gender.String,
				Age:                  int(value.Age.Int64),
				Province:             value.Province.String,
				City:                 value.City.String,
				Address:              value.Address.String,
				Birth:                value.Birth.Time,
				Description:          value.Description.String,
				ImageProfile:         value.ImageProfile.String,
				Account_Name:         value.AccountName.String,
				Account_Owner:        value.AccountNumber.String,
				Account_Number:       value.AccountNumber.String,
				Learning_Method:      value.LearningMethod.String,
				Learning_Method_Text: value.LearningMethodText.String,
				Rating:               value.Rating,
				Is_Active:            value.IsActive,
				Created_By:           value.CreatedBy,
				Created_Date:         value.CreatedDate,
				Modified_By:          value.ModifiedBy.String,
				Modified_Date:        value.ModifiedDate.Time,
				Mentor_Schedule:      mentorScheduleResult,
				Mentor_Experience:    mentorExperienceResult,
			})
		}
	}

	return mentorResult, count, err
}

func (mentorUsecase *mentorUsecase) RegisterMentor(r model.Mentors) (bool, error) {

	result, err := mentorUsecase.mentorRepository.RegisterMentor(r)
	if err != nil {
		return false, utils.WrapError(err, "mentorUsecase.mentorRepository.RegisterMentor : ")
	}

	return result, err
}
