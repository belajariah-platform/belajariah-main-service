package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
)

type exerciseReadingUsecase struct {
	exerciseReadingRepository repository.ExerciseReadingRepository
}

type ExerciseReadingUsecase interface {
	GetAllExerciseReading(query model.Query) ([]shape.ExerciseReading, int, error)
}

func InitExerciseReadingUsecase(exerciseReadingRepository repository.ExerciseReadingRepository) ExerciseReadingUsecase {
	return &exerciseReadingUsecase{
		exerciseReadingRepository,
	}
}

func (exerciseReadingUsecase *exerciseReadingUsecase) GetAllExerciseReading(query model.Query) ([]shape.ExerciseReading, int, error) {
	var filterQuery string
	var exerciseReadings []model.ExerciseReading
	var exerciseReadingResult []shape.ExerciseReading

	filterQuery = utils.GetFilterHandler(query.Filters)

	exerciseReadings, err := exerciseReadingUsecase.exerciseReadingRepository.GetAllExerciseReading(query.Skip, query.Take, filterQuery)
	count, errCount := exerciseReadingUsecase.exerciseReadingRepository.GetAllExerciseReadingCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range exerciseReadings {
			exerciseReadingResult = append(exerciseReadingResult, shape.ExerciseReading{
				ID:            value.ID,
				Code:          value.Code,
				Title_Code:    value.TitleCode,
				Surat_Code:    value.SuratCode,
				Ayat_Start:    value.AyatStart,
				Ayat_End:      value.AyatEnd,
				Is_Active:     value.IsActive,
				Created_By:    value.CreatedBy,
				Created_Date:  value.CreatedDate,
				Modified_By:   value.ModifiedBy.String,
				Modified_Date: value.ModifiedDate.Time,
				Is_Deleted:    value.IsDeleted,
			})
		}
	}
	exerciseReadingEmpty := make([]shape.ExerciseReading, 0)
	if len(exerciseReadingResult) == 0 {
		return exerciseReadingEmpty, count, err
	}
	return exerciseReadingResult, count, err
}
