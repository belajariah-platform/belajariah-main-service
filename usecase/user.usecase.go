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
	userRepository repository.UserRepository
}

type UserUsecase interface {
	LoginUser(users shape.Users) (shape.UserInfo, error, string)
	RegisterUser(users shape.Users) (bool, error, string)
	VerifyUser(users shape.Users) (bool, error, string)
	ChangePasswordUser(users shape.Users) (bool, error)
	GetUserInfo(email string) (shape.UserInfo, error)
}

func InitUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository,
	}
}

func (userUsecase *userUsecase) LoginUser(users shape.Users) (shape.UserInfo, error, string) {
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
			return shape.UserInfo{}, err, msg
		}
	}
	user, err := userUsecase.GetUserInfo(dataUser.Email)
	if !user.Is_Verified {
		msg = fmt.Sprintf(`Akun kamu belum terverifikasi`)
		return user, err, msg
	}
	return user, err, msg
}

func (userUsecase *userUsecase) RegisterUser(users shape.Users) (bool, error, string) {
	var user model.UserInfo
	var result bool
	var msg string
	var err error

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
	if user.Email == dataUser.Email {
		msg = fmt.Sprintf(`Email '%s' sudah ada`, dataUser.Email)
		return result, err, msg
	}
	result, err = userUsecase.userRepository.RegisterUser(dataUser)
	return result, err, msg
}

func (userUsecase *userUsecase) VerifyUser(users shape.Users) (bool, error, string) {
	var result bool
	var msg string
	var err error

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
	result, err = userUsecase.userRepository.VerifyUser(dataUser)
	return result, err, msg
}

func (userUsecase *userUsecase) ChangePasswordUser(users shape.Users) (bool, error) {
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
	result, err = userUsecase.userRepository.ChangePasswordUser(dataUser)
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
