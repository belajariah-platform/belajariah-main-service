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
	_getUserData = `
		SELECT
			id,
			code,
			role_code,
			role, 
			email,
			fullname
		FROM 
			belajariah.v_users
		WHERE 
			code = $1
	`
	_getUserInfo = `
		SELECT
			id,
			code,
			role_code,
			role, 
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
			is_verified,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date
		FROM 
			belajariah.v_users
		WHERE 
			email = $1
`
	_getEmailByVerifyCode = `
		SELECT
			id,
			email
		FROM 
			auth.users
		WHERE 
		verified_code = $1`

	_checkVerifyCodeUser = `
		SELECT 
			count(*)
		FROM 
			auth.users 
		WHERE 
			verified_code = $1
	`
	_CheckValidateLogin = `
		SELECT 
			email,
			password
		FROM 
			auth.users	
		WHERE
			email='%s'
`
	_loginUser = `
		SELECT 
			email,
			is_verified
		FROM 
			auth.users 
		WHERE 
			email = $1 AND
			password = $2 
`

	_registerUser = `
		INSERT INTO auth.users
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
			(SELECT code FROM auth.roles WHERE role = 'User'),
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
	_updateProfileUser = `
		UPDATE
			auth.user_info
		SET
			fullname=$1,
			phone=$2,
			profession=$3,
			gender=$4,
			birth=$5,
			province=$6,
			city=$7,
			address=$8,
			modified_by=$9,
			modified_date=$10
		WHERE user_code=$11
 	`

	_verifyUser = `
		UPDATE
			auth.users
		SET
			is_verified=true,
			verified_code=default
		WHERE 
			verified_code=$1
			returning code
`

	_changePassword = `
		UPDATE
			auth.users	
		SET
			password=$1,
		verified_code=default,
		modified_by=$2,
		modified_date=$3
		WHERE
			email=$4
		returning code
   `

	_resetVerificationCode = `
		UPDATE
			auth.users
		SET
		verified_code=$1
		WHERE
			email=$2
		returning code
	`
	_insertUserDetail = `
		INSERT INTO belajariah.user_info
		(
			user_code,
			fullname,
			phone,
			created_by,
			created_date,
			modified_by,
			modified_date
		)
		VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		)
	`
)

type userRepository struct {
	db *sqlx.DB
}

type UserRepository interface {
	GetUserData(code string) (model.UserInfo, error)
	GetUserInfo(email string) (model.UserInfo, error)
	GetEmailByVerifyCode(email string) (model.UserInfo, error)

	CheckVerifyCodeUser(users model.Users) (int, error)
	CheckValidateLogin(users model.Users) (model.Users, error)

	InsertUserDetail(data model.Users) (bool, error)
	LoginUser(users model.Users) (model.Users, error)
	UpdateProfileUser(users model.UserInfo) (bool, error)
	VerifyUser(users model.Users) (model.Users, bool, error)
	RegisterUser(users model.Users) (model.Users, bool, error)
	ChangePassword(users model.Users) (model.Users, bool, error)
	ResetVerificationCode(users model.Users) (model.Users, bool, error)
}

func InitUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db,
	}
}

func (userRepository *userRepository) CheckVerifyCodeUser(users model.Users) (int, error) {
	var count int
	row := userRepository.db.QueryRow(_checkVerifyCodeUser, users.VerifiedCode.String)
	sqlError := row.Scan(&count)

	if sqlError != nil {
		utils.PushLogf("SQL error on CheckVerifyCodeUser => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (userRepository *userRepository) CheckValidateLogin(users model.Users) (model.Users, error) {
	var userLogin model.Users
	var email, password string

	query := fmt.Sprintf(_CheckValidateLogin, users.Email)

	row := userRepository.db.QueryRow(query)
	sqlError := row.Scan(&email, &password)

	if sqlError != nil {
		utils.PushLogf("SQL error on CheckValidateLogin => ", sqlError.Error())
		return model.Users{}, nil
	} else {
		userLogin = model.Users{
			Email:    email,
			Password: password,
		}
		return userLogin, sqlError
	}
}

func (userRepository *userRepository) GetUserData(code string) (model.UserInfo, error) {
	var id int
	var emailUsr, roleCode, role, codes string
	var fullname sql.NullString
	var userRow model.UserInfo

	row := userRepository.db.QueryRow(_getUserData, code)
	sqlError := row.Scan(
		&id,
		&codes,
		&roleCode,
		&role,
		&emailUsr,
		&fullname,
	)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetUserData => ", sqlError.Error())
		return model.UserInfo{}, nil
	} else {
		userRow = model.UserInfo{
			ID:       id,
			Code:     codes,
			RoleCode: roleCode,
			Role:     role,
			Email:    emailUsr,
			FullName: fullname,
		}
		return userRow, sqlError
	}
}

func (userRepository *userRepository) GetUserInfo(email string) (model.UserInfo, error) {
	var userRow model.UserInfo
	row := userRepository.db.QueryRow(_getUserInfo, email)

	var id int
	var createdDate time.Time
	var phone, age sql.NullInt64
	var isVerified, isActive bool
	var modifiedDate, births sql.NullTime
	var emailUsr, roleCode, role, createdBy, code string
	var fullname, profession, gender, province, city, address, imageProfile, modifiedBy sql.NullString

	sqlError := row.Scan(
		&id,
		&code,
		&roleCode,
		&role,
		&emailUsr,
		&fullname,
		&phone,
		&profession,
		&gender,
		&age,
		&births,
		&province,
		&city,
		&address,
		&imageProfile,
		&isVerified,
		&isActive,
		&createdBy,
		&createdDate,
		&modifiedBy,
		&modifiedDate,
	)
	if sqlError != nil {
		return model.UserInfo{}, nil
	} else {
		userRow = model.UserInfo{
			ID:           id,
			Code:         code,
			RoleCode:     roleCode,
			Role:         role,
			Email:        emailUsr,
			FullName:     fullname,
			Phone:        phone,
			Profession:   profession,
			Gender:       gender,
			Age:          age,
			Birth:        births,
			Province:     province,
			City:         city,
			Address:      address,
			ImageProfile: imageProfile,
			IsVerified:   isVerified,
			IsActive:     isActive,
			CreatedBy:    createdBy,
			CreatedDate:  createdDate,
			ModifiedBy:   modifiedBy,
			ModifiedDate: modifiedDate,
		}

		return userRow, sqlError
	}
}

func (userRepository *userRepository) GetEmailByVerifyCode(code string) (model.UserInfo, error) {
	var id int
	var emailUsr string
	var userRow model.UserInfo

	row := userRepository.db.QueryRow(_getEmailByVerifyCode, code)
	sqlError := row.Scan(&id, &emailUsr)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetEmailByVerifyCode => ", sqlError.Error())
		return model.UserInfo{}, nil
	} else {
		userRow = model.UserInfo{
			ID:    id,
			Email: emailUsr,
		}
		return userRow, sqlError
	}
}

func (userRepository *userRepository) LoginUser(users model.Users) (model.Users, error) {
	var emailUsr string
	var isVerified bool
	var userRow model.Users

	row := userRepository.db.QueryRow(_loginUser, users.Email, users.Password)
	sqlError := row.Scan(&emailUsr, &isVerified)

	if sqlError != nil {
		utils.PushLogf("SQL error on LoginUser => ", sqlError.Error())
		return model.Users{}, nil
	} else {
		userRow = model.Users{
			Email:      emailUsr,
			IsVerified: isVerified,
		}
		return userRow, sqlError
	}
}

func (r *userRepository) VerifyUser(data model.Users) (model.Users, bool, error) {
	var code string
	var user model.Users

	tx, err := r.db.Beginx()
	if err != nil {
		return user, false, errors.New("userRepository: VerifyUser: error begin transaction")
	}

	err = tx.QueryRow(_verifyUser, data.VerifiedCode.String).Scan(&code)
	if err != nil {
		tx.Rollback()
		return user, false, utils.WrapError(err, "userRepository: VerifyUser: error update")
	}

	user = model.Users{Code: code}

	tx.Commit()
	return user, err == nil, nil
}

func (r *userRepository) ResetVerificationCode(data model.Users) (model.Users, bool, error) {
	var code string
	var user model.Users

	tx, err := r.db.Beginx()
	if err != nil {
		return user, false, errors.New("userRepository: ResetVerificationCode: error begin transaction")
	}

	err = tx.QueryRow(_resetVerificationCode, data.VerifiedCode.String, data.Email).Scan(&code)
	if err != nil {
		tx.Rollback()
		return user, false, utils.WrapError(err, "userRepository: ResetVerificationCode: error update")
	}

	user = model.Users{Code: code}

	tx.Commit()
	return user, err == nil, nil
}

func (r *userRepository) RegisterUser(data model.Users) (model.Users, bool, error) {
	var code string
	var user model.Users

	tx, err := r.db.Beginx()
	if err != nil {
		return user, false, errors.New("userRepository: RegisterUser: error begin transaction")
	}

	mutation := fmt.Sprintf(_registerUser,
		data.Email,
		data.Password,
		data.VerifiedCode.String,
		data.IsVerified,
		data.CreatedBy,
		utils.CurrentDateString(data.CreatedDate.UTC()),
		data.ModifiedBy.String,
		utils.CurrentDateString(data.ModifiedDate.Time.UTC()))

	err = tx.QueryRow(mutation).Scan(&code)
	user = model.Users{Code: code}

	if err != nil {
		tx.Rollback()
		return user, false, utils.WrapError(err, "userRepository: RegisterUser: error insert")
	}

	tx.Commit()
	return user, err == nil, nil
}

func (r *userRepository) InsertUserDetail(data model.Users) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("userRepository: InsertUserDetail: error begin transaction")
	}

	_, err = tx.Exec(_insertUserDetail,
		data.Code,
		data.FullName.String,
		data.Phone.Int64,
		data.CreatedBy,
		data.CreatedDate,
		data.ModifiedBy.String,
		data.ModifiedDate.Time,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "userRepository: InsertUserDetail: error insert")
	}

	tx.Commit()
	return err == nil, nil
}

func (r *userRepository) UpdateProfileUser(data model.UserInfo) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("userRepository: UpdateProfileUser: error begin transaction")
	}

	_, err = tx.Exec(_updateProfileUser,
		data.FullName.String,
		data.Phone.Int64,
		data.Profession.String,
		data.Gender.String,
		data.Birth.Time,
		data.Province.String,
		data.City.String,
		data.Address.String,
		data.ModifiedBy.String,
		data.ModifiedDate.Time,
		data.Code,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "userRepository: UpdateProfileUser: error update")
	}

	tx.Commit()
	return err == nil, nil
}

func (r *userRepository) ChangePassword(data model.Users) (model.Users, bool, error) {
	var code string
	var user model.Users

	tx, err := r.db.Beginx()
	if err != nil {
		return user, false, errors.New("userRepository: ChangePassword: error begin transaction")
	}

	err = tx.QueryRow(_changePassword,
		data.Password,
		data.ModifiedBy,
		data.ModifiedDate,
		data.Email,
	).Scan(&code)
	user = model.Users{Code: code}

	if err != nil {
		tx.Rollback()
		return user, false, utils.WrapError(err, "userRepository: ChangePassword: error update")
	}

	tx.Commit()
	return user, err == nil, nil
}
