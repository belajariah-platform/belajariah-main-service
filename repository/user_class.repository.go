package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type userClassRepository struct {
	db *sqlx.DB
}

type UserClassRepository interface {
	GetAllUserClass(skip, take int, filter, filterUser string) ([]model.UserClass, error)
	GetAllUserClassCount(filter, filterUser string) (int, error)

	UpdateUserClassProgress(userClass model.UserClass) (bool, error)
	UpdateUserClassTest(userClass model.UserClass, types string) (bool, error)
}

func InitUserClassRepository(db *sqlx.DB) UserClassRepository {
	return &userClassRepository{
		db,
	}
}

func (userClassRepository *userClassRepository) GetAllUserClass(skip, take int, filter, filterUser string) ([]model.UserClass, error) {
	var paymentList []model.UserClass
	query := fmt.Sprintf(`
	SELECT
		id,
		user_code,
		class_code,
		class_name,
		class_initial,
		class_category,
		class_description,
		class_image,
		class_rating,
		total_user,
		status_code,
		status,
		is_expired,
		start_date,
		expired_date,
		time_duration,
		progress,
		progress_index,
		progress_cur_index,
		progress_cur_subindex,
		pre_test_scores,
		post_test_scores,
		post_test_date,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM v_t_user_class
	WHERE 
		deleted_by IS NULL
		%s
	%s
	OFFSET %d
	LIMIT %d
	`, filterUser, filter, skip, take)

	rows, sqlError := userClassRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllUserClass => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var classRating float64
			var isExpired, isActive bool
			var id, userCode, totalUser, timeDuration int
			var startDate, expiredDate, createdDate time.Time
			var postTestDate, modifiedDate, deletedDate sql.NullTime
			var progress, preTestScores, postTestScores sql.NullFloat64
			var progressIndex, progressCurIndex, progressCurSubindex sql.NullInt64
			var status, statusCode, classCode, className, classCategory, createdBy string
			var classInitial, classDescription, classImage, modifiedBy, deletedBy sql.NullString

			sqlError := rows.Scan(
				&id,
				&userCode,
				&classCode,
				&className,
				&classInitial,
				&classCategory,
				&classDescription,
				&classImage,
				&classRating,
				&totalUser,
				&statusCode,
				&status,
				&isExpired,
				&startDate,
				&expiredDate,
				&timeDuration,
				&progress,
				&progressIndex,
				&progressCurIndex,
				&progressCurSubindex,
				&preTestScores,
				&postTestScores,
				&postTestDate,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllUserClass => ", sqlError)
			} else {
				paymentList = append(
					paymentList,
					model.UserClass{
						ID:                  id,
						UserCode:            userCode,
						ClassCode:           classCode,
						ClassName:           className,
						ClassInitial:        classInitial,
						ClassCategory:       classCategory,
						ClassDescription:    classDescription,
						ClassImage:          classImage,
						ClassRating:         classRating,
						TotalUser:           totalUser,
						StatusCode:          statusCode,
						Status:              status,
						IsExpired:           isExpired,
						StartDate:           startDate,
						ExpiredDate:         expiredDate,
						TimeDuration:        timeDuration,
						Progress:            progress,
						ProgressIndex:       progressIndex,
						ProgressCurIndex:    progressCurIndex,
						ProgressCurSubindex: progressCurSubindex,
						PreTestScores:       preTestScores,
						PostTestScores:      postTestScores,
						PostTestDate:        postTestDate,
						IsActive:            isActive,
						CreatedBy:           createdBy,
						CreatedDate:         createdDate,
						ModifiedBy:          modifiedBy,
						ModifiedDate:        modifiedDate,
						DeletedBy:           deletedBy,
						DeletedDate:         deletedDate,
					},
				)
			}
		}
	}
	return paymentList, sqlError
}

func (userClassRepository *userClassRepository) GetAllUserClassCount(filter, filterUser string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		v_t_user_class  
	WHERE 
		deleted_by IS NULL
		%s
	%s
	`, filterUser, filter)

	row := userClassRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllUserClassCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}
func (userClassRepository *userClassRepository) UpdateUserClassProgress(userClass model.UserClass) (bool, error) {
	var err error
	var result bool

	tx, errTx := userClassRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in UpdateUserClassProgress", errTx)
	} else {
		err = updateUserClassProgress(tx, userClass)
		if err != nil {
			utils.PushLogf("err in user-class---", err)
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to UpdateUserClassProgress", err)
	}

	return result, err
}

func updateUserClassProgress(tx *sql.Tx, userClass model.UserClass) error {
	_, err := tx.Exec(`
	UPDATE
		transact_user_class
	 SET
		progress=$1,
		modified_by=$2,
		modified_date=$3,
		progress_index=$4,
		progress_cur_index=$5,
		progress_cur_subindex=$6
 	WHERE
 		id=$7 AND
		user_code=$8 
	`,
		userClass.Progress.Float64,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
		userClass.ProgressIndex.Int64,
		userClass.ProgressCurIndex.Int64,
		userClass.ProgressCurSubindex.Int64,
		userClass.ID,
		userClass.UserCode,
	)
	return err
}

func (userClassRepository *userClassRepository) UpdateUserClassTest(userClass model.UserClass, types string) (bool, error) {
	var err error
	var result bool

	tx, errTx := userClassRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in UpdateUserClassTest", errTx)
	} else {
		if strings.ToLower(types) == "pre-test" {
			err = updateUserClassPreTest(tx, userClass)
			if err != nil {
				utils.PushLogf("err in user-class---", err)
			}
		} else if strings.ToLower(types) == "post-test" {
			err = updateUserClassPostTest(tx, userClass)
			if err != nil {
				utils.PushLogf("err in user-class---", err)
			}
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to UpdateUserClassTest", err)
	}

	return result, err
}

func updateUserClassPreTest(tx *sql.Tx, userClass model.UserClass) error {
	_, err := tx.Exec(`
	UPDATE
		transact_user_class
	 SET
		pre_test_scores=$1,
		modified_by=$2,
		modified_date=$3
 	WHERE
 		id=$4
	`,
		userClass.PreTestScores.Float64,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
		userClass.ID,
	)
	return err
}

func updateUserClassPostTest(tx *sql.Tx, userClass model.UserClass) error {
	_, err := tx.Exec(`
	UPDATE
		transact_user_class
	 SET
		post_test_scores=$1,
		post_test_date=$2,
		modified_by=$3,
		modified_date=$4
 	WHERE
 		id=$5
	`,
		userClass.PostTestScores.Float64,
		userClass.PostTestDate.Time,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
		userClass.ID,
	)
	return err
}
