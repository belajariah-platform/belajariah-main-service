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
		rating,
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
	var id, taskCompleted, taskInprogress int
	var isActive bool
	var rating float64
	var createdDate time.Time
	var phone, age sql.NullInt64
	var modifiedDate sql.NullTime
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
		taskInprogress,
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
		rating,
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
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := mentorRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllMentor => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var id, taskCompleted, taskInprogress int
			var isActive bool
			var rating float64
			var createdDate time.Time
			var phone, age sql.NullInt64
			var modifiedDate sql.NullTime
			var emailUsr, roleCode, role, createdBy string
			var fullname, profession, gender, province, city, address, imageCode, imageFilename, imageFilepath, modifiedBy sql.NullString

			sqlError := rows.Scan(
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
				taskInprogress,
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
