package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
)

type learningUsecase struct {
	learningRepository        repository.LearningRepository
	exerciseReadingRepository repository.ExerciseReadingRepository
}

type LearningUsecase interface {
	GetAllLearning(query model.Query) ([]shape.Learning, int, error)
}

func InitLearningUsecase(learningRepository repository.LearningRepository, exerciseReadingRepository repository.ExerciseReadingRepository) LearningUsecase {
	return &learningUsecase{
		learningRepository,
		exerciseReadingRepository,
	}
}

func (learningUsecase *learningUsecase) GetAllLearning(query model.Query) ([]shape.Learning, int, error) {
	var count int
	var isDone bool
	var filterQuery string
	var learnings []model.Learning
	var learningResult []shape.Learning

	filterQuery = utils.GetFilterHandler(query.Filters)

	learnings, err := learningUsecase.learningRepository.GetAllLearning(query.Skip, query.Take, filterQuery)

	if err == nil {
		for _, value := range learnings {

			var subLearning []model.SubLearning
			var subLearningResult []shape.SubLearning
			subLearning, err := learningUsecase.learningRepository.GetAllSubLearning(value.Code)
			if err == nil {
				for _, sublearn := range subLearning {
					subLearningResult = append(subLearningResult, shape.SubLearning{
						ID:             sublearn.ID,
						Code:           sublearn.Code,
						Title_Code:     sublearn.TitleCode,
						Sub_Title:      sublearn.SubTitle.String,
						Video_Duration: sublearn.VideoDuration.Float64,
						Video:          sublearn.Video.String,
						Document:       sublearn.Document.String,
						Poster:         sublearn.Poster.String,
						Sequence:       int(sublearn.Sequence.Int64),
						Is_Done:        isDone,
						Is_Exercise:    sublearn.IsExercise.Bool,
						Is_Active:      sublearn.IsActive,
						Created_By:     sublearn.CreatedBy,
						Created_Date:   sublearn.CreatedDate,
						Modified_By:    sublearn.ModifiedBy.String,
						Modified_Date:  sublearn.ModifiedDate.Time,
						Deleted_By:     sublearn.DeletedBy.String,
						Deleted_Date:   sublearn.DeletedDate.Time,
					})
				}
				count = count + len(subLearning)
			}
			var exerciseResult shape.ExerciseReading
			if value.IsExercise.Bool {
				exercise, err := learningUsecase.exerciseReadingRepository.GetExerciseReading(value.Code)
				if err == nil {
					exerciseResult = shape.ExerciseReading{
						ID:            exercise.ID,
						Code:          exercise.Code,
						Title_Code:    exercise.TitleCode,
						Surat_Code:    exercise.SuratCode,
						Ayat_Start:    exercise.AyatStart,
						Ayat_End:      exercise.AyatEnd,
						Is_Active:     exercise.IsActive,
						Created_By:    exercise.CreatedBy,
						Created_Date:  exercise.CreatedDate,
						Modified_By:   exercise.ModifiedBy.String,
						Modified_Date: exercise.ModifiedDate.Time,
						Deleted_By:    exercise.DeletedBy.String,
						Deleted_Date:  exercise.DeletedDate.Time,
					}
				}
			}
			learningResult = append(learningResult, shape.Learning{
				ID:                 value.ID,
				Code:               value.Code,
				Class_Code:         value.ClassCode,
				Title:              value.Title,
				Document_Path:      value.DocumentPath.String,
				Document_Name:      value.DocumentName.String,
				Sequence:           int(value.Sequence.Int64),
				SubTitles:          subLearningResult,
				Exercises:          exerciseResult,
				Is_Exercise:        value.IsExercise.Bool,
				Is_Direct_Learning: value.IsDirectLearning,
				Is_Active:          value.IsActive,
				Created_By:         value.CreatedBy,
				Created_Date:       value.CreatedDate,
				Modified_By:        value.ModifiedBy.String,
				Modified_Date:      value.ModifiedDate.Time,
				Deleted_By:         value.DeletedBy.String,
				Deleted_Date:       value.DeletedDate.Time,
			})
		}
	}
	learningEmpty := make([]shape.Learning, 0)
	if len(learningResult) == 0 {
		return learningEmpty, count, err
	}

	return learningResult, count, err
}
