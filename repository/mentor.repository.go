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
	GetMentorInfo(email string) (model.MentorInfo, error)

	GetAllMentor(skip, take int, sort, search, filter string) ([]model.MentorInfo, error)
	GetAllMentorCount(filter string) (int, error)

	GetAllMentorSchedule(code string) ([]model.MentorSchedule, error)
	GetAllMentorExperience(code string) ([]model.MentorExperience, error)
}

func InitMentorRepository(db *sqlx.DB) MentorRepository {
	return &mentorRepository{
		db,
	}
}

func (mentorRepository *mentorRepository) GetMentorInfo(email string) (model.MentorInfo, error) {
	var mentorRow model.MentorInfo
	row := mentorRepository.db.QueryRow(`
	SELECT
		id,
		code,
		role_code,
		role,
		mentor_code, 
		email,
		fullname,
		phone,
		profession,
		gender,
		age,
		birth,
		province,
		city,
		address,
		image_profile,
		description,
		account_owner,
		account_name,
		account_number,
		learning_method,
		learning_method_text,
		rating,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date
	FROM 
		belajariah.v_mentors
	WHERE 
		email = $1`, email)
	var isActive bool
	var rating float64
	var id, mentorCode int
	var createdDate time.Time
	var phone, age sql.NullInt64
	var modifiedDate, birth sql.NullTime
	var emailUsr, roleCode, role, createdBy, code string
	var fullname, profession, gender, province, city, address, imageProfile, modifiedBy,
		accountName, accountOwner, accountNumber, description, learningMethod, learningMethodText sql.NullString

	sqlError := row.Scan(
		&id,
		&code,
		&roleCode,
		&role,
		&mentorCode,
		&emailUsr,
		&fullname,
		&phone,
		&profession,
		&gender,
		&age,
		&birth,
		&province,
		&city,
		&address,
		&imageProfile,
		&description,
		&accountOwner,
		&accountName,
		&accountNumber,
		&learningMethod,
		&learningMethodText,
		&rating,
		&isActive,
		&createdBy,
		&createdDate,
		&modifiedBy,
		&modifiedDate,
	)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetMentorInfo => ", sqlError.Error())
		return model.MentorInfo{}, nil
	} else {
		mentorRow = model.MentorInfo{
			ID:                 id,
			Code:               code,
			RoleCode:           roleCode,
			Role:               role,
			MentorCode:         mentorCode,
			Email:              emailUsr,
			FullName:           fullname,
			Phone:              phone,
			Profession:         profession,
			Gender:             gender,
			Age:                age,
			Province:           province,
			City:               city,
			Address:            address,
			ImageProfile:       imageProfile,
			Description:        description,
			AccountOwner:       accountOwner,
			AccountName:        accountName,
			AccountNumber:      accountNumber,
			LearningMethod:     learningMethod,
			LearningMethodText: learningMethodText,
			Rating:             rating,
			IsActive:           isActive,
			CreatedBy:          createdBy,
			CreatedDate:        createdDate,
			ModifiedBy:         modifiedBy,
			ModifiedDate:       modifiedDate,
		}
		return mentorRow, sqlError
	}
}

func (mentorRepository *mentorRepository) GetAllMentor(skip, take int, sort, search, filter string) ([]model.MentorInfo, error) {
	var mentorList []model.MentorInfo
	query := fmt.Sprintf(`
		SELECT
			id,
			code,
			role_code,
			role,
			mentor_code, 
			email,
			fullname,
			phone,
			profession,
			gender,
			age,
			birth,
			province,
			city,
			address,
			image_profile,
			description,
			account_owner,
			account_name,
			account_number,
			learning_method,
			learning_method_text,
			rating,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date
		FROM 
			belajariah.v_mentors
		WHERE 
			is_active=true
		%s %s %s
		OFFSET %d
		LIMIT %d
	`, filter, search, sort, skip, take)

	rows, sqlError := mentorRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllMentor => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var isActive bool
			var rating float64
			var id, mentorCode int
			var createdDate time.Time
			var phone, age sql.NullInt64
			var modifiedDate, birth sql.NullTime
			var emailUsr, roleCode, role, createdBy, code string
			var fullname, profession, gender, province, city, address, imageProfile, modifiedBy,
				accountName, accountOwner, accountNumber, description, learningMethod, learningMethodText sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&roleCode,
				&role,
				&mentorCode,
				&emailUsr,
				&fullname,
				&phone,
				&profession,
				&gender,
				&age,
				&birth,
				&province,
				&city,
				&address,
				&imageProfile,
				&description,
				&accountOwner,
				&accountName,
				&accountNumber,
				&learningMethod,
				&learningMethodText,
				&rating,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllMentor => ", sqlError.Error())
			} else {
				mentorList = append(mentorList,
					model.MentorInfo{
						ID:                 id,
						Code:               code,
						RoleCode:           roleCode,
						Role:               role,
						MentorCode:         mentorCode,
						Email:              emailUsr,
						FullName:           fullname,
						Phone:              phone,
						Profession:         profession,
						Gender:             gender,
						Age:                age,
						Province:           province,
						City:               city,
						Address:            address,
						ImageProfile:       imageProfile,
						Description:        description,
						AccountOwner:       accountOwner,
						AccountName:        accountName,
						AccountNumber:      accountNumber,
						LearningMethod:     learningMethod,
						LearningMethodText: learningMethodText,
						Rating:             rating,
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

func (mentorRepository *mentorRepository) GetAllMentorSchedule(code string) ([]model.MentorSchedule, error) {
	var mentorList []model.MentorSchedule
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		mentor_code,
		shift_name,
		start_date,
		end_date,
		time_zone,
		coalesce(sequence, 0) AS sequence,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM 
		master.master_mentor_schedule
	WHERE 
		is_active=true and 
		is_deleted=false and
		mentor_code='%s'
	`, code)

	rows, sqlError := mentorRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllMentorSchedule => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id, sequence int
			var isActive, isDeleted bool
			var modifiedDate sql.NullTime
			var modifiedBy sql.NullString
			var startAt, endAt, createdDate time.Time
			var createdBy, shiftName, mentorCode, code, timeZone string

			sqlError := rows.Scan(
				&id,
				&code,
				&mentorCode,
				&shiftName,
				&startAt,
				&endAt,
				&timeZone,
				&sequence,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&isDeleted,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllMentorSchedule => ", sqlError.Error())
			} else {
				mentorList = append(mentorList,
					model.MentorSchedule{
						ID:           id,
						Code:         code,
						MentorCode:   mentorCode,
						ShiftName:    shiftName,
						StartDate:    startAt,
						EndDate:      endAt,
						TimeZone:     timeZone,
						Sequence:     sequence,
						IsActive:     isActive,
						CreatedBy:    createdBy,
						CreatedDate:  createdDate,
						ModifiedBy:   modifiedBy,
						ModifiedDate: modifiedDate,
						IsDeleted:    isDeleted,
					},
				)
			}
		}
	}
	return mentorList, sqlError
}

func (mentorRepository *mentorRepository) GetAllMentorExperience(code string) ([]model.MentorExperience, error) {
	var mentorList []model.MentorExperience
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		mentor_code,
		experience,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM 
		belajariah.mentor_experience
	WHERE 
		is_active=true and 
		is_deleted=false and
		mentor_code='%s'
	`, code)

	rows, sqlError := mentorRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllMentorExperience => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var createdDate time.Time
			var isActive, isDeleted bool
			var modifiedDate sql.NullTime
			var modifiedBy sql.NullString
			var createdBy, experience, mentorCode, code string

			sqlError := rows.Scan(
				&id,
				&code,
				&mentorCode,
				&experience,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&isDeleted,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllMentorExperience => ", sqlError.Error())
			} else {
				mentorList = append(mentorList,
					model.MentorExperience{
						ID:           id,
						Code:         code,
						MentorCode:   mentorCode,
						Experience:   experience,
						IsActive:     isActive,
						CreatedBy:    createdBy,
						CreatedDate:  createdDate,
						ModifiedBy:   modifiedBy,
						ModifiedDate: modifiedDate,
						IsDeleted:    isDeleted,
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
		belajariah.v_mentors  
	WHERE 
		is_active=true
	%s
	`, filter)

	row := mentorRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllMentorCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}
