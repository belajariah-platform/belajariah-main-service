package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
)

type learningUsecase struct {
	learningRepository repository.LearningRepository
}

type LearningUsecase interface {
	GetAllLearning(query model.Query) ([]shape.Learning, int, error)
}

func InitLearningUsecase(learningRepository repository.LearningRepository) LearningUsecase {
	return &learningUsecase{
		learningRepository,
	}
}

func (learningUsecase *learningUsecase) GetAllLearning(query model.Query) ([]shape.Learning, int, error) {
	var filterQuery string
	var learnings []model.Learning
	var learningResult []shape.Learning

	filterQuery = utils.GetFilterHandler(query.Filters)

	learnings, err := learningUsecase.learningRepository.GetAllLearning(query.Skip, query.Take, filterQuery)
	count, errCount := learningUsecase.learningRepository.GetAllSubLearningCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range learnings {

			var subLearning []model.SubLearning
			var subLearningResult []shape.SubLearning
			subLearning, err := learningUsecase.learningRepository.GetAllSubLearning(value.Code)
			if err == nil {
				for _, val := range subLearning {
					subLearningResult = append(subLearningResult, shape.SubLearning{
						ID:             val.ID,
						Code:           val.Code,
						Title_Code:     val.TitleCode,
						Sub_Title:      val.SubTitle.String,
						Video_Duration: val.VideoDuration.Float64,
						Video:          val.Video.String,
						Document:       val.Document.String,
						Exercise_Image: val.ExerciseImage.String,
						Sequence:       int(val.Sequence.Int64),
						Is_Active:      val.IsActive,
						Created_By:     val.CreatedBy,
						Created_Date:   val.CreatedDate,
						Modified_By:    val.ModifiedBy.String,
						Modified_Date:  val.ModifiedDate.Time,
						Deleted_By:     val.DeletedBy.String,
						Deleted_Date:   val.DeletedDate.Time,
					})
				}
			}
			learningResult = append(learningResult, shape.Learning{
				ID:            value.ID,
				Code:          value.Code,
				Class_Code:    value.ClassCode,
				Title:         value.Title,
				Document:      value.Document.String,
				Sequence:      int(value.Sequence.Int64),
				SubTitles:     subLearningResult,
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
	learningEmpty := make([]shape.Learning, 0)
	if len(learningResult) == 0 {
		return learningEmpty, count, err
	}
	return learningResult, count, err
}
