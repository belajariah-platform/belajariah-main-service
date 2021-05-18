package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
)

type excelizeUsecase struct {
	excelizeRepository repository.ExcelizeRepository
}

type ExcelizeUsecase interface {
	GetAllExcelize(query model.Query) ([]shape.UserInfo, error)
}

func InitExcelizeUsecase(excelizeRepository repository.ExcelizeRepository) ExcelizeUsecase {
	return &excelizeUsecase{
		excelizeRepository,
	}
}

func (excelizeUsecase *excelizeUsecase) GetAllExcelize(query model.Query) ([]shape.UserInfo, error) {
	var filterQuery string
	var excelizes []model.UserInfo
	var excelizeResult []shape.UserInfo

	filterQuery = utils.GetFilterHandler(query.Filters)

	excelizes, err := excelizeUsecase.excelizeRepository.GetAllExcelize(query.Skip, query.Take, filterQuery)
	if err == nil {
		for _, user := range excelizes {
			excelizeResult = append(excelizeResult, shape.UserInfo{
				ID:             user.ID,
				Role_Code:      user.RoleCode,
				Role:           user.Role,
				Email:          user.Email,
				Full_Name:      user.FullName.String,
				Phone:          int(user.Phone.Int64),
				Profession:     user.Profession.String,
				Gender:         user.Gender.String,
				Age:            int(user.Age.Int64),
				Birth:          utils.HandleNullableDate(user.Birth.Time),
				Province:       user.Province.String,
				City:           user.City.String,
				Address:        user.Address.String,
				Image_Code:     user.ImageCode.String,
				Image_Filename: user.ImageFilename.String,
				Is_Verified:    user.IsVerified,
				Is_Active:      user.IsActive,
				Created_By:     user.CreatedBy,
				Created_Date:   user.CreatedDate,
				Modified_By:    user.ModifiedBy.String,
				Modified_Date:  user.ModifiedDate.Time,
			})
		}
	}
	excelizeEmpty := make([]shape.UserInfo, 0)
	if len(excelizeResult) == 0 {
		return excelizeEmpty, err
	}
	return excelizeResult, err
}
