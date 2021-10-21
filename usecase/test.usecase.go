package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
	"database/sql"
	"time"
)

type testUsecase struct {
	testRepository      repository.TestRepository
	userClassRepository repository.UserClassRepository
}

type TestUsecase interface {
	GetAllTest(query model.Query) ([]shape.ClassTest, int, error)
	CorrectionTest(test shape.ClassTestPost, email string) (bool, float64, error)
}

func InitTestUsecase(testRepository repository.TestRepository, userClassRepository repository.UserClassRepository) TestUsecase {
	return &testUsecase{
		testRepository,
		userClassRepository,
	}
}

func (testUsecase *testUsecase) GetAllTest(query model.Query) ([]shape.ClassTest, int, error) {
	var filterQuery string
	var tests []model.ClassTest
	var testResult []shape.ClassTest

	filterQuery = utils.GetFilterHandler(query.Filters)

	tests, err := testUsecase.testRepository.GetAllTest(query.Skip, query.Take, filterQuery)
	count, errCount := testUsecase.testRepository.GetAllTestCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range tests {
			testResult = append(testResult, shape.ClassTest{
				ID:             value.ID,
				Code:           value.Code,
				Class_Code:     value.ClassCode,
				Test_Type_Code: value.TestTypeCode,
				Question:       value.Question,
				Option_A:       value.OptionA,
				Option_B:       value.OptionB,
				Option_C:       value.OptionC,
				Option_D:       value.OptionD,
				Test_Image:     value.TestImage.String,
				Is_Active:      value.IsActive,
				Created_By:     value.CreatedBy,
				Created_Date:   value.CreatedDate,
				Modified_By:    value.ModifiedBy.String,
				Modified_Date:  value.ModifiedDate.Time,
				Is_Deleted:     value.IsDeleted,
			})
		}
	}
	testEmpty := make([]shape.ClassTest, 0)
	if len(testResult) == 0 {
		return testEmpty, count, err
	}
	return testResult, count, err
}

func (testUsecase *testUsecase) CorrectionTest(test shape.ClassTestPost, email string) (bool, float64, error) {
	var countResult int
	var testResult float64

	for _, value := range test.Answers {
		count, err := testUsecase.testRepository.CorrectionTest(value.Code, value.Answer)
		if err == nil {
			countResult = countResult + (count * 100)
		}
	}
	testResult = float64(countResult / len(test.Answers))
	dataTest := model.UserClass{
		ID: test.ID,
		PreTestScores: sql.NullFloat64{
			Float64: testResult,
		},
		PostTestScores: sql.NullFloat64{
			Float64: testResult,
		},
		PostTestDate: sql.NullTime{
			Time: time.Now(),
		},
		ModifiedBy: sql.NullString{
			String: email,
		},
		ModifiedDate: sql.NullTime{
			Time: time.Now(),
		},
	}
	result, err := testUsecase.userClassRepository.UpdateUserClassTest(dataTest, test.Test_Type)
	if err != nil && !result {
		testResult = 0
	}
	return result, testResult, err
}
