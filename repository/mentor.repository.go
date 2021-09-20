package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type mentorRepository struct {
	db *sqlx.DB
}

type MentorRepository interface {
	GetMentorInfo(email string) (model.Mentor, error)
	GetAllMentor(skip, take int, sort, search, filter string) ([]model.Mentor, error)
	GetAllMentorSchedule(code int) ([]model.MentorSchedule, error)
	GetAllMentorCount(filter string) (int, error)
}

func InitMentorRepository(db *sqlx.DB) MentorRepository {
	return &mentorRepository{
		db,
	}
}

func (mentorRepository *mentorRepository) GetMentorInfo(email string) (model.Mentor, error) {
	var mentorRow model.Mentor
	row := mentorRepository.db.QueryRow(`
	SELECT
		id,
		role_code,
		role, 
		email,
		full_name,
		phone,
		profession,
		gender,
		age,
		province,
		city,
		address,
		image_code,
		image_filename,
		image_filepath,
		CASE
			WHEN rating IS NULL THEN 0
			ELSE rating
		END AS rating,
		task_completed,
		task_inprogress,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date
	FROM 
		private.v_mentors
	WHERE 
		email = $1`, email)
	var isActive bool
	var rating float64
	var createdDate time.Time
	var phone, age sql.NullInt64
	var modifiedDate sql.NullTime
	var id, taskCompleted, taskInprogress int
	var emailUsr, roleCode, role, createdBy string
	var fullname, profession, gender, province, city, address, imageCode, imageFilename, imageFilepath, modifiedBy sql.NullString

	sqlError := row.Scan(
		&id,
		&roleCode,
		&role,
		&emailUsr,
		&fullname,
		&phone,
		&profession,
		&gender,
		&age,
		&province,
		&city,
		&address,
		&imageCode,
		&imageFilename,
		&imageFilepath,
		&rating,
		&taskCompleted,
		&taskInprogress,
		&isActive,
		&createdBy,
		&createdDate,
		&modifiedBy,
		&modifiedDate,
	)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetMentorInfo => ", sqlError)
		return model.Mentor{}, nil
	} else {
		mentorRow = model.Mentor{
			ID:             id,
			RoleCode:       roleCode,
			Role:           role,
			Email:          emailUsr,
			FullName:       fullname,
			Phone:          phone,
			Profession:     profession,
			Gender:         gender,
			Age:            age,
			Province:       province,
			City:           city,
			Address:        address,
			ImageCode:      imageCode,
			ImageFilename:  imageFilename,
			ImageFilepath:  imageFilepath,
			Rating:         rating,
			TaskCompleted:  taskCompleted,
			TaskInprogress: taskInprogress,
			IsActive:       isActive,
			CreatedBy:      createdBy,
			CreatedDate:    createdDate,
			ModifiedBy:     modifiedBy,
			ModifiedDate:   modifiedDate,
		}
		return mentorRow, sqlError
	}
}

func (mentorRepository *mentorRepository) GetAllMentor(skip, take int, sort, search, filter string) ([]model.Mentor, error) {
	var mentorList []model.Mentor
	query := fmt.Sprintf(`
	SELECT
		id,
		role_code,
		role, 
		mentor_code,
		class_code,
		email,
		full_name,
		phone,
		profession,
		gender,
		age,
		province,
		city,
		address,
		image_code,
		image_filename,
		image_filepath,
		CASE
			WHEN rating IS NULL THEN 0
			ELSE rating
		END AS rating,
		learning_method,
		learning_method_text,
		task_completed,
		task_inprogress,
		description,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date
	FROM 
		private.v_mentors
	WHERE 
		is_active=true
	%s %s %s
	OFFSET %d
	LIMIT %d
	`, filter, search, sort, skip, take)

	rows, sqlError := mentorRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllMentor => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var id, taskCompleted, taskInprogress, mentorCode int
			var isActive bool
			var rating float64
			var createdDate time.Time
			var phone, age sql.NullInt64
			var modifiedDate sql.NullTime
			var emailUsr, roleCode, role, createdBy, classCode string
			var fullname, profession, gender, province, city, address, imageCode, imageFilename, imageFilepath, modifiedBy, description, learningMethodText, learningMethod sql.NullString

			sqlError := rows.Scan(
				&id,
				&roleCode,
				&role,
				&mentorCode,
				&classCode,
				&emailUsr,
				&fullname,
				&phone,
				&profession,
				&gender,
				&age,
				&province,
				&city,
				&address,
				&imageCode,
				&imageFilename,
				&imageFilepath,
				&rating,
				&learningMethod,
				&learningMethodText,
				&taskCompleted,
				&taskInprogress,
				&description,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllMentor => ", sqlError)
			} else {
				mentorList = append(mentorList,
					model.Mentor{
						ID:                 id,
						RoleCode:           roleCode,
						Role:               role,
						MentorCode:         mentorCode,
						ClassCode:          classCode,
						Email:              emailUsr,
						FullName:           fullname,
						Phone:              phone,
						Profession:         profession,
						Gender:             gender,
						Age:                age,
						Province:           province,
						City:               city,
						Address:            address,
						ImageCode:          imageCode,
						ImageFilename:      imageFilename,
						ImageFilepath:      imageFilepath,
						Rating:             rating,
						LearningMethod:     learningMethod,
						LearningMethodText: learningMethodText,
						TaskCompleted:      taskCompleted,
						TaskInprogress:     taskInprogress,
						Description:        description,
						IsActive:           isActive,
						CreatedBy:          createdBy,
						CreatedDate:        createdDate,
						ModifiedBy:         modifiedBy,
						ModifiedDate:       modifiedDate,
					},
				)
			}
		}
	}
	return mentorList, sqlError
}

func (mentorRepository *mentorRepository) GetAllMentorSchedule(code int) ([]model.MentorSchedule, error) {
	var mentorList []model.MentorSchedule
	query := fmt.Sprintf(`
	SELECT
		id,
		mentor_code,
		shift_name,
		start_at,
		end_at,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date
	FROM 
		public.master_mentor_schedule
	WHERE 
		is_active=true and 
		deleted_by is null and
		mentor_code=%d
	`, code)

	rows, sqlError := mentorRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllMentorSchedule => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var isActive bool
			var id, mentorCode int
			var modifiedDate sql.NullTime
			var modifiedBy sql.NullString
			var createdBy, shiftName string
			var startAt, endAt, createdDate time.Time

			sqlError := rows.Scan(
				&id,
				&mentorCode,
				&shiftName,
				&startAt,
				&endAt,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllMentorSchedule => ", sqlError)
			} else {
				mentorList = append(mentorList,
					model.MentorSchedule{
						ID:           id,
						MentorCode:   mentorCode,
						ShiftName:    shiftName,
						StartAt:      startAt,
						EndAt:        endAt,
						IsActive:     isActive,
						CreatedBy:    createdBy,
						CreatedDate:  createdDate,
						ModifiedBy:   modifiedBy,
						ModifiedDate: modifiedDate,
					},
				)
			}
		}
	}
	return mentorList, sqlError
}

func (mentorRepository *mentorRepository) GetAllMentorCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		private.v_mentors  
	WHERE 
		is_active=true
	%s
	`, filter)

	row := mentorRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllMentorCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}
