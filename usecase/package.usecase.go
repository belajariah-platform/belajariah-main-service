package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
)

type packageUsecase struct {
	packageRepository repository.PackageRepository
}

type PackageUsecase interface {
	GetAllPackage(query model.Query) ([]shape.Package, int, error)
	GetAllBenefit(query model.Query) ([]shape.Benefit, error)
}

func InitPackageUsecase(packageRepository repository.PackageRepository) PackageUsecase {
	return &packageUsecase{
		packageRepository,
	}
}

func (packageUsecase *packageUsecase) GetAllPackage(query model.Query) ([]shape.Package, int, error) {
	var filterQuery string
	var packages []model.Package
	var packageResult []shape.Package

	filterQuery = utils.GetFilterHandler(query.Filters)

	packages, err := packageUsecase.packageRepository.GetAllPackage(query.Skip, query.Take, filterQuery)
	count, errCount := packageUsecase.packageRepository.GetAllPackageCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range packages {
			packageResult = append(packageResult, shape.Package{
				ID:             value.ID,
				Code:           value.Code,
				Class_Code:     value.ClassCode,
				Type:           value.Type,
				Price_Package:  value.PricePackage,
				Price_Discount: value.PriceDiscount.String,
				Duration:       value.Duration,
				Consultation:   int(value.Consultation.Int64),
				Webinar:        int(value.Webinar.Int64),
				Description:    value.Description.String,
				Is_Active:      value.IsActive,
				Created_By:     value.CreatedBy,
				Created_Date:   value.CreatedDate,
				Modified_By:    value.ModifiedBy.String,
				Modified_Date:  value.ModifiedDate.Time,
				Is_Deleted:     value.IsDeleted,
			})
		}
	}
	packageEmpty := make([]shape.Package, 0)
	if len(packageResult) == 0 {
		return packageEmpty, count, err
	}
	return packageResult, count, err
}

func (packageUsecase *packageUsecase) GetAllBenefit(query model.Query) ([]shape.Benefit, error) {
	var filterQuery string
	var benefits []model.Benefit
	var benefitResult []shape.Benefit

	filterQuery = utils.GetFilterHandler(query.Filters)

	benefits, err := packageUsecase.packageRepository.GetAllBenefit(query.Skip, query.Take, filterQuery)

	if err == nil {
		for _, value := range benefits {
			benefitResult = append(benefitResult, shape.Benefit{
				ID:            value.ID,
				Code:          value.Code,
				Class_Code:    value.ClassCode,
				Description:   value.Description,
				Icon_Benefit:  value.IconBenefit.String,
				Sequence:      value.Sequence,
				Is_Active:     value.IsActive,
				Created_By:    value.CreatedBy,
				Created_Date:  value.CreatedDate,
				Modified_By:   value.ModifiedBy.String,
				Modified_Date: value.ModifiedDate.Time,
				Is_Deleted:    value.IsDeleted,
			})
		}
	}
	benefitEmpty := make([]shape.Benefit, 0)
	if len(benefitResult) == 0 {
		return benefitEmpty, err
	}
	return benefitResult, err
}
