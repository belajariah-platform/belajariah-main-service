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
	_getAllClass = `
		SELECT
			id,
			code,
			class_category_code,
			class_category,
			class_name,
			class_initial,
			class_description,
			class_image,
			class_image_header,
			class_video,
			class_document,
			class_rating,
			total_review,
			total_video,
			coalesce(total_video_duration, 0),
			instructor_name,
			instructor_description,
			instructor_biografi,
			instrcutor_image,
			color_path,
			is_direct,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM master.v_m_class
		WHERE 
			is_deleted = false AND
			is_active=true
		%s
		OFFSET %d
		LIMIT %d
	`
	_getAllClassQuran = `
		SELECT
			id,
			code,
			class_category_code,
			class_category,
			class_name,
			class_initial,
			class_description,
			class_image,
			class_image_header,
			class_video,
			class_document,
			color_path,
			is_direct,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM master.v_m_class_quran
		%s
	`
	_getCountClass = `
		SELECT COUNT(*) FROM 
			master.v_m_class  
		WHERE 
			is_deleted = false AND
			is_active=true
		%s
	`
	_getCountClassQuran = `
		SELECT COUNT(*) FROM 
			master.v_m_class_quran  
		%s
	`
)

type classRepository struct {
	db *sqlx.DB
}

type ClassRepository interface {
	GetAllClass(skip, take int, filter string) ([]model.Class, error)
	GetAllClassCount(filter string) (int, error)

	GetAllClassQuran(filter string) (*[]model.ClassQuran, error)
	GetAllClassQuranCount(filter string) (int, error)
}

func InitClassRepository(db *sqlx.DB) ClassRepository {
	return &classRepository{
		db,
	}
}

func (r *classRepository) GetAllClass(skip, take int, filter string) ([]model.Class, error) {
	var classList []model.Class
	query := fmt.Sprintf(_getAllClass, filter, skip, take)

	rows, sqlError := r.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllClass => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var classRating float64
			var createdDate time.Time
			var modifiedDate sql.NullTime
			var isActive, isDirect, isDeleted bool
			var id, totalReview, totalVideo, totalVideoDuration int
			var code, classCategoryCode, classCategory, className, createdBy string
			var classInitial, classDescription, classImage, classImageHeader, classVideo, classDocument, instructorDescription, instructorBiografi, instructorImage, modifiedBy, instructorName, colorPath sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&classCategoryCode,
				&classCategory,
				&className,
				&classInitial,
				&classDescription,
				&classImage,
				&classImageHeader,
				&classVideo,
				&classDocument,
				&classRating,
				&totalReview,
				&totalVideo,
				&totalVideoDuration,
				&instructorName,
				&instructorDescription,
				&instructorBiografi,
				&instructorImage,
				&colorPath,
				&isDirect,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&isDeleted,
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
						ClassImageHeader:      classImageHeader,
						ClassVideo:            classVideo,
						ClassDocument:         classDocument,
						ClassRating:           classRating,
						TotalReview:           totalReview,
						TotalVideo:            totalVideo,
						TotalVideoDuration:    totalVideoDuration,
						InstructorName:        instructorName,
						InstructorDescription: instructorDescription,
						InstructorBiografi:    instructorBiografi,
						InstructorImage:       instructorImage,
						ColorPath:             colorPath,
						IsDirect:              isDirect,
						IsActive:              isActive,
						CreatedBy:             createdBy,
						CreatedDate:           createdDate,
						ModifiedBy:            modifiedBy,
						ModifiedDate:          modifiedDate,
						IsDeleted:             isDeleted,
					},
				)
			}
		}
	}

	return classList, sqlError
}

func (r *classRepository) GetAllClassCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(_getCountClass, filter)

	row := r.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllClassCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (r *classRepository) GetAllClassQuran(filter string) (*[]model.ClassQuran, error) {
	var result []model.ClassQuran
	query := fmt.Sprintf(_getAllClassQuran, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "classRepository.GetAllClassQuran :  error get")
	}

	return &result, nil
}

func (r *classRepository) GetAllClassQuranCount(filter string) (int, error) {
	var count int

	query := fmt.Sprintf(_getCountClassQuran, filter)

	row := r.db.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return 0, utils.WrapError(err, "classRepository: GetCount: error query row")
	}

	return count, err
}
