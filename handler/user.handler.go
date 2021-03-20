package handler

import (
	"belajariah-main-service/shape"
	"belajariah-main-service/usecase"
	"belajariah-main-service/utils"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
}

type UserHandler interface {
	GetUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	VerifyUser(ctx *gin.Context)
	RegisterUser(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
	VerifyPasswordUser(ctx *gin.Context)
	ResetVerificationUser(ctx *gin.Context)
}

func InitUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase,
	}
}

func (userHandler *userHandler) LoginUser(ctx *gin.Context) {
	var err error
	var result bool
	var token, message string
	var loginJSON shape.Users
	var userInfo shape.UserInfo

	if err := ctx.ShouldBindJSON(&loginJSON); err == nil {
		if loginJSON.Email != "" && loginJSON.Password != "" {
			userInfo, result, err, message = userHandler.userUsecase.LoginUser(loginJSON)
			if err == nil && userInfo.Is_Verified {
				token, err = getAuthToken(loginJSON.Email)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"Error generate token ": err,
					})
				}
			}
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Email dan password tidak boleh kosong",
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   err,
		"token":   token,
		"result":  result,
		"data":    userInfo,
		"message": message,
	})
}

func (userHandler *userHandler) ResetVerificationUser(ctx *gin.Context) {
	var users shape.Users

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result, err := userHandler.userUsecase.ResetVerificationUser(users)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": result,
			"error":  "",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": result,
			"error":  err.Error(),
		})
	}
}

func (userHandler *userHandler) VerifyUser(ctx *gin.Context) {
	var users shape.Users

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result, err, msg := userHandler.userUsecase.VerifyUser(users)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": result,
			"error":  "",
			"mesage": msg,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": result,
			"error":  err.Error(),
		})
	}
}

func (userHandler *userHandler) VerifyPasswordUser(ctx *gin.Context) {
	var users shape.Users

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result, err, message := userHandler.userUsecase.VerifyPasswordUser(users)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result":  result,
			"error":   "",
			"message": message,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result":  result,
			"error":   err.Error(),
			"message": message,
		})
	}
}

func (userHandler *userHandler) GetUser(ctx *gin.Context) {
	email := ctx.Param("email")

	result, err := userHandler.userUsecase.GetUserInfo(email)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": result,
			"error":  "",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": result,
			"error":  err.Error(),
		})
	}
}

func (userHandler *userHandler) RegisterUser(ctx *gin.Context) {
	var users shape.Users

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result, err, msg := userHandler.userUsecase.RegisterUser(users)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": result,
			"error":  "",
			"mesage": msg,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": result,
			"error":  err.Error(),
		})
	}
}

func (userHandler *userHandler) UpdateProfile(ctx *gin.Context) {
	var users shape.UsersPost
	var email string
	for _, value := range ctx.Request.Header["Email"] {
		email = value
		break
	}
	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := userHandler.userUsecase.UpdateProfileUser(users, email)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": result,
			"error":  "",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": result,
			"error":  err.Error(),
		})
	}
}

func (userHandler *userHandler) ChangePassword(ctx *gin.Context) {
	var users shape.Users

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result, err := userHandler.userUsecase.ChangePasswordUser(users)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": result,
			"error":  "",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": result,
			"error":  err.Error(),
		})
	}
}

func getAuthToken(email string) (string, error) {
	var err error
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		utils.PushLogf("Error generate token :", err)
		return "", err
	}
	return token, nil
}
