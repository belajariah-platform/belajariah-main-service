package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"
)

type userUsecase struct {
	emailUsecase   EmailUsecase
	userRepository repository.UserRepository
}

type UserUsecase interface {
	RegisterUser(users shape.Users) (bool, error, string)
	ResetVerificationUser(users shape.Users) (bool, error)
	LoginUser(users shape.Users) (shape.UserInfo, bool, error, string)
	UpdateProfileUser(users shape.UsersPost, email string) (bool, error)

	ChangePasswordPublic(users shape.Users) (bool, error)
	ChangePasswordPrivate(users shape.Users) (bool, error, string)

	GetUserInfo(email string) (shape.UserInfo, error)
	VerifyAccount(users shape.Users) (bool, error, string)
	GoogleLogin(users shape.Users) (shape.UserInfo, bool, error)

	CheckEmail(email string) (shape.UserInfo, error)
	VerifyEmail(users shape.Users) (string, int, error)
}

func InitUserUsecase(emailUsecase EmailUsecase, userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		emailUsecase,
		userRepository,
	}
}

func (userUsecase *userUsecase) LoginUser(users shape.Users) (shape.UserInfo, bool, error, string) {
	var msg string

	dataUser := model.Users{
		Email:    users.Email,
		Password: users.Password,
	}
	userLogin, err := userUsecase.userRepository.CheckValidateLogin(dataUser)
	if err != nil {
		return shape.UserInfo{}, false, utils.WrapError(err, "userUsecase.userRepository.CheckValidateLogin : "), msg
	}

	isPassword := utils.CheckPasswordHash(dataUser.Password, userLogin.Password)
	if userLogin == (model.Users{}) || !isPassword {
		msg = fmt.Sprintf(`Email dan kata sandi salah`)
		return shape.UserInfo{}, false, err, msg
	}

	user, err := userUsecase.GetUserInfo(dataUser.Email)
	if err != nil {
		return shape.UserInfo{}, false, utils.WrapError(err, "userUsecase.userRepository.GetUserInfo : "), msg
	}

	if !user.Is_Verified {
		msg = fmt.Sprintf(`Akun kamu belum terverifikasi`)
		return shape.UserInfo{}, false, err, msg
	}

	return user, true, err, msg
}

func (userUsecase *userUsecase) GoogleLogin(users shape.Users) (shape.UserInfo, bool, error) {
	var emailType string = "Registration Success"
	var result bool = true
	var userID string

	hashPassword, err := utils.GenerateHashPassword(users.Password)
	if err != nil {
		return shape.UserInfo{}, false, utils.WrapError(err, "utils.GenerateHashPassword : ")
	}

	dataUser := model.Users{
		Email:    users.Email,
		Password: hashPassword,
		FullName: sql.NullString{
			String: users.Full_Name,
		},
		Phone: sql.NullInt64{
			Int64: users.Phone,
		},
		IsVerified:  true,
		CreatedBy:   users.Email,
		CreatedDate: time.Now(),
		ModifiedBy: sql.NullString{
			String: users.Email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}

	userLogin, err := userUsecase.userRepository.CheckValidateLogin(dataUser)
	if err != nil {
		return shape.UserInfo{}, false, utils.WrapError(err, "userUsecase.userRepository.CheckValidateLogin : ")
	}

	if userLogin == (model.Users{}) {
		userID, result, err = userUsecase.userRepository.RegisterUser(dataUser)
		fmt.Println(userLogin)
		fmt.Println(userID)
		if err != nil {
			return shape.UserInfo{}, false, utils.WrapError(err, "userUsecase.userRepository.RegisterUser : ")
		}

		if userID != "" {
			dataEmail := model.EmailBody{
				UserCode: userID,
				BodyTemp: emailType,
			}
			userUsecase.emailUsecase.SendEmail(dataEmail)
		}
	}

	user, err := userUsecase.GetUserInfo(dataUser.Email)
	if err != nil {
		return shape.UserInfo{}, false, utils.WrapError(err, "userUsecase.userRepository.GetUserInfo : ")
	}

	if user == (shape.UserInfo{}) {
		return shape.UserInfo{}, false, err
	}

	return user, result, err
}

func (userUsecase *userUsecase) ResetVerificationUser(users shape.Users) (bool, error) {
	var emailType string = "Account Verification"
	var userAuth model.Users
	var result bool
	var err error

	dataUser := model.Users{
		Email: users.Email,
		VerifiedCode: sql.NullString{
			String: utils.
				GenerateVerifyCode(users.Email),
		},
		ModifiedBy: sql.NullString{
			String: users.Email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}

	userAuth, result, err = userUsecase.userRepository.ResetVerificationCode(dataUser)
	if err != nil {
		return false, utils.WrapError(err, "userUsecase.userRepository.ResetVerificationCode : ")
	}

	if userAuth.ID != 0 {
		dataEmail := model.EmailBody{
			UserCode:         userAuth.Code,
			BodyTemp:         emailType,
			VerificationCode: dataUser.VerifiedCode.String,
		}
		userUsecase.emailUsecase.SendEmail(dataEmail)
	}
	return result, err
}

func (userUsecase *userUsecase) RegisterUser(users shape.Users) (bool, error, string) {
	var emailType string = "Account Verification"
	var user model.UserInfo
	var msg, userID string
	var err error
	var result bool

	hashPassword, err := utils.GenerateHashPassword(users.Password)
	if err != nil {
		return false, utils.WrapError(err, "utils.GenerateHashPassword : "), msg
	}

	dataUser := model.Users{
		Email:    users.Email,
		Password: hashPassword,
		FullName: sql.NullString{
			String: users.Full_Name,
		},
		Phone: sql.NullInt64{
			Int64: users.Phone,
		},
		VerifiedCode: sql.NullString{
			String: utils.
				GenerateVerifyCode(users.Email),
		},
		IsVerified:  false,
		CreatedBy:   users.Email,
		CreatedDate: time.Now(),
		ModifiedBy: sql.NullString{
			String: users.Email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}

	user, err = userUsecase.userRepository.GetUserInfo(dataUser.Email)
	if err != nil {
		return false, utils.WrapError(err, "userUsecase.userRepository.GetUserInfo : "), msg
	}

	if user.Email == dataUser.Email && user.IsVerified {
		msg = fmt.Sprintf(`Email '%s' sudah ada`, dataUser.Email)
		return result, err, msg
	} else if user.Email == dataUser.Email && !user.IsVerified {
		msg = fmt.Sprintf(`Akun sudah terdaftar silahkan verifikasi email kamu`)
		return result, err, msg
	}

	firstLoop := true
	for firstLoop {
		count, err := userUsecase.userRepository.CheckVerifyCodeUser(dataUser)
		if err != nil {
			utils.PushLogf("userUsecase.userRepository.CheckVerifyCodeUser :", err)
		}

		if count == 0 {
			firstLoop = false
		} else {
			dataUser.VerifiedCode.String = utils.GenerateVerifyCode(dataUser.Email)
		}
	}

	userID, result, err = userUsecase.userRepository.RegisterUser(dataUser)
	if err != nil {
		return false, utils.WrapError(err, "userUsecase.userRepository.RegisterUser : "), msg
	}

	if err == nil && userID != "" {
		dataEmail := model.EmailBody{
			UserCode:         userID,
			BodyTemp:         emailType,
			VerificationCode: dataUser.VerifiedCode.String,
		}
		userUsecase.emailUsecase.SendEmail(dataEmail)
	}

	return result, err, msg
}

func (userUsecase *userUsecase) UpdateProfileUser(users shape.UsersPost, email string) (bool, error) {
	var result bool
	var err error
	dataUser := model.UserInfo{
		ID: users.User_Code,
		FullName: sql.NullString{
			String: users.Full_Name,
		},
		Phone: sql.NullInt64{
			Int64: users.Phone,
		},
		Profession: sql.NullString{
			String: users.Profession,
		},
		Gender: sql.NullString{
			String: users.Gender,
		},
		Birth: sql.NullTime{
			Time: users.Birth,
		},
		Province: sql.NullString{
			String: users.Province,
		},
		City: sql.NullString{
			String: users.City,
		},
		Address: sql.NullString{
			String: users.Address,
		},
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}

	result, err = userUsecase.userRepository.UpdateProfileUser(dataUser)
	if err != nil {
		return false, utils.WrapError(err, "userUsecase.userRepository.UpdateProfileUser : ")
	}

	return result, err
}

func (userUsecase *userUsecase) ChangePasswordPrivate(users shape.Users) (bool, error, string) {
	var err error
	var msg string
	var result bool

	hashPassword, err := utils.GenerateHashPassword(users.New_Password)
	if err != nil {
		return false, utils.WrapError(err, "utils.GenerateHashPassword : "), msg
	}

	dataUser := model.Users{
		Email:       users.Email,
		OldPassword: users.Old_Password,
		Password:    hashPassword,
	}

	userLogin, err := userUsecase.userRepository.CheckValidateLogin(dataUser)
	if err != nil {
		return false, utils.WrapError(err, "userUsecase.userRepository.CheckValidateLogin : "), msg
	}

	isPassword := utils.CheckPasswordHash(dataUser.OldPassword, userLogin.Password)
	if userLogin == (model.Users{}) || !isPassword {
		msg = fmt.Sprintf(`Kata sandi lama salah`)
		return result, err, msg
	}

	user, result, err := userUsecase.userRepository.ChangePassword(dataUser)
	if err != nil {
		return false, utils.WrapError(err, "userUsecase.userRepository.CheckValidateLogin : "), msg
	}

	utils.PushLogf("[SUCCESS CHANGE PASSWORD] :", user)

	return result, err, msg
}

func (userUsecase *userUsecase) ChangePasswordPublic(users shape.Users) (bool, error) {
	var emailType string = "Change Password"
	var user model.Users
	var result bool
	var err error

	hashPassword, err := utils.GenerateHashPassword(users.Password)
	if err != nil {
		return false, utils.WrapError(err, "utils.GenerateHashPassword : ")
	}

	dataUser := model.Users{
		Email:    users.Email,
		Password: hashPassword,
		ModifiedBy: sql.NullString{
			String: users.Email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}

	user, result, err = userUsecase.userRepository.ChangePassword(dataUser)
	if err != nil {
		return false, utils.WrapError(err, "userUsecase.userRepository.ChangePassword : ")
	}

	dataEmail := model.EmailBody{
		UserEmail: dataUser.Email,
		BodyTemp:  emailType,
		UserCode:  user.Code,
	}
	userUsecase.emailUsecase.SendEmail(dataEmail)

	return result, err
}

func (userUsecase *userUsecase) VerifyAccount(users shape.Users) (bool, error, string) {
	var emailType string = "Registration Success"
	var user model.Users
	var result bool
	var err error
	var msg string

	dataUser := model.Users{
		Email: users.Email,
		VerifiedCode: sql.NullString{
			String: users.Verified_Code,
		},
	}

	count, err := userUsecase.userRepository.CheckVerifyCodeUser(dataUser)
	if err != nil {
		return false, utils.WrapError(err, "userUsecase.userRepository.CheckVerifyCodeUser : "), msg
	}

	if count == 0 {
		msg = fmt.Sprintf(`Kode verifikasi salah`)
		return result, err, msg
	}

	user, result, err = userUsecase.userRepository.VerifyUser(dataUser)
	if err != nil {
		return false, utils.WrapError(err, "userUsecase.userRepository.VerifyUser : "), msg
	}

	dataEmail := model.EmailBody{
		BodyTemp: emailType,
		UserCode: user.Code,
	}
	userUsecase.emailUsecase.SendEmail(dataEmail)

	return result, err, msg
}

func (userUsecase *userUsecase) VerifyEmail(users shape.Users) (string, int, error) {
	var email string

	dataUser := model.Users{
		Email: users.Email,
		VerifiedCode: sql.NullString{
			String: users.Verified_Code,
		},
	}
	count, err := userUsecase.userRepository.CheckVerifyCodeUser(dataUser)
	if err != nil {
		return email, count, utils.WrapError(err, "userUsecase.userRepository.CheckVerifyCodeUser : ")
	}

	if count == 0 {
		return email, count, err
	}

	user, err := userUsecase.userRepository.GetEmailByVerifyCode(dataUser.VerifiedCode.String)
	if err != nil {
		return email, count, utils.WrapError(err, "userUsecase.userRepository.GetEmailByVerifyCode : ")
	}

	email = user.Email
	return email, count, err
}

func (userUsecase *userUsecase) GetUserInfo(email string) (shape.UserInfo, error) {
	user, err := userUsecase.userRepository.GetUserInfo(email)
	if err != nil {
		return shape.UserInfo{}, utils.WrapError(err, "userUsecase.userRepository.GetUserInfo : ")
	}

	if user == (model.UserInfo{}) {
		return shape.UserInfo{}, nil
	}

	userResult := shape.UserInfo{
		ID:            user.ID,
		Code:          user.Code,
		Role_Code:     user.RoleCode,
		Role:          user.Role,
		Email:         user.Email,
		Full_Name:     user.FullName.String,
		Phone:         int(user.Phone.Int64),
		Profession:    user.Profession.String,
		Gender:        user.Gender.String,
		Age:           int(user.Age.Int64),
		Birth:         utils.HandleNullableDate(user.Birth.Time),
		Province:      user.Province.String,
		City:          user.City.String,
		Address:       user.Address.String,
		Image_Profile: user.ImageProfile.String,
		Is_Verified:   user.IsVerified,
		Is_Active:     user.IsActive,
		Created_By:    user.CreatedBy,
		Created_Date:  user.CreatedDate,
		Modified_By:   user.ModifiedBy.String,
		Modified_Date: user.ModifiedDate.Time,
	}

	return userResult, err
}

func (userUsecase *userUsecase) CheckEmail(email string) (shape.UserInfo, error) {
	var emailType string = "Account Verification"

	user, err := userUsecase.userRepository.GetUserInfo(email)
	if err != nil {
		return shape.UserInfo{}, utils.WrapError(err, "userUsecase.userRepository.GetUserInfo : ")
	}

	if user == (model.UserInfo{}) || !user.IsVerified {
		return shape.UserInfo{}, nil
	}

	userResult := shape.UserInfo{
		ID:            user.ID,
		Code:          user.Code,
		Role_Code:     user.RoleCode,
		Role:          user.Role,
		Email:         user.Email,
		Full_Name:     user.FullName.String,
		Phone:         int(user.Phone.Int64),
		Profession:    user.Profession.String,
		Gender:        user.Gender.String,
		Age:           int(user.Age.Int64),
		Birth:         utils.HandleNullableDate(user.Birth.Time),
		Province:      user.Province.String,
		City:          user.City.String,
		Address:       user.Address.String,
		Image_Profile: user.ImageProfile.String,
		Is_Verified:   user.IsVerified,
		Is_Active:     user.IsActive,
		Created_By:    user.CreatedBy,
		Created_Date:  user.CreatedDate,
		Modified_By:   user.ModifiedBy.String,
		Modified_Date: user.ModifiedDate.Time,
	}

	dataUser := model.Users{
		Email: email,
		VerifiedCode: sql.NullString{
			String: utils.
				GenerateVerifyCode(email),
		},
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}

	firstLoop := true
	for firstLoop {
		count, err := userUsecase.userRepository.CheckVerifyCodeUser(dataUser)
		if err != nil {
			return shape.UserInfo{}, utils.WrapError(err, "userUsecase.userRepository.CheckVerifyCodeUser : ")
		}

		if count == 0 {
			firstLoop = false
		} else {
			dataUser.VerifiedCode.String = utils.GenerateVerifyCode(email)
		}
	}

	userAuth, result, err := userUsecase.userRepository.ResetVerificationCode(dataUser)
	if err != nil {
		return shape.UserInfo{}, utils.WrapError(err, "userUsecase.userRepository.ResetVerificationCode : ")
	}

	if result && userAuth.ID != 0 {
		dataEmail := model.EmailBody{
			UserCode:         userAuth.Code,
			BodyTemp:         emailType,
			VerificationCode: dataUser.VerifiedCode.String,
		}
		userUsecase.emailUsecase.SendEmail(dataEmail)
	}
	return userResult, err
}
