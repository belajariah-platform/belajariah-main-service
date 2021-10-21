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
	GetMentorInfo(email string) (shape.Mentor, error)
	GetAllMentor(query model.Query) ([]shape.Mentor, int, error)
}

func InitMentorUsecase(mentorRepository repository.MentorRepository) MentorUsecase {
	return &mentorUsecase{
		mentorRepository,
	}
}

func (mentorUsecase *mentorUsecase) GetMentorInfo(email string) (shape.Mentor, error) {
	mentor, err := mentorUsecase.mentorRepository.GetMentorInfo(email)
	if mentor == (model.Mentor{}) {
		return shape.Mentor{}, nil
	}
	mentorResult := shape.Mentor{
		ID:              mentor.ID,
		Role_Code:       mentor.RoleCode,
		Role:            mentor.Role,
		Class_Code:      mentor.ClassCode,
		Email:           mentor.Email,
		Full_Name:       mentor.FullName.String,
		Phone:           int(mentor.Phone.Int64),
		Profession:      mentor.Profession.String,
		Gender:          mentor.Gender.String,
		Age:             int(mentor.Age.Int64),
		Province:        mentor.Province.String,
		City:            mentor.City.String,
		Address:         mentor.Address.String,
		Description:     mentor.Description.String,
		Image_Code:      mentor.ImageCode.String,
		Image_Filename:  mentor.ImageFilename.String,
		Image_Filepath:  mentor.ImageFilepath.String,
		Rating:          mentor.Rating,
		Task_Completed:  mentor.TaskCompleted,
		Task_Inprogress: mentor.TaskInprogress,
		Is_Active:       mentor.IsActive,
		Created_By:      mentor.CreatedBy,
		Created_Date:    mentor.CreatedDate,
		Modified_By:     mentor.ModifiedBy.String,
		Modified_Date:   mentor.ModifiedDate.Time,
	}
	return mentorResult, err
}

func (mentorUsecase *mentorUsecase) GetAllMentor(query model.Query) ([]shape.Mentor, int, error) {
	var mentors []model.Mentor
	var mentorResult []shape.Mentor
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

	mentors, err := mentorUsecase.mentorRepository.GetAllMentor(query.Skip, query.Take, sorting, search, filterQuery)
	count, errCount := mentorUsecase.mentorRepository.GetAllMentorCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range mentors {

			var mentorSchedule []model.MentorSchedule
			var mentorScheduleResult []shape.MentorSchedule
			mentorSchedule, err := mentorUsecase.mentorRepository.GetAllMentorSchedule(value.MentorCode)
			if err == nil {
				for _, schedule := range mentorSchedule {
					mentorScheduleResult = append(mentorScheduleResult, shape.MentorSchedule{
						ID:            schedule.ID,
						Mentor_Code:   schedule.MentorCode,
						Shift_Name:    schedule.ShiftName,
						Start_At:      schedule.StartAt,
						End_At:        schedule.EndAt,
						Is_Active:     schedule.IsActive,
						Created_By:    schedule.CreatedBy,
						Created_Date:  schedule.CreatedDate,
						Modified_By:   schedule.ModifiedBy.String,
						Modified_Date: schedule.ModifiedDate.Time,
						Time_Zone:     schedule.TimeZone.String,
					})
				}
			}

			mentorResult = append(mentorResult, shape.Mentor{
				ID:                   value.ID,
				Role_Code:            value.RoleCode,
				Role:                 value.Role,
				Mentor_Code:          value.MentorCode,
				Class_Code:           value.ClassCode,
				Email:                value.Email,
				Full_Name:            value.FullName.String,
				Phone:                int(value.Phone.Int64),
				Profession:           value.Profession.String,
				Gender:               value.Gender.String,
				Age:                  int(value.Age.Int64),
				Province:             value.Province.String,
				City:                 value.City.String,
				Address:              value.Address.String,
				Description:          value.Description.String,
				Image_Code:           value.ImageCode.String,
				Image_Filename:       value.ImageFilename.String,
				Image_Filepath:       value.ImageFilepath.String,
				Rating:               value.Rating,
				Learning_Method:      value.LearningMethod.String,
				Learning_Method_Text: value.LearningMethodText.String,
				Task_Completed:       value.TaskCompleted,
				Task_Inprogress:      value.TaskInprogress,
				Is_Active:            value.IsActive,
				Created_By:           value.CreatedBy,
				Created_Date:         value.CreatedDate,
				Modified_By:          value.ModifiedBy.String,
				Modified_Date:        value.ModifiedDate.Time,
				Schedule:             mentorScheduleResult,
			})
		}
	}
	mentorEmpty := make([]shape.Mentor, 0)
	if len(mentorResult) == 0 {
		return mentorEmpty, count, err
	}
	return mentorResult, count, err
}
