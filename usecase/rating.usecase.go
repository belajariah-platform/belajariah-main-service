package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
	"database/sql"
	"time"
)

type ratingUsecase struct {
	ratingRepository repository.RatingRepository
}

type RatingUsecase interface {
	GetAllRatingClass(query model.Query) ([]shape.Rating, int, error)
	GiveRatingClass(rating shape.RatingPost, email string) (bool, error)
	GiveRatingMentor(rating shape.RatingPost, email string) (bool, error)
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

	filterQuery = utils.GetFilterHandler(query.Filters)

	ratings, err := ratingUsecase.ratingRepository.GetAllRatingClass(query.Skip, query.Take, filterQuery)
	count, errCount := ratingUsecase.ratingRepository.GetAllRatingClassCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range ratings {
			ratingResult = append(ratingResult, shape.Rating{
				ID:            value.ID,
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
				Deleted_By:    value.DeletedBy.String,
				Deleted_Date:  value.DeletedDate.Time,
			})
		}
	}
	ratingEmpty := make([]shape.Rating, 0)
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
	return result, err
}
