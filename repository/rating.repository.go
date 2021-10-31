package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	_getAllRatingClass = `
		SELECT
			id,
			code,
			class_code,
			class_name,
			class_initial,
			user_code,
			user_name,
			rating,
			comment,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM 
			transaction.v_t_class_rating
		WHERE 
			is_deleted=false
		%s
		OFFSET %d
		LIMIT %d
	`
	_getAllRatingClassCount = `
		SELECT COUNT(*) FROM 
			transaction.v_t_class_rating  
		WHERE 
			is_deleted=false
		%s
	`
	_giveRatingClass = `
		INSERT INTO transaction.transact_class_rating
		(
			class_code,
			user_code,
			rating,
			comment,
			created_by,
			created_date,
			modified_by,
			modified_date
		)
		VALUES (
			$1, 
			$2, 
			$3, 
			$4, 
			$5, 
			$6, 
			$7, 
			$8
			)
	`
	_giveRatingMentor = `
		INSERT INTO transaction.transact_mentor_rating
		(
			mentor_code,
			user_code,
			rating,
			comment,
			created_by,
			created_date,
			modified_by,
			modified_date
		)
		VALUES (
			$1, 
			$2, 
			$3, 
			$4, 
			$5, 
			$6, 
			$7, 
			$8
			)
	`
)

type ratingRepository struct {
	db *sqlx.DB
}

type RatingRepository interface {
	GetAllRatingClass(skip, take int, filter string) ([]model.Rating, error)
	GetAllRatingClassCount(filter string) (int, error)
	GiveRatingClass(rating model.RatingPost) (bool, error)
	GiveRatingMentor(rating model.RatingPost) (bool, error)
}

func InitRatingRepository(db *sqlx.DB) RatingRepository {
	return &ratingRepository{
		db,
	}
}

func (ratingRepository *ratingRepository) GetAllRatingClass(skip, take int, filter string) ([]model.Rating, error) {
	var ratingList []model.Rating
	query := fmt.Sprintf(_getAllRatingClass, filter, skip, take)

	rows, sqlError := ratingRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllRatingClass => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var rating float64
			var createdDate time.Time
			var isActive, isDeleted bool
			var modifiedDate sql.NullTime
			var comment, classInitial, modifiedBy sql.NullString
			var classCode, className, userName, createdBy, code, userCode string

			sqlError := rows.Scan(
				&id,
				&code,
				&classCode,
				&className,
				&classInitial,
				&userCode,
				&userName,
				&rating,
				&comment,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&isDeleted,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllRatingClass => ", sqlError.Error())
			} else {
				ratingList = append(
					ratingList,
					model.Rating{
						ID:           id,
						Code:         code,
						ClassCode:    classCode,
						ClassName:    className,
						ClassInitial: classInitial,
						UserCode:     userCode,
						UserName:     userName,
						Rating:       rating,
						Comment:      comment,
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
	return ratingList, sqlError
}

func (ratingRepository *ratingRepository) GetAllRatingClassCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(_getAllRatingClassCount, filter)

	row := ratingRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllRatingClassCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (r *ratingRepository) GiveRatingClass(data model.RatingPost) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("ratingRepository: GiveRatingClass: error begin transaction")
	}

	_, err = tx.Exec(_giveRatingClass,
		data.ClassCode,
		data.UserCode,
		data.Rating,
		data.Comment.String,
		data.CreatedBy,
		data.CreatedDate,
		data.ModifiedBy.String,
		data.ModifiedDate.Time,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "ratingRepository: GiveRatingClass: error insert")
	}

	tx.Commit()
	return err == nil, nil
}

func (r *ratingRepository) GiveRatingMentor(data model.RatingPost) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("ratingRepository: GiveRatingMentor: error begin transaction")
	}

	_, err = tx.Exec(_giveRatingMentor,
		data.MentorCode,
		data.UserCode,
		data.Rating,
		data.Comment.String,
		data.CreatedBy,
		data.CreatedDate,
		data.ModifiedBy.String,
		data.ModifiedDate.Time,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "ratingRepository: GiveRatingMentor: error insert")
	}

	tx.Commit()
	return err == nil, nil
}
