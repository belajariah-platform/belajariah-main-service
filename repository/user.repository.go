package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
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

	LoginUser(users model.Users) (model.Users, error)
	RegisterUser(users model.Users) (string, bool, error)
	UpdateProfileUser(users model.UserInfo) (bool, error)
	VerifyUser(users model.Users) (model.Users, bool, error)
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

	row := userRepository.db.QueryRow(`
	SELECT 
		count(*)
	FROM 
		auth.users 
	WHERE 
		verified_code = $1
	`, users.VerifiedCode.String,
	)
	sqlError := row.Scan(&count)

	if sqlError != nil {
		utils.PushLogf("SQL error on CheckVerifyCodeUser => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (userRepository *userRepository) CheckValidateLogin(users model.Users) (model.Users, error) {
	var userLogin model.Users
	query := fmt.Sprintf(`
	SELECT 
		email,
		password
	FROM 
		auth.users	
	WHERE
		email='%s'
	`, users.Email)

	row := userRepository.db.QueryRow(query)

	var email, password string
	sqlError := row.Scan(
		&email,
		&password,
	)
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
	var userRow model.UserInfo
	row := userRepository.db.QueryRow(`
	SELECT
		id,
		role_code,
		role, 
		email,
		full_name
	FROM 
		auth.v_users
	WHERE 
		code = $1`, code)

	var id int
	var emailUsr, roleCode, role string
	var fullname sql.NullString

	sqlError := row.Scan(
		&id,
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
	row := userRepository.db.QueryRow(`
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
		email = $1`, email)

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
		return model.UserInfo{}, sqlError
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
	var userRow model.UserInfo

	row := userRepository.db.QueryRow(`
	SELECT
		id,
		email
	FROM 
		auth.users
	WHERE 
	verified_code = $1`, code)

	var id int
	var emailUsr string

	sqlError := row.Scan(
		&id,
		&emailUsr,
	)
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
	var userRow model.Users
	row := userRepository.db.QueryRow(`
	SELECT 
		email,
		is_verified
	FROM 
		auth.users 
	WHERE 
		email = $1 AND
		password = $2 
	`, users.Email, users.Password)

	var emailUsr string
	var isVerified bool

	sqlError := row.Scan(
		&emailUsr,
		&isVerified,
	)
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

func (userRepository *userRepository) VerifyUser(users model.Users) (model.Users, bool, error) {
	var err error
	var result bool
	var user model.Users

	tx, errTx := userRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in verify account", errTx)
	} else {
		user, err = verifyUser(tx, users)
		if err != nil {
			utils.PushLogf("err---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf(users.VerifiedCode.String, "failed to verify account")
	}

	return user, result, err
}

func verifyUser(tx *sql.Tx, users model.Users) (model.Users, error) {
	var code string
	var userRow model.Users
	err := tx.QueryRow(`
	UPDATE
 		auth.users
	 SET
		is_verified=true,
		verified_code=default
 	WHERE
		verified_code=$1
		returning code
	`,
		users.VerifiedCode.String,
	).Scan(&code)

	if err != nil {
		utils.PushLogf("SQL error on Return verify user => ", err.Error())
		return model.Users{}, nil
	} else {
		userRow = model.Users{
			Code: code,
		}
		return userRow, err
	}
}

func (userRepository *userRepository) ResetVerificationCode(users model.Users) (model.Users, bool, error) {
	var err error
	var result bool
	var user model.Users

	tx, errTx := userRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in ResetVerificationCode", errTx)
	} else {
		user, err = resetVerificationCode(tx, users)
		if err != nil {
			utils.PushLogf("err in authentication", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to ResetVerificationCode")
	}
	return user, result, err
}

func resetVerificationCode(tx *sql.Tx, users model.Users) (model.Users, error) {
	var code string
	var userRow model.Users
	err := tx.QueryRow(`
	UPDATE
 		auth.users
	 SET
		verified_code=$1
 	WHERE
 		email=$2
		returning code
	`,
		users.VerifiedCode.String,
		users.Email,
	).Scan(&code)

	if err != nil {
		utils.PushLogf("SQL error on Return reset verification code => ", err.Error())
		return model.Users{}, nil
	} else {
		userRow = model.Users{
			Code: code,
		}
		return userRow, err
	}
}

func (userRepository *userRepository) RegisterUser(users model.Users) (string, bool, error) {
	var err error
	var result bool
	var email, code string

	tx, errTx := userRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in register", errTx)
	} else {
		code, email, err = registerUser(tx, users)
		if err == nil {
			insertUserDetail(tx, email, users)
		} else if err != nil {
			utils.PushLogf("err---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf(users.Email, "failed to register")
	}

	return code, result, err
}

func registerUser(tx *sql.Tx, users model.Users) (string, string, error) {
	var email, code string
	err := tx.QueryRow(`
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
		(SELECT code FROM auth.user_role WHERE "role" = 'Users'),
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8
	) returning email, code
	`,
		users.Email,
		users.Password,
		users.VerifiedCode.String,
		users.IsVerified,
		users.CreatedBy,
		users.CreatedDate,
		users.ModifiedBy.String,
		users.ModifiedDate.Time,
	).Scan(&email, &code)
	return code, email, err
}

func insertUserDetail(tx *sql.Tx, email string, users model.Users) error {
	_, err := tx.Exec(`
	INSERT INTO auth.user_detail
	(
		user_code,
		full_name,
		phone,
		created_by,
		created_date,
		modified_by,
		modified_date
	)
	VALUES(
		(SELECT id FROM auth.users WHERE email = $1),
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	);
	`,
		email,
		users.FullName.String,
		users.Phone.Int64,
		users.CreatedBy,
		users.CreatedDate,
		users.ModifiedBy.String,
		users.ModifiedDate.Time,
	)
	return err
}

func (userRepository *userRepository) UpdateProfileUser(users model.UserInfo) (bool, error) {
	var err error
	var result bool

	tx, errTx := userRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in update profile", errTx)
	} else {
		err = updateProfile(tx, users)
		if err != nil {
			utils.PushLogf("err---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf(users.Email, "failed to update profile")
	}

	return result, err
}

func updateProfile(tx *sql.Tx, users model.UserInfo) error {

	_, err := tx.Exec(`
	UPDATE
		auth.user_detail
	 SET
		full_name=$1,
		phone=$2,
		profession=$3,
		gender=$4,
		birth=$5,
		province=$6,
		city=$7,
		address=$8,
		modified_by=$9,
		modified_date=$10
 	WHERE
 		user_code=$11
	`,
		users.FullName.String,
		users.Phone.Int64,
		users.Profession.String,
		users.Gender.String,
		users.Birth.Time,
		users.Province.String,
		users.City.String,
		users.Address.String,
		users.ModifiedBy.String,
		users.ModifiedDate.Time,
		users.ID,
	)
	return err
}

func (userRepository *userRepository) ChangePassword(users model.Users) (model.Users, bool, error) {
	var err error
	var result bool
	var user model.Users

	tx, errTx := userRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in change password", errTx)
	} else {
		user, err = changePassword(tx, users)
		if err != nil {
			utils.PushLogf("err---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf(users.VerifiedCode.String, "failed to change password")
	}

	return user, result, err
}

func changePassword(tx *sql.Tx, users model.Users) (model.Users, error) {
	var id int
	var userRow model.Users
	err := tx.QueryRow(`
	UPDATE
 		auth.users
	 SET
	 	password=$1,
		verified_code=default,
		modified_by=$2,
		modified_date=$3
 	WHERE
 		email=$4
		returning id
	`,
		users.Password,
		users.ModifiedBy,
		users.ModifiedDate,
		users.Email,
	).Scan(&id)

	if err != nil {
		utils.PushLogf("SQL error on Return changePassword => ", err)
		return model.Users{}, nil
	} else {
		userRow = model.Users{
			ID: id,
		}
		return userRow, err
	}
}
