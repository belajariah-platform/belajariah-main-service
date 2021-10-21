package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
)

type exerciseUsecase struct {
	exerciseRepository repository.ExerciseRepository
}

type ExerciseUsecase interface {
	GetAllExercise(query model.Query) ([]shape.Exercise, int, error)
}

func InitExerciseUsecase(exerciseRepository repository.ExerciseRepository) ExerciseUsecase {
	return &exerciseUsecase{
		exerciseRepository,
	}
}

func (exerciseUsecase *exerciseUsecase) GetAllExercise(query model.Query) ([]shape.Exercise, int, error) {
	var filterQuery string
	var exercises []model.Exercise
	var exerciseResult []shape.Exercise

	filterQuery = utils.GetFilterHandler(query.Filters)

	exercises, err := exerciseUsecase.exerciseRepository.GetAllExercise(query.Skip, query.Take, filterQuery)
	count, errCount := exerciseUsecase.exerciseRepository.GetAllExerciseCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range exercises {
			exerciseResult = append(exerciseResult, shape.Exercise{
				ID:             value.ID,
				Code:           value.Code,
				Subtitle_Code:  value.SubtitleCode,
				Image_Exercise: value.ImageExercise.String,
				Is_Active:      value.IsActive,
				Created_By:     value.CreatedBy,
				Created_Date:   value.CreatedDate,
				Modified_By:    value.ModifiedBy.String,
				Modified_Date:  value.ModifiedDate.Time,
				Is_Deleted:     value.IsDeleted,
			})
		}
	}
	exerciseEmpty := make([]shape.Exercise, 0)
	if len(exerciseResult) == 0 {
		return exerciseEmpty, count, err
	}
	return exerciseResult, count, err
}
