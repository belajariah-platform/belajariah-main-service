package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

type ratingUsecase struct {
	ratingRepository repository.RatingRepository
}

type RatingUsecase interface {
	GetAllRatingClass(query model.Query) ([]shape.Rating, int, error)
	GiveRatingClass(rating shape.RatingPost, email string) (bool, error)
	GiveRatingMentor(rating shape.RatingPost, email string) (bool, error)

	GetAllRatingClassQuran(r model.RatingQuranRequest) ([]model.RatingQuran, int, error)
	GiveRatingClassQuran(ctx *gin.Context, r model.RatingQuranRequest) (bool, error)
}

func InitRatingUsecase(ratingRepository repository.RatingRepository) RatingUsecase {
	return &ratingUsecase{
		ratingRepository,
	}
}

func (ratingUsecase *ratingUsecase) GetAllRatingClass(query model.Query) ([]shape.Rating, int, error) {
	var filterQuery string
	var ratings []model.Rating
	var ratingResult []shape.Rating
	ratingEmpty := make([]shape.Rating, 0)

	filterQuery = utils.GetFilterHandler(query.Filters)

	ratings, err := ratingUsecase.ratingRepository.GetAllRatingClass(query.Skip, query.Take, filterQuery)
	if err != nil {
		return ratingEmpty, 0, utils.WrapError(err, "ratingUsecase.ratingRepository.GetAllRatingClass : ")
	}

	count, errCount := ratingUsecase.ratingRepository.GetAllRatingClassCount(filterQuery)
	if err != nil {
		return ratingEmpty, 0, utils.WrapError(err, "ratingUsecase.userRepository.GetAllRatingClassCount : ")
	}

	if err == nil && errCount == nil {
		for _, value := range ratings {
			ratingResult = append(ratingResult, shape.Rating{
				ID:            value.ID,
				Code:          value.Code,
				Class_Code:    value.ClassCode,
				Class_Name:    value.ClassName,
				Class_Initial: value.ClassInitial.String,
				Rating:        value.Rating,
				Comment:       value.Comment.String,
				User_Code:     value.UserCode,
				User_Name:     value.UserName,
				Is_Active:     value.IsActive,
				Created_By:    value.CreatedBy,
				Created_Date:  value.CreatedDate,
				Modified_By:   value.ModifiedBy.String,
				Modified_Date: value.ModifiedDate.Time,
				Is_Deleted:    value.IsDeleted,
			})
		}
	}

	if len(ratingResult) == 0 {
		return ratingEmpty, count, err
	}
	return ratingResult, count, err
}

func (ratingUsecase *ratingUsecase) GiveRatingClass(rating shape.RatingPost, email string) (bool, error) {
	dataRating := model.RatingPost{
		ClassCode: rating.Class_Code,
		UserCode:  rating.User_Code,
		Rating:    rating.Rating,
		Comment: sql.NullString{
			String: rating.Comment,
		},
		CreatedBy:   email,
		CreatedDate: time.Now(),
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}
	result, err := ratingUsecase.ratingRepository.GiveRatingClass(dataRating)
	if err != nil {
		return false, utils.WrapError(err, "ratingUsecase.userRepository.GiveRatingClass : ")
	}

	return result, err
}

func (ratingUsecase *ratingUsecase) GiveRatingMentor(rating shape.RatingPost, email string) (bool, error) {
	dataRating := model.RatingPost{
		MentorCode: rating.Mentor_Code,
		UserCode:   rating.User_Code,
		Rating:     rating.Rating,
		Comment: sql.NullString{
			String: rating.Comment,
		},
		CreatedBy:   email,
		CreatedDate: time.Now(),
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}
	result, err := ratingUsecase.ratingRepository.GiveRatingMentor(dataRating)
	if err != nil {
		return false, utils.WrapError(err, "ratingUsecase.userRepository.GiveRatingMentor : ")
	}

	return result, err
}

func (u *ratingUsecase) GetAllRatingClassQuran(r model.RatingQuranRequest) ([]model.RatingQuran, int, error) {
	var orderDefault = "ORDER BY code asc"
	var filterDefault = "is_deleted = false and is_active = true"

	filterFinal := utils.GetFilterOrderHandler(filterDefault, orderDefault, r.Query)
	result, err := u.ratingRepository.GetAllClassRatingQuran(filterFinal)
	if err != nil {
		return nil, 0, utils.WrapError(err, "ratingUsecase.GetAllClassRatingQuran")
	}

	ratingEmpty := make([]model.RatingQuran, 0)
	if len(*result) == 0 {
		return ratingEmpty, 0, err
	}

	return *result, len(*result), nil
}

func (u *ratingUsecase) GiveRatingClassQuran(ctx *gin.Context, r model.RatingQuranRequest) (bool, error) {
	email := ctx.Request.Header.Get("email")

	r.Data.ModifiedBy.String = email
	r.Data.ModifiedDate.Time = time.Now()

	result, err := u.ratingRepository.GiveRatingClassQuran(r.Data)
	if err != nil {
		return false, utils.WrapError(err, "ratingUsecase.GiveRatingClassQuran")
	}

	return result, nil
}
