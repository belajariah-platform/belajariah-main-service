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
	GetUserClass(filter string) (model.UserClass, error)
	GetAllUserClassCount(filter, filterUser string) (int, error)
	GetAllUserClass(skip, take int, sorting, filter, filterUser string) ([]model.UserClass, error)

	InsertUserClass(userClass model.UserClass) (bool, error)
	UpdateUserClass(userClass model.UserClass) (bool, error)
	UpdateUserClassExpired(userClass model.UserClass) (bool, error)
	UpdateUserClassProgress(userClass model.UserClass) (bool, error)
	DeleteUserClass(userClass model.UserClass) (time.Time, bool, error)
	UpdateUserClassTest(userClass model.UserClass, types string) (bool, error)

	CheckAllUserClassExpired() ([]model.UserClass, error)
	CheckAllUserClassBeforeExpired(interval model.TimeInterval) ([]model.UserClass, error)
}

func InitUserClassRepository(db *sqlx.DB) UserClassRepository {
	return &userClassRepository{
		db,
	}
}

func (userClassRepository *userClassRepository) GetUserClass(filter string) (model.UserClass, error) {
	var userClassRow model.UserClass
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		user_code,
		class_code,
		total_user,
		type_code,
		type,
		status_code,
		status,
		package_code,
		package_type,
		is_expired,
		start_date,
		expired_date,
		time_duration,
		progress,
		progress_count,
		progress_index,
		progress_subindex,
		pre_test_scores,
		post_test_scores,
		post_test_date,
		pre_test_total,
		post_test_total,
		total_consultation,
		total_webinar
	FROM 
		transaction.v_t_user_class
	WHERE 
		is_deleted=false
		%s
	`, filter)
	row := userClassRepository.db.QueryRow(query)

	var isExpired bool
	var postTestDate sql.NullTime
	var id, totalUser, timeDuration int
	var startDate, expiredDate time.Time
	var preTestTotal, postTestTotal sql.NullInt64
	var progress, preTestScores, postTestScores sql.NullFloat64
	var packageCode, packageType, typeCode, types, status, statusCode, classCode, userCode, code string
	var progressCount, progressIndex, progressSubindex, totalConsultation, totalWebinar sql.NullInt64

	sqlError := row.Scan(
		&id,
		&code,
		&userCode,
		&classCode,
		&totalUser,
		&typeCode,
		&types,
		&statusCode,
		&status,
		&packageCode,
		&packageType,
		&isExpired,
		&startDate,
		&expiredDate,
		&timeDuration,
		&progress,
		&progressCount,
		&progressIndex,
		&progressSubindex,
		&preTestScores,
		&postTestScores,
		&postTestDate,
		&preTestTotal,
		&postTestTotal,
		&totalConsultation,
		&totalWebinar,
	)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetUserClass => ", sqlError.Error())
		return model.UserClass{}, nil
	} else {
		userClassRow = model.UserClass{
			ID:                id,
			Code:              code,
			UserCode:          userCode,
			ClassCode:         classCode,
			TotalUser:         totalUser,
			TypeCode:          typeCode,
			Type:              types,
			StatusCode:        statusCode,
			Status:            status,
			PackageCode:       packageCode,
			PackageType:       packageType,
			IsExpired:         isExpired,
			StartDate:         startDate,
			ExpiredDate:       expiredDate,
			TimeDuration:      timeDuration,
			Progress:          progress,
			ProgressCount:     progressCount,
			ProgressIndex:     progressIndex,
			ProgressSubindex:  progressSubindex,
			PreTestScores:     preTestScores,
			PostTestScores:    postTestScores,
			PostTestDate:      postTestDate,
			PreTestTotal:      preTestTotal,
			PostTestTotal:     postTestTotal,
			TotalConsultation: totalConsultation,
			TotalWebinar:      totalWebinar,
		}
		return userClassRow, sqlError
	}
}

func (userClassRepository *userClassRepository) GetAllUserClass(skip, take int, sort, filter, filterUser string) ([]model.UserClass, error) {
	var userClassList []model.UserClass
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		user_code,
		class_code,
		class_name,
		class_initial,
		class_category,
		class_description,
		class_image,
		class_rating,
		total_user,
		type_code,
		type,
		status_code,
		status,
		package_code,
		package_type,
		is_expired,
		start_date,
		expired_date,
		time_duration,
		progress,
		progress_count,
		progress_index,
		progress_subindex,
		pre_test_scores,
		post_test_scores,
		post_test_date,
		pre_test_total,
		post_test_total,
		total_consultation,
		total_webinar,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM transaction.v_t_user_class
	WHERE 
		is_deleted=false
		%s
	%s %s
	OFFSET %d
	LIMIT %d
	`, filterUser, filter, sort, skip, take)

	rows, sqlError := userClassRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllUserClass => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var classRating float64
			var isExpired, isActive bool
			var id, totalUser, timeDuration int
			var preTestTotal, postTestTotal sql.NullInt64
			var startDate, expiredDate, createdDate time.Time
			var postTestDate, modifiedDate, deletedDate sql.NullTime
			var progress, preTestScores, postTestScores sql.NullFloat64
			var classInitial, classDescription, classImage, modifiedBy, deletedBy sql.NullString
			var packageCode, packageType, typeCode, types, status, statusCode, classCode, className, classCategory, createdBy,
				code, userCode string
			var progressCount, progressIndex, progressSubindex, totalConsultation, totalWebinar sql.NullInt64

			sqlError := rows.Scan(
				&id,
				&code,
				&userCode,
				&classCode,
				&className,
				&classInitial,
				&classCategory,
				&classDescription,
				&classImage,
				&classRating,
				&totalUser,
				&typeCode,
				&types,
				&statusCode,
				&status,
				&packageCode,
				&packageType,
				&isExpired,
				&startDate,
				&expiredDate,
				&timeDuration,
				&progress,
				&progressCount,
				&progressIndex,
				&progressSubindex,
				&preTestScores,
				&postTestScores,
				&postTestDate,
				&preTestTotal,
				&postTestTotal,
				&totalConsultation,
				&totalWebinar,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllUserClass => ", sqlError.Error())
			} else {
				userClassList = append(
					userClassList,
					model.UserClass{
						ID:                id,
						Code:              code,
						UserCode:          userCode,
						ClassCode:         classCode,
						ClassName:         className,
						ClassInitial:      classInitial,
						ClassCategory:     classCategory,
						ClassDescription:  classDescription,
						ClassImage:        classImage,
						ClassRating:       classRating,
						TotalUser:         totalUser,
						TypeCode:          typeCode,
						Type:              types,
						StatusCode:        statusCode,
						Status:            status,
						PackageCode:       packageCode,
						PackageType:       packageType,
						IsExpired:         isExpired,
						StartDate:         startDate,
						ExpiredDate:       expiredDate,
						TimeDuration:      timeDuration,
						Progress:          progress,
						ProgressCount:     progressCount,
						ProgressIndex:     progressIndex,
						ProgressSubindex:  progressSubindex,
						PreTestScores:     preTestScores,
						PostTestScores:    postTestScores,
						PostTestDate:      postTestDate,
						PreTestTotal:      preTestTotal,
						PostTestTotal:     postTestTotal,
						TotalConsultation: totalConsultation,
						TotalWebinar:      totalWebinar,
						IsActive:          isActive,
						CreatedBy:         createdBy,
						CreatedDate:       createdDate,
						ModifiedBy:        modifiedBy,
						ModifiedDate:      modifiedDate,
						DeletedBy:         deletedBy,
						DeletedDate:       deletedDate,
					},
				)
			}
		}
	}
	return userClassList, sqlError
}

func (userClassRepository *userClassRepository) GetAllUserClassCount(filter, filterUser string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		transaction.v_t_user_class  
	WHERE 
		is_deleted=false
		%s
	%s
	`, filterUser, filter)

	row := userClassRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllUserClassCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (userClassRepository *userClassRepository) InsertUserClass(userClass model.UserClass) (bool, error) {
	var err error
	var result bool

	tx, errTx := userClassRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in insertUserClass", errTx)
	} else {
		err = insertUserClass(tx, userClass)
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
		utils.PushLogf("failed to InsertUserClass", err)
	}

	return result, err
}

func insertUserClass(tx *sql.Tx, userClass model.UserClass) error {
	_, err := tx.Exec(`
	INSERT INTO transact_user_class
	(
		user_code,
		class_code,
		package_code,
		type_code,
		status_code,
		start_date,
		expired_date,
		time_duration,
		created_by,
		created_date,
		modified_by,
		modified_date,
		total_consultation,
		total_webinar
	)
	VALUES(
		$1,
		$2,
		$3,
		(SELECT code 
			FROM master_enum me 
			WHERE lower(value)=lower('new class') LIMIT 1),
		(SELECT code 
			FROM master_enum me 
			WHERE lower(value)=lower('start') LIMIT 1),
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		$12
	);
	`,
		userClass.UserCode,
		userClass.ClassCode,
		userClass.PackageCode,
		userClass.StartDate,
		userClass.ExpiredDate,
		userClass.TimeDuration,
		userClass.CreatedBy,
		userClass.CreatedDate,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
		userClass.TotalConsultation.Int64,
		userClass.TotalWebinar.Int64,
	)
	return err
}

func (userClassRepository *userClassRepository) DeleteUserClass(userClass model.UserClass) (time.Time, bool, error) {
	var err error
	var result bool
	var date time.Time

	tx, errTx := userClassRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in DeleteUserClass", errTx)
	} else {
		date, err = deleteUserClass(tx, userClass)
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
		utils.PushLogf("failed to DeleteUserClass", err)
	}

	return date, result, err
}

func deleteUserClass(tx *sql.Tx, userClass model.UserClass) (time.Time, error) {
	var expiredDate time.Time
	err := tx.QueryRow(`
	UPDATE
		transact_user_class
	 SET
		deleted_by=$1,
		deleted_date=$2
 	WHERE
 		class_code=$3 AND
		user_code=$4 
		returning expired_date
	`,
		userClass.DeletedBy.String,
		userClass.DeletedDate.Time,
		userClass.ClassCode,
		userClass.UserCode,
	).Scan(&expiredDate)
	return expiredDate, err
}

func (userClassRepository *userClassRepository) UpdateUserClass(userClass model.UserClass) (bool, error) {
	var err error
	var result bool

	tx, errTx := userClassRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in UpdateUserClass", errTx)
	} else {
		err = updateUserClass(tx, userClass)
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
		utils.PushLogf("failed to UpdateUserClass", err)
	}

	return result, err
}

func updateUserClass(tx *sql.Tx, userClass model.UserClass) error {
	_, err := tx.Exec(`
	UPDATE
		transact_user_class
	 SET
		package_code=$1,
		type_code=(SELECT code 
			FROM master_enum me 
			WHERE value='Extend Class' LIMIT 1),
		is_expired=false,
		start_date=$2,
		expired_date=$3,
		time_duration=$4,
		pre_test_scores=default,
		post_test_scores=default,
		post_test_date=default,
		modified_by=$5,
		modified_date=$6,
		total_consultation=$7,
		total_webinar=$8,
		pre_test_total=0,
		post_test_total=0
 	WHERE
 		class_code=$9 AND
		user_code=$10 
	`,
		userClass.PackageCode,
		userClass.StartDate,
		userClass.ExpiredDate,
		userClass.TimeDuration,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
		userClass.TotalConsultation.Int64,
		userClass.TotalWebinar.Int64,
		userClass.ClassCode,
		userClass.UserCode,
	)
	return err
}

func (userClassRepository *userClassRepository) UpdateUserClassExpired(userClass model.UserClass) (bool, error) {
	var err error
	var result bool

	tx, errTx := userClassRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in UpdateUserClassExpired", errTx)
	} else {
		err = updateUserClassExpired(tx, userClass)
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
		utils.PushLogf("failed to UpdateUserClassExpired", err)
	}

	return result, err
}

func updateUserClassExpired(tx *sql.Tx, userClass model.UserClass) error {
	_, err := tx.Exec(`
	UPDATE
		transact_user_class
	 SET
		is_expired=true,
		modified_by=$1,
		modified_date=$2
 	WHERE
 		id=$3 AND
		user_code=$4 
	`,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
		userClass.ID,
		userClass.UserCode,
	)
	return err
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
		status_code=(SELECT code 
			FROM master_enum me 
			WHERE value=$2 LIMIT 1),
		modified_by=$3,
		modified_date=$4,
		progress_index=$5,
		progress_subindex=$6,
		progress_count=$7
 	WHERE
 		id=$8 AND
		user_code=$9
	`,
		userClass.Progress.Float64,
		userClass.Status,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
		userClass.ProgressIndex.Int64,
		userClass.ProgressSubindex.Int64,
		userClass.ProgressCount.Int64,
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
		pre_test_total=pre_test_total+1,
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
			post_test_total=post_test_total+1,
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

func (userClassRepository *userClassRepository) CheckAllUserClassExpired() ([]model.UserClass, error) {
	var userClassList []model.UserClass

	rows, sqlError := userClassRepository.db.Query(`
	SELECT
		id,
		code,
		user_code,
		class_code,
		status_code,
		status,
		is_expired,
		start_date,
		expired_date,
		time_duration,
		progress
	FROM transaction.v_t_user_class
	WHERE  	
		is_deleted=false AND
		is_expired = false AND
		expired_date <= now() 
	`)

	if sqlError != nil {
		utils.PushLogf("SQL error on CheckAllUserClassExpired => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var isExpired bool
			var id, timeDuration int
			var progress sql.NullFloat64
			var startDate, expiredDate time.Time
			var status, statusCode, classCode, userCode, code string

			sqlError := rows.Scan(
				&id,
				&code,
				&userCode,
				&classCode,
				&statusCode,
				&status,
				&isExpired,
				&startDate,
				&expiredDate,
				&timeDuration,
				&progress,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on CheckAllUserClassExpired => ", sqlError.Error())
			} else {
				userClassList = append(userClassList, model.UserClass{
					ID:           id,
					Code:         code,
					UserCode:     userCode,
					ClassCode:    classCode,
					StatusCode:   statusCode,
					Status:       status,
					IsExpired:    isExpired,
					StartDate:    startDate,
					ExpiredDate:  expiredDate,
					TimeDuration: timeDuration,
					Progress:     progress,
				})
			}
		}
	}
	return userClassList, sqlError
}

func (userClassRepository *userClassRepository) CheckAllUserClassBeforeExpired(interval model.TimeInterval) ([]model.UserClass, error) {
	var userClassList []model.UserClass

	rows, sqlError := userClassRepository.db.Query(`
	SELECT
		id,
		code,
		user_code,
		class_code,
		status_code,
		status,
		is_expired,
		start_date,
		expired_date,
		time_duration,
		progress
	FROM transaction.v_t_user_class
	WHERE  	
		is_deleted=false AND
		is_expired = false AND
		DATE_PART('day', expired_date::timestamp - now()::timestamp) * 24 * 60 * 60 + 
		DATE_PART('hour', expired_date::timestamp - now()::timestamp) * 60 * 60 +
		DATE_PART('minute', expired_date::timestamp - now()::timestamp) * 60 +
		DATE_PART('second', expired_date::timestamp - now()::timestamp) <= $1 AND
		DATE_PART('day', expired_date::timestamp - now()::timestamp) * 24 * 60 * 60 + 
		DATE_PART('hour', expired_date::timestamp - now()::timestamp) * 60 * 60 +
		DATE_PART('minute', expired_date::timestamp - now()::timestamp) * 60 +
		DATE_PART('second', expired_date::timestamp - now()::timestamp) >= $2
	`, interval.Interval1, interval.Interval2)

	if sqlError != nil {
		utils.PushLogf("SQL error on CheckAllUserClassExpired => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var isExpired bool
			var id, timeDuration int
			var progress sql.NullFloat64
			var startDate, expiredDate time.Time
			var status, statusCode, classCode, code, userCode string

			sqlError := rows.Scan(
				&id,
				&code,
				&userCode,
				&classCode,
				&statusCode,
				&status,
				&isExpired,
				&startDate,
				&expiredDate,
				&timeDuration,
				&progress,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on CheckAllUserClassExpired => ", sqlError.Error())
			} else {
				userClassList = append(userClassList, model.UserClass{
					ID:           id,
					Code:         code,
					UserCode:     userCode,
					ClassCode:    classCode,
					StatusCode:   statusCode,
					Status:       status,
					IsExpired:    isExpired,
					StartDate:    startDate,
					ExpiredDate:  expiredDate,
					TimeDuration: timeDuration,
					Progress:     progress,
				})
			}
		}
	}
	return userClassList, sqlError
}
