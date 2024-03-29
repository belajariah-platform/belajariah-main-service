package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	_getMentorInfo = `
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
		country_number_code,
		rating,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date
	FROM 
		belajariah.v_mentors
	WHERE 
		email = $1`

	_getAllMentor = `
	SELECT
		id,
		code,
		class_code,
		mentor_code, 
		email,
		fullname,
		phone,
		profession,
		gender,
		age,
		province,
		city,
		address,
		image_profile,
		description,
		learning_method,
		learning_method_text,
		rating,
		minimum_rate,
		allow_contact_from,
		country_number_code,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date
	FROM 
		belajariah.v_mentor_class_detail
	WHERE 
		is_active=true
	%s %s %s
	OFFSET %d
	LIMIT %d
`

	_getAllMentorSchedule = `
	SELECT
		id,
		code,
		mentor_code,
		class_code,
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
		mentor_code='%s' and
		class_code='%s'
	`

	_getAllMentorExperience = `
	SELECT
		id,
		code,
		mentor_code,
		experience,
		experience_type,
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
	`
	_getAllMentorClass = `
	SELECT
		id,
		code,
		mentor_code,
		mentor_name,
		class_code,
		class_name,
		class_initial,
		coalesce(minimum_rate, 0) as minimum_rate,
		allow_contact_from,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM 
		belajariah.v_mentor_class
	WHERE 
		is_active=true and 
		is_deleted=false and
		mentor_code='%s'
	`

	_getAllMentorCount = `
		SELECT COUNT(*) FROM 
			belajariah.v_mentor_class_detail  
		WHERE 
			is_active=true
		%s
		`
	_registerMentor = `
		INSERT INTO auth.mentors
		(
			role_code,
			email,
			password,
			verified_code,
			is_verified,
			created_by,
			created_date,
			modified_by,
			modified_date
		)
		VALUES(
			(SELECT code FROM auth.roles WHERE role = 'Mentor' LIMIT 1),
			'%s',
			'%s',
			'%s',
			 %t,
			'%s',
			'%s',
			'%s',
			'%s'
		) returning code
		`
	_insertMentorDetail = `
		INSERT INTO belajariah.mentor_info
		(
			mentor_code,
			fullname,
			phone,
			created_by,
			created_date,
			modified_by,
			modified_date,
			profession,
			gender,
			province,
			city,
			description,
			age
		)
		VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13
		)
	`
)

type mentorRepository struct {
	db *sqlx.DB
}

type MentorRepository interface {
	GetMentorInfo(email string) (model.MentorInfo, error)

	GetAllMentor(skip, take int, sort, search, filter string) ([]model.MentorInfo, error)
	GetAllMentorCount(filter string) (int, error)

	GetAllMentorClass(code string) ([]model.MentorClass, error)
	GetAllMentorExperience(code string) ([]model.MentorExperience, error)
	GetAllMentorSchedule(code, classCode string) ([]model.MentorSchedule, error)

	RegisterMentor(data model.Mentors) (model.Mentors, bool, error)
	InsertMentorDetail(data model.Mentors) (bool, error)
}

func InitMentorRepository(db *sqlx.DB) MentorRepository {
	return &mentorRepository{
		db,
	}
}

func (r *mentorRepository) GetMentorInfo(email string) (model.MentorInfo, error) {
	var mentorRow model.MentorInfo
	row := r.db.QueryRow(_getMentorInfo, email)
	var isActive bool
	var rating float64
	var id, mentorCode int
	var createdDate time.Time
	var phone, age sql.NullInt64
	var modifiedDate, birth sql.NullTime
	var emailUsr, roleCode, role, createdBy, code string
	var fullname, profession, gender, province, city, address, imageProfile, modifiedBy,
		accountName, accountOwner, accountNumber, description, learningMethod, learningMethodText, countryNumberCode sql.NullString

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
		&countryNumberCode,
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
			CountryNumberCode:  countryNumberCode,
		}
		return mentorRow, sqlError
	}
}

func (r *mentorRepository) GetAllMentor(skip, take int, sort, search, filter string) ([]model.MentorInfo, error) {
	var mentorList []model.MentorInfo

	query := fmt.Sprintf(_getAllMentor, filter, search, sort, skip, take)
	rows, sqlError := r.db.Query(query)

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
			var minimumRate sql.NullInt64
			var modifiedDate sql.NullTime
			var emailUsr, createdBy, code, classCode string
			var fullname, profession, gender, province, city, address, imageProfile, modifiedBy,
				description, learningMethod, learningMethodText, allowContactFrom, countryNumberCode sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&classCode,
				&mentorCode,
				&emailUsr,
				&fullname,
				&phone,
				&profession,
				&gender,
				&age,
				&province,
				&city,
				&address,
				&imageProfile,
				&description,
				&learningMethod,
				&learningMethodText,
				&rating,
				&minimumRate,
				&allowContactFrom,
				&countryNumberCode,
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
						Class_Code:         classCode,
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
						LearningMethod:     learningMethod,
						LearningMethodText: learningMethodText,
						Rating:             rating,
						MinimumRate:        minimumRate,
						AllowContactFrom:   allowContactFrom,
						CountryNumberCode:  countryNumberCode,
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

func (r *mentorRepository) GetAllMentorSchedule(code, classCode string) ([]model.MentorSchedule, error) {
	var mentorList []model.MentorSchedule
	query := fmt.Sprintf(_getAllMentorSchedule, code, classCode)

	rows, sqlError := r.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllMentorSchedule => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id, sequence int
			var isActive, isDeleted bool
			var modifiedDate sql.NullTime
			var modifiedBy, classCode sql.NullString
			var startAt, endAt, createdDate time.Time
			var createdBy, shiftName, mentorCode, code, timeZone string

			sqlError := rows.Scan(
				&id,
				&code,
				&mentorCode,
				&classCode,
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
						ClassCode:    classCode.String,
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

func (r *mentorRepository) GetAllMentorExperience(code string) ([]model.MentorExperience, error) {
	var mentorList []model.MentorExperience
	query := fmt.Sprintf(_getAllMentorExperience, code)

	rows, sqlError := r.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllMentorExperience => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var createdDate time.Time
			var isActive, isDeleted bool
			var modifiedDate sql.NullTime
			var modifiedBy, experienceType sql.NullString
			var createdBy, experience, mentorCode, code string

			sqlError := rows.Scan(
				&id,
				&code,
				&mentorCode,
				&experience,
				&experienceType,
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
						ID:             id,
						Code:           code,
						MentorCode:     mentorCode,
						Experience:     experience,
						ExperienceType: experienceType,
						IsActive:       isActive,
						CreatedBy:      createdBy,
						CreatedDate:    createdDate,
						ModifiedBy:     modifiedBy,
						ModifiedDate:   modifiedDate,
						IsDeleted:      isDeleted,
					},
				)
			}
		}
	}
	return mentorList, sqlError
}

func (r *mentorRepository) GetAllMentorClass(code string) ([]model.MentorClass, error) {
	var mentorList []model.MentorClass
	query := fmt.Sprintf(_getAllMentorClass, code)

	rows, sqlError := r.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllMentorClass => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var createdDate time.Time
			var isActive, isDeleted bool
			var minimumRate sql.NullInt64
			var modifiedDate sql.NullTime
			var createdBy, classCode, className, mentorCode, code string
			var modifiedBy, classInitial, mentorName, allowContactFrom sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&mentorCode,
				&mentorName,
				&classCode,
				&className,
				&classInitial,
				&minimumRate,
				&allowContactFrom,
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
					model.MentorClass{
						ID:               id,
						Code:             code,
						MentorCode:       mentorCode,
						MentorName:       mentorName,
						ClassCode:        classCode,
						ClassName:        className,
						ClassInitial:     classInitial,
						MinimumRate:      minimumRate,
						AllowContactFrom: allowContactFrom,
						IsActive:         isActive,
						CreatedBy:        createdBy,
						CreatedDate:      createdDate,
						ModifiedBy:       modifiedBy,
						ModifiedDate:     modifiedDate,
						IsDeleted:        isDeleted,
					},
				)
			}
		}
	}
	return mentorList, sqlError
}

func (r *mentorRepository) GetAllMentorCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(_getAllMentorCount, filter)

	row := r.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllMentorCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (r *mentorRepository) RegisterMentor(data model.Mentors) (model.Mentors, bool, error) {
	var code string

	tx, err := r.db.Beginx()
	if err != nil {
		return model.Mentors{}, false, errors.New("mentorRepository: RegisterMentor: error begin transaction")
	}

	data.CreatedDate = time.Now()
	data.ModifiedDate.Time = time.Now()

	mutation := fmt.Sprintf(_registerMentor,
		data.Email,
		data.Password,
		data.VerifiedCode.String,
		data.IsVerified,
		data.Email,
		utils.CurrentDateString(data.CreatedDate.UTC()),
		data.Email,
		utils.CurrentDateString(data.ModifiedDate.Time.UTC()),
	)

	err = tx.QueryRow(mutation).Scan(&code)
	if err != nil {
		tx.Rollback()
		return model.Mentors{}, false, utils.WrapError(err, "mentorRepository: RegisterMentor: error insert")
	}

	tx.Commit()
	return model.Mentors{Code: code}, err == nil, nil
}

func (r *mentorRepository) InsertMentorDetail(data model.Mentors) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("mentorRepository: InsertMentorDetail: error begin transaction")
	}

	_, err = tx.Exec(_insertMentorDetail,
		data.Code,
		data.Full_Name,
		data.Phone.Int64,
		data.Email,
		utils.CurrentDateString(data.CreatedDate.UTC()),
		data.Email,
		utils.CurrentDateString(data.ModifiedDate.Time.UTC()),
		data.Profession,
		data.Gender,
		data.Province,
		data.City,
		data.Description,
		data.Age,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "mentorRepository: InsertMentorDetail: error insert")
	}

	tx.Commit()
	return err == nil, nil
}
