package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
	"fmt"
)

type learningUsecase struct {
	learningRepository        repository.LearningRepository
	exerciseReadingRepository repository.ExerciseReadingRepository
}

type LearningUsecase interface {
	GetAllLearning(query model.Query) ([]shape.Learning, int, error)
	GetAllLearningQuran(r model.LearningQuranRequest) ([]model.LearningQuran, error)
}

func InitLearningUsecase(learningRepository repository.LearningRepository, exerciseReadingRepository repository.ExerciseReadingRepository) LearningUsecase {
	return &learningUsecase{
		learningRepository,
		exerciseReadingRepository,
	}
}

func (u *learningUsecase) GetAllLearning(query model.Query) ([]shape.Learning, int, error) {
	var count int
	var isDone bool
	var filterQuery string
	var learnings []model.Learning
	var learningResult []shape.Learning

	filterQuery = utils.GetFilterHandler(query.Filters)

	learnings, err := u.learningRepository.GetAllLearning(query.Skip, query.Take, filterQuery)

	if err == nil {
		for _, value := range learnings {
			var subLearning []model.SubLearning
			var subLearningResult []shape.SubLearning
			subLearning, err := u.learningRepository.GetAllSubLearning(value.Code)
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
						Sequence:       sublearn.Sequence,
						Is_Done:        isDone,
						Is_Exercise:    sublearn.IsExercise,
						Is_Active:      sublearn.IsActive,
						Created_By:     sublearn.CreatedBy,
						Created_Date:   sublearn.CreatedDate,
						Modified_By:    sublearn.ModifiedBy.String,
						Modified_Date:  sublearn.ModifiedDate.Time,
						Is_Deleted:     value.IsDeleted,
					})
				}
				count = count + len(subLearning)
			}

			var exerciseResult []shape.ExerciseReading
			if value.IsExercise {
				var exercises []model.ExerciseReading
				filter := fmt.Sprintf(`AND title_code='%s'`, value.Code)
				exercises, err := u.exerciseReadingRepository.GetAllExerciseReading(0, 100, filter)
				if err == nil {
					for _, exercise := range exercises {
						exerciseResult = append(exerciseResult, shape.ExerciseReading{
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
							Is_Deleted:    exercise.IsDeleted,
						})
					}
				}
			}

			sublearningEmpty := make([]shape.SubLearning, 0)
			if len(subLearningResult) == 0 {
				subLearningResult = sublearningEmpty
			}

			exerciseEmpty := make([]shape.ExerciseReading, 0)
			if len(subLearningResult) == 0 {
				exerciseResult = exerciseEmpty
			}

			learningResult = append(learningResult, shape.Learning{
				ID:                 value.ID,
				Code:               value.Code,
				Class_Code:         value.ClassCode,
				Title:              value.Title,
				Document_Path:      value.DocumentPath.String,
				Document_Name:      value.DocumentName.String,
				Sequence:           value.Sequence,
				Is_Exercise:        value.IsExercise,
				Is_Direct_Learning: value.IsDirectLearning,
				Is_Active:          value.IsActive,
				Created_By:         value.CreatedBy,
				Created_Date:       value.CreatedDate,
				Modified_By:        value.ModifiedBy.String,
				Modified_Date:      value.ModifiedDate.Time,
				Is_Deleted:         value.IsDeleted,
				Exercises:          exerciseResult,
				SubTitles:          subLearningResult,
			})
		}
	}

	learningEmpty := make([]shape.Learning, 0)
	if len(learningResult) == 0 {
		return learningEmpty, count, err
	}

	return learningResult, count, err
}

func (u *learningUsecase) GetAllLearningQuran(r model.LearningQuranRequest) ([]model.LearningQuran, error) {
	var orderDefault = "ORDER BY code asc"
	var filterDefault = "is_deleted = false and is_active = true"
	filterFinal := utils.GetFilterOrderHandler(filterDefault, orderDefault, r.Query)

	result, err := u.learningRepository.GetAllLearningQuran(filterFinal)
	if err != nil {
		return nil, utils.WrapError(err, "coachingProgramUsecase.GetAllLearningQuran")
	}

	learningEmpty := make([]model.LearningQuran, 0)
	if len(*result) == 0 {
		return learningEmpty, err
	}

	return *result, nil
}
