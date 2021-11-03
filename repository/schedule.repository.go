package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	_getAllScheduleSql = `
		SELECT 
			id,
			code,
			user_class_history_code,
			class_code,
			class_name,
			user_code,
			user_name,
			mentor_code,
			mentor_name,
			description,
			shift_name,
			planning_start_time,
			planning_end_time,
			actual_user_start_date,
			actual_user_end_date,
			actual_mentor_start_date,
			actual_mentor_end_date,
			sequence,
			is_completed_user,
			is_completed_mentor,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM
			transaction.v_t_schedule
		%s
	`
	_getAllMasterSchedule = `
	SELECT
		id,
		code,
		mentor_code,
		shift_name,
		start_date AS planning_start_time,
		end_date AS planning_end_time,
		coalesce(sequence, 0) AS sequence,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM 
		master.master_mentor_schedule
	WHERE 
		is_active=true and 
		is_deleted=false
		%s'
	`
	_insertScheduleSql = `
		INSERT INTO transaction.transact_schedule
		(
			user_class_history_code,
			class_code,
			user_code,
			mentor_code,
			description,
			shift_name,
			planning_start_time,
			planning_end_time,
			actual_user_start_date,
			actual_user_end_date,
			actual_mentor_start_date,
			actual_mentor_end_date,
			sequence,
			is_completed_user,
			is_completed_mentor,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date
		)
		VALUES (
			:user_class_history_code,
			:class_code,
			:user_code,
			:mentor_code,
			:description,
			:shift_name,
			:planning_start_time,
			:planning_end_time,
			:actual_user_start_date,
			:actual_user_end_date,
			:actual_mentor_start_date,
			:actual_mentor_end_date,
			:sequence,
			:is_completed_user,
			:is_completed_mentor,
			:is_active,
			:created_by,
			:created_date,
			:modified_by,
			:modified_date
			);
	`
	_updateScheduleUserSql = `
		UPDATE
			transaction.transact_schedule
		SET
			actual_user_start_date=:actual_user_start_date,
			actual_user_end_date=:actual_user_end_date,
			is_completed_user=:is_completed_user,
			modified_by=:modified_by,
			modified_date=:modified_date
		WHERE
			code=:code
`
	_updateScheduleMentorSql = `
		UPDATE
			transaction.transact_schedule
		SET
			actual_mentor_start_date=:actual_mentor_start_date,
			actual_mentor_end_date=:actual_mentor_end_date
			is_completed_mentor=:is_completed_mentor,
			modified_by=:modified_by,
			modified_date=:modified_date
		WHERE
			code=:code
		`
)

type scheduleRepository struct {
	db *sqlx.DB
}

type ScheduleRepository interface {
	GetAllSchedule(filter string) (*[]model.Schedule, error)
	GetAllMasterSchedule(filter string) (*[]model.Schedule, error)

	InsertSchedule(data model.Schedule) (bool, error)
	UpdateScheduleUser(data model.Schedule) (bool, error)
	UpdateScheduleMentor(data model.Schedule) (bool, error)
}

func InitScheduleRepository(db *sqlx.DB) ScheduleRepository {
	return &scheduleRepository{
		db,
	}
}

func (r *scheduleRepository) GetAllSchedule(filter string) (*[]model.Schedule, error) {
	var result []model.Schedule
	query := fmt.Sprintf(_getAllScheduleSql, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "scheduleRepository.GetAllSchedule :  error get")
	}

	return &result, nil
}

func (r *scheduleRepository) GetAllMasterSchedule(filter string) (*[]model.Schedule, error) {
	var result []model.Schedule
	query := fmt.Sprintf(_getAllScheduleSql, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "scheduleRepository.GetAllMasterSchedule :  error get")
	}

	return &result, nil
}

func (r *scheduleRepository) InsertSchedule(data model.Schedule) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("scheduleRepository.InsertSchedule: error begin transaction")
	}

	_, err = tx.NamedExec(_insertScheduleSql, data)
	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "scheduleRepository: InsertSchedule: error insert")
	}

	tx.Commit()
	return err == nil, nil
}

func (r *scheduleRepository) UpdateScheduleUser(data model.Schedule) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("scheduleRepository.UpdateScheduleUser: error begin transaction")
	}

	_, err = tx.NamedExec(_updateScheduleUserSql, data)
	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "scheduleRepository: UpdateScheduleUser: error update")
	}

	tx.Commit()
	return err == nil, nil
}

func (r *scheduleRepository) UpdateScheduleMentor(data model.Schedule) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("scheduleRepository.UpdateScheduleMentor: error begin transaction")
	}

	_, err = tx.NamedExec(_updateScheduleMentorSql, data)
	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "scheduleRepository: UpdateScheduleMentor: error update")
	}

	tx.Commit()
	return err == nil, nil
}
