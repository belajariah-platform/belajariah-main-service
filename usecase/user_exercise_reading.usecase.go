package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"database/sql"
	"fmt"
	"time"
)

type userExerciseReadingUsecase struct {
	userExerciseReadingRepository repository.UserExerciseReadingRepository
}

type UserExerciseReadingUsecase interface {
	InserteUserExerciseReading(userExercise shape.UserExerciseReading, email string) (bool, error)
}

func InitUserExerciseReadingUsecase(userExerciseReadingRepository repository.UserExerciseReadingRepository) UserExerciseReadingUsecase {
	return &userExerciseReadingUsecase{
		userExerciseReadingRepository,
	}
}

func (userExerciseReadingUsecase *userExerciseReadingUsecase) InserteUserExerciseReading(userExercise shape.UserExerciseReading, email string) (bool, error) {
	var result bool

	dataUserExercise := model.UserExerciseReading{
		UserCode:      userExercise.User_Code,
		ClassCode:     userExercise.Class_Code,
		RecordingCode: userExercise.Recording_Code,
		Duration:      userExercise.Duration,
		ExpiredDate:   userExercise.Expired_Date,
		CreatedBy:     email,
		CreatedDate:   time.Now(),
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}

	filter := fmt.Sprintf(`user_code=%d AND class_code='%s' AND expired_date='%s'`,
		userExercise.User_Code,
		userExercise.Class_Code,
		userExercise.Expired_Date,
	)

	count, err := userExerciseReadingUsecase.userExerciseReadingRepository.GetAllUserExerciseReadingCount(filter)
	if count >= userExercise.Max_Recording {
		return result, err
	}
	result, err = userExerciseReadingUsecase.userExerciseReadingRepository.InsertUserExerciseReading(dataUserExercise)
	return result, err
}
