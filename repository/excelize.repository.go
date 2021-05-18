package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type excelizeRepository struct {
	db *sqlx.DB
}

type ExcelizeRepository interface {
	GetAllExcelize(skip, take int, filter string) ([]model.UserInfo, error)
}

func InitExcelizeRepository(db *sqlx.DB) ExcelizeRepository {
	return &excelizeRepository{
		db,
	}
}

func (excelizeRepository *excelizeRepository) GetAllExcelize(skip, take int, filter string) ([]model.UserInfo, error) {
	var excelizeList []model.UserInfo
	query := fmt.Sprintf(`
	SELECT
		id,
		role_code,
		role, 
		email,
		full_name,
		phone,
		profession,
		gender,
		age,
		birth,
		province,
		city,
		address,
		image_code,
		image_filename,
		image_filepath,
		is_verified,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date
	FROM private.v_users
	WHERE 
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)
	fmt.Println(query)
	rows, sqlError := excelizeRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllExcelize => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {

			var id int
			var createdDate time.Time
			var phone, age sql.NullInt64
			var modifiedDate, births sql.NullTime
			var isVerified, isActive bool
			var emailUsr, roleCode, role, createdBy string
			var fullname, profession, gender, province, city, address, imageCode, imageFilename, imageFilepath, modifiedBy sql.NullString

			sqlError := rows.Scan(
				&id,
				&roleCode,
				&role,
				&emailUsr,
				&fullname,
				&phone,
				&profession,
				&gender,
				&age,
				&births,
				&province,
				&city,
				&address,
				&imageCode,
				&imageFilename,
				&imageFilepath,
				&isVerified,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllExcelize => ", sqlError)
			} else {
				excelizeList = append(
					excelizeList,
					model.UserInfo{
						ID:            id,
						RoleCode:      roleCode,
						Role:          role,
						Email:         emailUsr,
						FullName:      fullname,
						Phone:         phone,
						Profession:    profession,
						Gender:        gender,
						Age:           age,
						Birth:         births,
						Province:      province,
						City:          city,
						Address:       address,
						ImageCode:     imageCode,
						ImageFilename: imageFilename,
						ImageFilepath: imageFilepath,
						IsVerified:    isVerified,
						IsActive:      isActive,
						CreatedBy:     createdBy,
						CreatedDate:   createdDate,
						ModifiedBy:    modifiedBy,
						ModifiedDate:  modifiedDate,
					},
				)
			}
		}
	}
	return excelizeList, sqlError
}
