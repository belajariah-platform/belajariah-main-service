package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
)

type enumUsecase struct {
	enumRepository repository.EnumRepository
}

type EnumUsecase interface {
	GetAllEnum(query model.Query) ([]shape.Enum, int, error)
}

func InitEnumUsecase(enumRepository repository.EnumRepository) EnumUsecase {
	return &enumUsecase{
		enumRepository,
	}
}

func (enumUsecase *enumUsecase) GetAllEnum(query model.Query) ([]shape.Enum, int, error) {
	var filterQuery string
	var enums []model.Enum
	var enumResult []shape.Enum

	filterQuery = utils.GetFilterHandler(query.Filters)

	enums, err := enumUsecase.enumRepository.GetAllEnum(query.Skip, query.Take, filterQuery)
	count, errCount := enumUsecase.enumRepository.GetAllEnumCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range enums {
			enumResult = append(enumResult, shape.Enum{
				ID:            value.ID,
				Code:          value.Code,
				Type:          value.Type,
				Value:         value.Value,
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
	enumEmpty := make([]shape.Enum, 0)
	if len(enumResult) == 0 {
		return enumEmpty, count, err
	}
	return enumResult, count, err
}
