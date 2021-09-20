package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type classRepository struct {
	db *sqlx.DB
}

type ClassRepository interface {
	GetAllClass(skip, take int, filter string) ([]model.Class, error)
	GetAllClassCount(filter string) (int, error)
}

func InitClassRepository(db *sqlx.DB) ClassRepository {
	return &classRepository{
		db,
	}
}

func (classRepository *classRepository) GetAllClass(skip, take int, filter string) ([]model.Class, error) {
	var classList []model.Class
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		class_category_code,
		class_category,
		class_name,
		class_initial,
		class_description,
		class_image,
		class_video,
		class_rating,
		total_review,
		total_video,
		coalesce(total_video_duration, 0),
		instructor_name,
		instructor_description,
		instructor_biografi,
		instrcutor_image,
		is_direct,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM v_m_class
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := classRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllClass => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var createdDate time.Time
			var isActive, isDirect bool
			var id, totalReview, totalVideo int
			var modifiedDate, deletedDate sql.NullTime
			var classRating, totalVideoDuration float64
			var classInitial, classDescription, classImage, classVideo, instructorDescription, instructorBiografi, instructorImage, modifiedBy, deletedBy sql.NullString
			var code, classCategoryCode, classCategory, className, instructorName, createdBy string

			sqlError := rows.Scan(
				&id,
				&code,
				&classCategoryCode,
				&classCategory,
				&className,
				&classInitial,
				&classDescription,
				&classImage,
				&classVideo,
				&classRating,
				&totalReview,
				&totalVideo,
				&totalVideoDuration,
				&instructorName,
				&instructorDescription,
				&instructorBiografi,
				&instructorImage,
				&isDirect,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
			)
			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllClass => ", sqlError.Error())
			} else {
				classList = append(
					classList,
					model.Class{
						ID:                    id,
						Code:                  code,
						ClassCategoryCode:     classCategoryCode,
						ClassCategory:         classCategory,
						ClassName:             className,
						ClassInitial:          classInitial,
						ClassDescription:      classDescription,
						ClassImage:            classImage,
						ClassVideo:            classVideo,
						ClassRating:           classRating,
						TotalReview:           totalReview,
						TotalVideo:            totalVideo,
						TotalVideoDuration:    totalVideoDuration,
						InstructorName:        instructorName,
						InstructorDescription: instructorDescription,
						InstructorBiografi:    instructorBiografi,
						InstructorImage:       instructorImage,
						IsDirect:              isDirect,
						IsActive:              isActive,
						CreatedBy:             createdBy,
						CreatedDate:           createdDate,
						ModifiedBy:            modifiedBy,
						ModifiedDate:          modifiedDate,
						DeletedBy:             deletedBy,
						DeletedDate:           deletedDate,
					},
				)
			}
		}
	}
	return classList, sqlError
}

func (classRepository *classRepository) GetAllClassCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		v_m_class  
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	`, filter)

	row := classRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllClassCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}
