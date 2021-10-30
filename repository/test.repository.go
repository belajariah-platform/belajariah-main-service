package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	_getAllTest = `
		SELECT
			id,
			code,
			class_code,
			test_type_code,
			question,
			option_a,
			option_b,
			option_c,
			option_d,
			answer,
			test_image,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM master.v_m_test
		WHERE 
			is_deleted = false AND
			is_active=true
		%s
		OFFSET %d
		LIMIT %d
	`
	_getAllTestCount = `
		SELECT COUNT(*) FROM 
			master.v_m_test  
		WHERE 
			is_deleted=false AND
			is_active=true
		%s
	`
	_correctionTest = `
		SELECT COUNT(*) FROM 
			master.v_m_test vmct
		WHERE 
			code = '%s' AND answer = %d
	`
)

type testRepository struct {
	db *sqlx.DB
}

type TestRepository interface {
	GetAllTest(skip, take int, filter string) ([]model.ClassTest, error)
	GetAllTestCount(filter string) (int, error)
	CorrectionTest(code string, answer int) (int, error)
}

func InitTestRepository(db *sqlx.DB) TestRepository {
	return &testRepository{
		db,
	}
}

func (testRepository *testRepository) GetAllTest(skip, take int, filter string) ([]model.ClassTest, error) {
	var testList []model.ClassTest
	query := fmt.Sprintf(_getAllTest, filter, skip, take)

	rows, sqlError := testRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllTest => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var id, Answer int
			var createdDate time.Time
			var isActive, isDeleted bool
			var modifiedDate sql.NullTime
			var TestImage, modifiedBy sql.NullString
			var code, ClassCode, TestTypeCode, Question, OptionA, OptionB, OptionC, OptionD, createdBy string

			sqlError := rows.Scan(
				&id,
				&code,
				&ClassCode,
				&TestTypeCode,
				&Question,
				&OptionA,
				&OptionB,
				&OptionC,
				&OptionD,
				&Answer,
				&TestImage,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&isDeleted,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllTest => ", sqlError)
			} else {
				testList = append(
					testList,
					model.ClassTest{
						ID:           id,
						Code:         code,
						ClassCode:    ClassCode,
						TestTypeCode: TestTypeCode,
						Question:     Question,
						OptionA:      OptionA,
						OptionB:      OptionB,
						OptionC:      OptionC,
						OptionD:      OptionD,
						Answer:       Answer,
						TestImage:    TestImage,
						IsActive:     isActive,
						CreatedBy:    createdBy,
						CreatedDate:  createdDate,
						ModifiedBy:   modifiedBy,
						ModifiedDate: modifiedDate,
						IsDeleted:    isDeleted,
					},
				)
			}
		}
	}
	return testList, sqlError
}

func (testRepository *testRepository) GetAllTestCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(_getAllTestCount, filter)

	row := testRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllTestCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}

func (testRepository *testRepository) CorrectionTest(code string, answer int) (int, error) {
	var count int
	query := fmt.Sprintf(_correctionTest, code, answer)

	row := testRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on CorrectionTest => ", sqlError)
		count = 0
	}
	return count, sqlError
}
