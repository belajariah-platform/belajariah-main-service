package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	_getUserClass = `
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
			post_test_total
		FROM 
			transaction.v_t_user_class
		WHERE 
			is_deleted=false
			%s
	`
	_getAllUserClass = `
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
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM transaction.v_t_user_class
		WHERE 
			is_deleted=false
			%s
		%s %s
		OFFSET %d
		LIMIT %d
	`
	_getAllUserClassCount = `
		SELECT COUNT(*) FROM 
			transaction.v_t_user_class  
		WHERE 
			is_deleted=false
			%s
		%s
	`
	_insertUserClass = `
		INSERT INTO transaction.transact_user_class
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
			modified_date
		)
		VALUES(
			$1,
			$2,
			$3,
			(SELECT code 
				FROM master.master_enum me 
				WHERE lower(value)=lower('new class') LIMIT 1),
			(SELECT code 
				FROM master.master_enum me 
				WHERE lower(value)=lower('start') LIMIT 1),
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10
		) returning code
	`
	_updateUserClass = `
		UPDATE
			transaction.transact_user_class
		SET
			package_code=$1,
			type_code=(SELECT code 
				FROM master.master_enum me 
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
			pre_test_total=0,
			post_test_total=0
		WHERE
			class_code=$7 AND
			user_code=$8 
			returning code
	`
	_updateUserClassExpired = `
		UPDATE
			transaction.transact_user_class
		SET
			is_expired=true,
			modified_by=$1,
			modified_date=$2
		WHERE
			code=$3 AND
			user_code=$4 
	`
	_updateUserClassPreTest = `
		UPDATE
			transaction.transact_user_class
		SET
			pre_test_scores=$1,
			pre_test_total=pre_test_total+1,
			modified_by=$2,
			modified_date=$3
		WHERE
			code=$4
	`
	_updateUserClassPostTest = `
		UPDATE
			transaction.transact_user_class
		SET
			post_test_scores=$1,
			post_test_date=$2,
			post_test_total=post_test_total+1,
			modified_by=$3,
			modified_date=$4
		WHERE
			code=$5
	`
	_deleteUserClass = `
		UPDATE
			transaction.transact_user_class
		SET
			modified_by=$1,
			modified_date=$2,
			is_active=false,
			is_deleted=true
		WHERE
			class_code=$3 AND
			user_code=$4 
			returning expired_date
	`
	_checkAllUserClassExpired = `
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
	`
	_checkAllUserClassBeforeExpired = `
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
	`
	//--------------------------------------------------------------
	_getAllUserClassQuran = `
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
			color_path,
			package_code,
			package_type,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM transaction.v_t_user_class_quran
		%s
	`
	_getAllUserClassQuranCount = `
		SELECT COUNT(*) FROM 
			transaction.v_t_user_class_quran  
		%s
	`
	_getAllUserClassQuranDetail = `
	SELECT
		id,
		code,
		user_class_code,
		mentor_code,
		mentor_name,
		mentor_image,
		mentor_city,
		package_code,
		package_type,
		is_completed,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted,
		phone,
		user_name
	FROM transaction.v_t_user_class_quran_detail
	%s
	`
	_getAllUserClassQuranSchedule = `
	SELECT
		id,
		code,
		user_class_detail_code,
		start_date,
		finsih_date,
		user_message,
		mentor_message,
		sequence,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted,
		material,
		user_score
	FROM transaction.v_t_user_class_quran_schedule
	%s
	`
	_insertUserClassQuran = `
		INSERT INTO transaction.transact_user_class_quran
		(
			user_code,
			class_code,
			package_code,
			created_by,
			created_date,
			modified_by,
			modified_date
		)
		VALUES(
			$1,
			$2,
			'',
			$3,
			$4,
			$5,
			$6
		) returning code
	`
	_insertUserClassQuranDetail = `
		INSERT INTO transaction.transact_user_class_quran_detail
		(
			user_class_code,
			package_code,
			created_by,
			created_date,
			modified_by,
			modified_date
		)
		VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		) returning code
	`
	_insertUserClassQuranSchedule = `
		INSERT INTO transaction.transact_user_class_quran_schedule
		(
			user_class_detail_code,
			start_date,
			finsih_date,
			mentor_message,
			"sequence",
			created_by,
			created_date,
			modified_by,
			modified_date,
			material
		)
		VALUES(
			$1,
			$2,
			$3,
			$4,
			(SELECT max(sequence)+1 FROM transaction.transact_user_class_quran_schedule WHERE user_class_detail_code=$1)
			$5,
			$6,
			$7,
			$8,
			$9
		) returning code
	`
	_updateUserClassQuranSchedule = `
		UPDATE
			transaction.transact_user_class_quran_schedule
		SET
			start_date=$1,
			finsih_date=$2,
			mentor_message=$3,
			material=$4,
			modified_by=$5,
			modified_date=$6
		WHERE
			code=$7
	`
	_updateUserClassProgress = `
		UPDATE
			transaction.transact_user_class_quran_detail
		SET
			is_completed=true,
			modified_by=$1,
			modified_date=$2
		WHERE
			code=$3
	`
)

type userClassRepository struct {
	db *sqlx.DB
}

type UserClassRepository interface {
	GetUserClass(filter string) (model.UserClass, error)
	GetAllUserClassCount(filter, filterUser string) (int, error)
	GetAllUserClass(skip, take int, sorting, filter, filterUser string) ([]model.UserClass, error)

	DeleteUserClass(userClass model.UserClass) (time.Time, bool, error)
	InsertUserClass(userClass model.UserClass) (model.UserClass, bool, error)
	UpdateUserClass(userClass model.UserClass) (model.UserClass, bool, error)

	UpdateUserClassExpired(userClass model.UserClass) (bool, error)
	UpdateUserClassTest(userClass model.UserClass, types string) (bool, error)

	//------------------------------------------------------------------------

	GetAllUserClassQuranCount(filter string) (int, error)
	GetAllUserClassQuran(filter string) (*[]model.UserClassQuran, error)
	GetAllUserClassQuranDetail(filter string) (*[]model.UserClassQuranDetail, error)
	GetAllUserClassQuranSchedule(filter string) (*[]model.UserClassQuranSchedule, error)

	InsertUserClassQuran(userClass model.UserClassQuran) (model.UserClassQuran, bool, error)
	InsertUserClassQuranDetail(userClass model.UserClassQuran) (model.UserClassQuran, bool, error)
	InsertUserClassQuranSchedule(userClass model.UserClassQuranSchedule) (model.UserClassQuranSchedule, bool, error)
	UpdateUserClassQuranSchedule(userClass model.UserClassQuranSchedule) (bool, error)
	UpdateUserClassQuranProgress(userClass model.UserClassQuran) (bool, error)

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

	query := fmt.Sprintf(_getUserClass, filter)
	row := userClassRepository.db.QueryRow(query)

	var isExpired bool
	var id, totalUser, timeDuration int
	var preTestTotal, postTestTotal sql.NullInt64
	var postTestDate, startDate, expiredDate sql.NullTime
	var progress, preTestScores, postTestScores sql.NullFloat64
	var packageCode, packageType, typeCode, types, status, statusCode, classCode, userCode, code string
	var progressCount, progressIndex, progressSubindex sql.NullInt64

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
	)

	if sqlError != nil {
		return model.UserClass{}, nil
	} else {
		userClassRow = model.UserClass{
			ID:               id,
			Code:             code,
			UserCode:         userCode,
			ClassCode:        classCode,
			TotalUser:        totalUser,
			TypeCode:         typeCode,
			Type:             types,
			StatusCode:       statusCode,
			Status:           status,
			PackageCode:      packageCode,
			PackageType:      packageType,
			IsExpired:        isExpired,
			StartDate:        startDate,
			ExpiredDate:      expiredDate,
			TimeDuration:     timeDuration,
			Progress:         progress,
			ProgressCount:    progressCount,
			ProgressIndex:    progressIndex,
			ProgressSubindex: progressSubindex,
			PreTestScores:    preTestScores,
			PostTestScores:   postTestScores,
			PostTestDate:     postTestDate,
			PreTestTotal:     preTestTotal,
			PostTestTotal:    postTestTotal,
		}
		return userClassRow, sqlError
	}
}

func (userClassRepository *userClassRepository) GetAllUserClass(skip, take int, sort, filter, filterUser string) ([]model.UserClass, error) {
	var userClassList []model.UserClass
	query := fmt.Sprintf(_getAllUserClass, filterUser, filter, sort, skip, take)

	rows, sqlError := userClassRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllUserClass => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var classRating float64
			var createdDate time.Time
			var id, totalUser, timeDuration int
			var isExpired, isActive, is_deleted bool
			var postTestDate, modifiedDate, startDate, expiredDate sql.NullTime
			var preTestTotal, postTestTotal sql.NullInt64
			var progress, preTestScores, postTestScores sql.NullFloat64
			var classInitial, classDescription, classImage, modifiedBy sql.NullString
			var packageCode, packageType, typeCode, types, status, statusCode, classCode, className, classCategory, createdBy,
				code, userCode string
			var progressCount, progressIndex, progressSubindex sql.NullInt64

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
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&is_deleted,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllUserClass => ", sqlError.Error())
			} else {
				userClassList = append(
					userClassList,
					model.UserClass{
						ID:               id,
						Code:             code,
						UserCode:         userCode,
						ClassCode:        classCode,
						ClassName:        className,
						ClassInitial:     classInitial,
						ClassCategory:    classCategory,
						ClassDescription: classDescription,
						ClassImage:       classImage,
						ClassRating:      classRating,
						TotalUser:        totalUser,
						TypeCode:         typeCode,
						Type:             types,
						StatusCode:       statusCode,
						Status:           status,
						PackageCode:      packageCode,
						PackageType:      packageType,
						IsExpired:        isExpired,
						StartDate:        startDate,
						ExpiredDate:      expiredDate,
						TimeDuration:     timeDuration,
						Progress:         progress,
						ProgressCount:    progressCount,
						ProgressIndex:    progressIndex,
						ProgressSubindex: progressSubindex,
						PreTestScores:    preTestScores,
						PostTestScores:   postTestScores,
						PostTestDate:     postTestDate,
						PreTestTotal:     preTestTotal,
						PostTestTotal:    postTestTotal,
						IsActive:         isActive,
						CreatedBy:        createdBy,
						CreatedDate:      createdDate,
						ModifiedBy:       modifiedBy,
						ModifiedDate:     modifiedDate,
						IsDeleted:        is_deleted,
					},
				)
			}
		}
	}
	return userClassList, sqlError
}

func (userClassRepository *userClassRepository) GetAllUserClassCount(filter, filterUser string) (int, error) {
	var count int
	query := fmt.Sprintf(_getAllUserClassCount, filterUser, filter)

	row := userClassRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllUserClassCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (r *userClassRepository) GetAllUserClassQuran(filter string) (*[]model.UserClassQuran, error) {
	var result []model.UserClassQuran
	query := fmt.Sprintf(_getAllUserClassQuran, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "userClassRepository.GetAllUserClassQuran :  error get")
	}

	return &result, nil
}

func (userClassRepository *userClassRepository) GetAllUserClassQuranCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(_getAllUserClassQuranCount, filter)

	row := userClassRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllUserClassQuranCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (r *userClassRepository) GetAllUserClassQuranDetail(filter string) (*[]model.UserClassQuranDetail, error) {
	var result []model.UserClassQuranDetail
	query := fmt.Sprintf(_getAllUserClassQuranDetail, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "userClassRepository.GetAllUserClassQuranDetail :  error get")
	}

	return &result, nil
}

func (r *userClassRepository) GetAllUserClassQuranSchedule(filter string) (*[]model.UserClassQuranSchedule, error) {
	var result []model.UserClassQuranSchedule
	query := fmt.Sprintf(_getAllUserClassQuranSchedule, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "userClassRepository.GetAllUserClassQuranSchedule :  error get")
	}

	return &result, nil
}

func (r *userClassRepository) InsertUserClass(userClass model.UserClass) (model.UserClass, bool, error) {
	var code string
	var userClassRow model.UserClass

	tx, err := r.db.Beginx()
	if err != nil {
		return userClassRow, false, errors.New("userClassRepository: InsertUserClass: error begin transaction")
	}

	err = tx.QueryRow(_insertUserClass,
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
	).Scan(&code)

	if err != nil {
		tx.Rollback()
		return userClassRow, false, utils.WrapError(err, "userClassRepository: InsertUserClass: error insert")
	}

	userClassRow = model.UserClass{Code: code}

	tx.Commit()
	return userClassRow, err == nil, nil
}

func (r *userClassRepository) InsertUserClassQuran(userClass model.UserClassQuran) (model.UserClassQuran, bool, error) {
	var code string
	var userClassRow model.UserClassQuran

	tx, err := r.db.Beginx()
	if err != nil {
		return userClassRow, false, errors.New("userClassRepository: InsertUserClassQuran: error begin transaction")
	}

	err = tx.QueryRow(_insertUserClassQuran,
		userClass.UserCode,
		userClass.ClassCode,
		userClass.CreatedBy,
		userClass.CreatedDate,
		userClass.CreatedBy,
		userClass.CreatedDate,
	).Scan(&code)

	if err != nil {
		tx.Rollback()
		return userClassRow, false, utils.WrapError(err, "userClassRepository: InsertUserClassQuran: error insert")
	}

	userClassRow = model.UserClassQuran{Code: code}

	tx.Commit()
	return userClassRow, err == nil, nil
}

func (r *userClassRepository) InsertUserClassQuranDetail(userClass model.UserClassQuran) (model.UserClassQuran, bool, error) {
	var code string
	var userClassRow model.UserClassQuran

	tx, err := r.db.Beginx()
	if err != nil {
		return userClassRow, false, errors.New("userClassRepository: InsertUserClassQuranDetail: error begin transaction")
	}

	err = tx.QueryRow(_insertUserClassQuranDetail,
		userClass.Code,
		userClass.PackageCode.String,
		userClass.CreatedBy,
		userClass.CreatedDate,
		userClass.CreatedBy,
		userClass.CreatedDate,
	).Scan(&code)

	if err != nil {
		tx.Rollback()
		return userClassRow, false, utils.WrapError(err, "userClassRepository: InsertUserClassQuranDetail: error insert")
	}

	userClassRow = model.UserClassQuran{Code: code}

	tx.Commit()
	return userClassRow, err == nil, nil
}

func (r *userClassRepository) InsertUserClassQuranSchedule(userClass model.UserClassQuranSchedule) (model.UserClassQuranSchedule, bool, error) {
	var code string
	var userClassRow model.UserClassQuranSchedule

	tx, err := r.db.Beginx()
	if err != nil {
		return userClassRow, false, errors.New("userClassRepository: InsertUserClassQuranSchedule: error begin transaction")
	}

	err = tx.QueryRow(_insertUserClassQuranSchedule,
		userClass.UserClassDetailCode,
		userClass.StartDate.Time,
		userClass.FinishDate.Time,
		userClass.MentorMessage.String,
		userClass.CreatedBy,
		userClass.CreatedDate,
		userClass.CreatedBy,
		userClass.CreatedDate,
		userClass.Material.String,
	).Scan(&code)

	if err != nil {
		tx.Rollback()
		return userClassRow, false, utils.WrapError(err, "userClassRepository: InsertUserClassQuranSchedule: error insert")
	}

	userClassRow = model.UserClassQuranSchedule{Code: code}

	tx.Commit()
	return userClassRow, err == nil, nil
}

func (r *userClassRepository) UpdateUserClassQuranSchedule(userClass model.UserClassQuranSchedule) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("userClassRepository: InsertUserClassQuranSchedule: error begin transaction")
	}

	_, err = tx.Exec(_updateUserClassQuranSchedule,
		userClass.StartDate.Time,
		userClass.FinishDate.Time,
		userClass.MentorMessage.String,
		userClass.Material.String,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "userClassRepository: InsertUserClassQuranSchedule: error insert")
	}

	tx.Commit()
	return err == nil, nil
}

func (r *userClassRepository) UpdateUserClassQuranProgress(userClass model.UserClassQuran) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("userClassRepository: UpdateUserClassQuranProgress: error begin transaction")
	}

	_, err = tx.Exec(_updateUserClassProgress,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
		userClass.Code,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "userClassRepository: UpdateUserClassQuranProgress: error update")
	}

	tx.Commit()
	return err == nil, nil
}

func (r *userClassRepository) UpdateUserClass(userClass model.UserClass) (model.UserClass, bool, error) {
	var code string
	var userClassRow model.UserClass

	tx, err := r.db.Beginx()
	if err != nil {
		return userClassRow, false, errors.New("userClassRepository: UpdateUserClass: error begin transaction")
	}

	err = tx.QueryRow(_updateUserClass,
		userClass.PackageCode,
		userClass.StartDate,
		userClass.ExpiredDate,
		userClass.TimeDuration,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
		userClass.ClassCode,
		userClass.UserCode,
	).Scan(&code)

	if err != nil {
		tx.Rollback()
		return userClassRow, false, utils.WrapError(err, "userClassRepository: UpdateUserClass: error update")
	}

	userClassRow = model.UserClass{Code: code}

	tx.Commit()
	return userClassRow, err == nil, nil
}

func (r *userClassRepository) UpdateUserClassExpired(userClass model.UserClass) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("userClassRepository: UpdateUserClassExpired: error begin transaction")
	}

	_, err = tx.Exec(_updateUserClassExpired,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
		userClass.Code,
		userClass.UserCode,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "userClassRepository: UpdateUserClassExpired: error update")
	}

	tx.Commit()
	return err == nil, nil
}

func (r *userClassRepository) UpdateUserClassTest(userClass model.UserClass, types string) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("userClassRepository: UpdateUserClassTest: error begin transaction")
	}

	if strings.ToLower(types) == "pre-test" {
		_, err = tx.Exec(_updateUserClassPreTest,
			userClass.PreTestScores.Float64,
			userClass.ModifiedBy.String,
			userClass.ModifiedDate.Time,
			userClass.Code,
		)

		if err != nil {
			tx.Rollback()
			return false, utils.WrapError(err, "userClassRepository: _updateUserClassPreTest: error update")
		}

	} else if strings.ToLower(types) == "post-test" {
		_, err = tx.Exec(_updateUserClassPostTest,
			userClass.PostTestDate.Time,
			userClass.ModifiedBy.String,
			userClass.ModifiedDate.Time,
			userClass.Code,
		)

		if err != nil {
			tx.Rollback()
			return false, utils.WrapError(err, "userClassRepository: _updateUserClassPostTest: error update")
		}
	}

	tx.Commit()
	return err == nil, nil
}

func (r *userClassRepository) DeleteUserClass(userClass model.UserClass) (time.Time, bool, error) {
	var date time.Time

	tx, err := r.db.Beginx()
	if err != nil {
		return date, false, errors.New("userClassRepository: DeleteUserClass: error begin transaction")
	}

	err = tx.QueryRow(_deleteUserClass,
		userClass.ModifiedBy.String,
		userClass.ModifiedDate.Time,
		userClass.ClassCode,
		userClass.UserCode,
	).Scan(&date)

	if err != nil {
		tx.Rollback()
		return date, false, utils.WrapError(err, "userClassRepository: DeleteUserClass: error delete")
	}

	tx.Commit()
	return date, err == nil, nil
}

func (userClassRepository *userClassRepository) CheckAllUserClassExpired() ([]model.UserClass, error) {
	var userClassList []model.UserClass

	rows, sqlError := userClassRepository.db.Query(_checkAllUserClassExpired)

	if sqlError != nil {
		utils.PushLogf("SQL error on CheckAllUserClassExpired => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var isExpired bool
			var id, timeDuration int
			var progress sql.NullFloat64
			var startDate, expiredDate sql.NullTime
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

	rows, sqlError := userClassRepository.db.Query(_checkAllUserClassBeforeExpired, interval.Interval1, interval.Interval2)

	if sqlError != nil {
		utils.PushLogf("SQL error on CheckAllUserClassExpired => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var isExpired bool
			var id, timeDuration int
			var progress sql.NullFloat64
			var startDate, expiredDate sql.NullTime
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
