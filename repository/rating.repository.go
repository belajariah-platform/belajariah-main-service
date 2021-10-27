package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
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
	query := fmt.Sprintf(`
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
	`, filter, skip, take)

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
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		transaction.v_t_class_rating  
	WHERE 
		is_deleted=false
	%s
	`, filter)

	row := ratingRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllRatingClassCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (ratingRepository *ratingRepository) GiveRatingClass(rating model.RatingPost) (bool, error) {
	var err error
	var result bool

	tx, errTx := ratingRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in GiveRatingClass", errTx)
	} else {
		err = insertRatingClass(tx, rating)
		if err != nil {
			utils.PushLogf("err in rating---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to GiveRatingClass", err)
	}

	return result, err
}

func insertRatingClass(tx *sql.Tx, rating model.RatingPost) error {
	sqlQuery := `
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
		);
`
	_, err := tx.Exec(sqlQuery,
		rating.ClassCode,
		rating.UserCode,
		rating.Rating,
		rating.Comment.String,
		rating.CreatedBy,
		rating.CreatedDate,
		rating.ModifiedBy.String,
		rating.ModifiedDate.Time,
	)
	return err
}

func (ratingRepository *ratingRepository) GiveRatingMentor(rating model.RatingPost) (bool, error) {
	var err error
	var result bool

	tx, errTx := ratingRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in GiveRatingMentor", errTx)
	} else {
		err = insertRatingMentor(tx, rating)
		if err != nil {
			utils.PushLogf("err in rating---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to GiveRatingMentor", err)
	}

	return result, err
}

func insertRatingMentor(tx *sql.Tx, rating model.RatingPost) error {
	sqlQuery := `
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
		);
`
	_, err := tx.Exec(sqlQuery,
		rating.MentorCode,
		rating.UserCode,
		rating.Rating,
		rating.Comment.String,
		rating.CreatedBy,
		rating.CreatedDate,
		rating.ModifiedBy.String,
		rating.ModifiedDate.Time,
	)
	return err
}
