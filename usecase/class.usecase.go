package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
)

type classUsecase struct {
	classRepository repository.ClassRepository
}

type ClassUsecase interface {
	GetAllClass(query model.Query) ([]shape.Class, int, error)
}

func InitClassUsecase(classRepository repository.ClassRepository) ClassUsecase {
	return &classUsecase{
		classRepository,
	}
}

func (classUsecase *classUsecase) GetAllClass(query model.Query) ([]shape.Class, int, error) {
	var filterQuery string
	var classes []model.Class
	var classResult []shape.Class

	filterQuery = utils.GetFilterHandler(query.Filters)

	classes, err := classUsecase.classRepository.GetAllClass(query.Skip, query.Take, filterQuery)
	count, errCount := classUsecase.classRepository.GetAllClassCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range classes {
			classResult = append(classResult, shape.Class{
				ID:                     value.ID,
				Code:                   value.Code,
				Class_Category_Code:    value.ClassCategoryCode,
				Class_Category:         value.ClassCategory,
				Class_Name:             value.ClassName,
				Class_Initial:          value.ClassInitial.String,
				Class_Description:      value.ClassDescription.String,
				Class_Image:            value.ClassImage.String,
				Class_Video:            value.ClassVideo.String,
				Class_Rating:           value.ClassRating,
				Total_Review:           value.TotalReview,
				Total_Video:            value.TotalVideo,
				Total_Video_Duration:   value.TotalVideoDuration,
				Instructor_Name:        value.InstructorName,
				Instructor_Description: value.InstructorDescription.String,
				Instructor_Biografi:    value.InstructorBiografi.String,
				Instructor_Image:       value.InstructorImage.String,
				Is_Active:              value.IsActive,
				Created_By:             value.CreatedBy,
				Created_Date:           value.CreatedDate,
				Modified_By:            value.ModifiedBy.String,
				Modified_Date:          value.ModifiedDate.Time,
				Deleted_By:             value.DeletedBy.String,
				Deleted_Date:           value.DeletedDate.Time,
			})
		}
	}
	classEmpty := make([]shape.Class, 0)
	if len(classResult) == 0 {
		return classEmpty, count, err
	}
	return classResult, count, err
}
