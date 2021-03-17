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
	LoginUser(users shape.Users) (shape.UserInfo, bool, error, string)
	UpdateProfileUser(users shape.UsersPost, email string) (bool, error)
	ResetVerificationUser(users shape.Users) (bool, error)
	RegisterUser(users shape.Users) (bool, error, string)
	VerifyPasswordUser(users shape.Users) (bool, error)
	VerifyUser(users shape.Users) (bool, error, string)
	ChangePasswordUser(users shape.Users) (bool, error)
	GetUserInfo(email string) (shape.UserInfo, error)
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
	if err == nil {
		isPassword := utils.CheckPasswordHash(dataUser.Password, userLogin.Password)
		if userLogin == (model.Users{}) || !isPassword {
			msg = fmt.Sprintf(`Email dan kata sandi salah`)
			return shape.UserInfo{}, false, err, msg
		}
	}
	user, err := userUsecase.GetUserInfo(dataUser.Email)
	if !user.Is_Verified {
		msg = fmt.Sprintf(`Akun kamu belum terverifikasi`)
		return shape.UserInfo{}, false, err, msg
	}
	return user, true, err, msg
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
	if err == nil && userAuth.ID != 0 {
		dataEmail := model.EmailBody{
			UserCode:         userAuth.ID,
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
	var err error
	var msg string
	var userID int
	var result bool

	hashPassword, err := utils.GenerateHashPassword(users.Password)
	if err != nil {
		utils.PushLogf("error :", err)
		return result, err, msg
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
	if user.Email == dataUser.Email && user.IsVerified {
		msg = fmt.Sprintf(`Email '%s' sudah ada`, dataUser.Email)
		return result, err, msg
	} else if user.Email == dataUser.Email && !user.IsVerified {
		msg = fmt.Sprintf(`Silahkan verifikasi email anda`)
		return result, err, msg
	}
	userID, result, err = userUsecase.userRepository.RegisterUser(dataUser)
	if err == nil && userID != 0 {
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
	fmt.Println(users)
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
	return result, err
}

func (userUsecase *userUsecase) VerifyPasswordUser(users shape.Users) (bool, error) {
	var err error
	var result bool
	var user model.Users

	dataUser := model.Users{
		Email: users.Email,
		VerifiedCode: sql.NullString{
			String: users.Verified_Code,
		},
	}

	user, result, err = userUsecase.userRepository.VerifyUser(dataUser)
	utils.PushLogf(user.Email)
	return result, err
}

func (userUsecase *userUsecase) VerifyUser(users shape.Users) (bool, error, string) {
	var err error
	var msg string
	var result bool
	var user model.Users
	var emailType string = "Registration Success"

	dataUser := model.Users{
		Email: users.Email,
		VerifiedCode: sql.NullString{
			String: users.Verified_Code,
		},
	}
	count, err := userUsecase.userRepository.CheckVerifyCodeUser(dataUser)
	if count == 0 {
		msg = fmt.Sprintf(`Kode verifikasi salah`)
		return result, err, msg
	}
	user, result, err = userUsecase.userRepository.VerifyUser(dataUser)
	if err == nil {
		dataEmail := model.EmailBody{
			BodyTemp: emailType,
			UserCode: user.ID,
		}
		userUsecase.emailUsecase.SendEmail(dataEmail)
	}
	return result, err, msg
}

func (userUsecase *userUsecase) ChangePasswordUser(users shape.Users) (bool, error) {
	var emailType string = "Change Password"
	var user model.Users
	var result bool
	var err error

	hashPassword, err := utils.GenerateHashPassword(users.Password)
	if err != nil {
		utils.PushLogf("error :", err)
		return result, err
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

	user, result, err = userUsecase.userRepository.ChangePasswordUser(dataUser)
	if err == nil {
		dataEmail := model.EmailBody{
			UserEmail: dataUser.Email,
			BodyTemp:  emailType,
			UserCode:  user.ID,
		}
		userUsecase.emailUsecase.SendEmail(dataEmail)
	}
	return result, err
}

func (userUsecase *userUsecase) GetUserInfo(email string) (shape.UserInfo, error) {
	user, err := userUsecase.userRepository.GetUserInfo(email)
	if user == (model.UserInfo{}) {
		return shape.UserInfo{}, nil
	}
	userResult := shape.UserInfo{
		ID:             user.ID,
		Role_Code:      user.RoleCode,
		Role:           user.Role,
		Email:          user.Email,
		Full_Name:      user.FullName.String,
		Phone:          int(user.Phone.Int64),
		Profession:     user.Profession.String,
		Gender:         user.Gender.String,
		Age:            int(user.Age.Int64),
		Birth:          utils.HandleNullableDate(user.Birth.Time),
		Province:       user.Province.String,
		City:           user.City.String,
		Address:        user.Address.String,
		Image_Code:     user.ImageCode.String,
		Image_Filename: user.ImageFilename.String,
		Is_Verified:    user.IsVerified,
		Is_Active:      user.IsActive,
		Created_By:     user.CreatedBy,
		Created_Date:   user.CreatedDate,
		Modified_By:    user.ModifiedBy.String,
		Modified_Date:  user.ModifiedDate.Time,
	}
	return userResult, err
}
