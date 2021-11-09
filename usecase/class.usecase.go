package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
	"fmt"
)

type classUsecase struct {
	classRepository   repository.ClassRepository
	packageRepository repository.PackageRepository
}

type ClassUsecase interface {
	GetAllClass(query model.Query) ([]shape.Class, int, error)
}

func InitClassUsecase(classRepository repository.ClassRepository, packageRepository repository.PackageRepository) ClassUsecase {
	return &classUsecase{
		classRepository,
		packageRepository,
	}
}

func (classUsecase *classUsecase) GetAllClass(query model.Query) ([]shape.Class, int, error) {
	var classes []model.Class
	var classResult []shape.Class
	var filterQuery, priceStart, priceStartDiscount, priceEnd, priceEndDiscount string

	filterQuery = utils.GetFilterHandler(query.Filters)

	classes, err := classUsecase.classRepository.GetAllClass(query.Skip, query.Take, filterQuery)
	count, errCount := classUsecase.classRepository.GetAllClassCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range classes {
			var packages []model.Package

			filterQuery = fmt.Sprintf(`AND class_code = '%s'`, value.Code)
			packages, err := classUsecase.packageRepository.GetAllPackage(query.Skip, query.Take, filterQuery)
			if err == nil {
				for index, values := range packages {
					if index == 0 {
						priceStart = values.PricePackage
						priceStartDiscount = values.PriceDiscount.String
					} else if index+1 == len(packages) {
						priceEnd = values.PricePackage
						priceEndDiscount = values.PriceDiscount.String
					}
				}
			}

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
				Class_Document:         value.ClassDocument.String,
				Class_Rating:           value.ClassRating,
				Total_Review:           value.TotalReview,
				Total_Video:            value.TotalVideo,
				Total_Video_Duration:   value.TotalVideoDuration,
				Instructor_Name:        value.InstructorName.String,
				Instructor_Description: value.InstructorDescription.String,
				Instructor_Biografi:    value.InstructorBiografi.String,
				Instructor_Image:       value.InstructorImage.String,
				Is_Direct:              value.IsDirect,
				Is_Active:              value.IsActive,
				Created_By:             value.CreatedBy,
				Created_Date:           value.CreatedDate,
				Modified_By:            value.ModifiedBy.String,
				Modified_Date:          value.ModifiedDate.Time,
				Is_Deleted:             value.IsDeleted,
				Price_Start:            priceStart,
				Price_Start_Discount:   priceStartDiscount,
				Price_End_Discount:     priceEndDiscount,
				Price_End:              priceEnd,
			})
		}
	}
	classEmpty := make([]shape.Class, 0)
	if len(classResult) == 0 {
		return classEmpty, count, err
	}
	return classResult, count, err
}
