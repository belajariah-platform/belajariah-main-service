package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	_getAllMasterSql = `
		SELECT 
			id, 
			code, 
			title,
			description,
			expired_date,
			quota_user,
			image_header,
			image_banner,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM
			master.master_coaching_program
		%s
	`
	_getAllCpSql = `
		SELECT 
			id, 
			code, 
			cp_code,
			program_description,
			user_code,
			fullname,
			gender,
			email,
			wa_no,
			age,
			address,
			profession,
			question_1,
			question_2,
			question_3,
			question_4,
			question_5,
			question_6,
			question_7,
			question_8,
			is_confirmed,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM
			transaction.v_t_coaching_program
		%s
	`
	_insertCpSql = `
		INSERT INTO transaction.transact_coaching_program
		(
			cp_code,
			user_code,
			fullname,
			gender,
			email,
			wa_no,
			age,
			address,
			profession,
			question_1,
			question_2,
			question_3,
			question_4,
			question_5,
			question_6,
			question_7,
			question_8,
			created_by,
			created_date,
			modified_by,
			modified_date
		)
		VALUES (
			:cp_code,
			:user_code,
			:fullname,
			:gender,
			:email,
			:wa_no,
			:age,
			:address,
			:profession,
			:question_1,
			:question_2,
			:question_3,
			:question_4,
			:question_5,
			:question_6,
			:question_7,
			:question_8,
			:created_by,
			:created_date,
			:modified_by,
			:modified_date
			);
	`
	_confirmCpSql = `
		UPDATE
			transaction.transact_coaching_program
		SET
			is_confirmed=:is_confirmed,
			modified_by=:modified_by,
			modified_date=:modified_date
		WHERE
			user_code=:user_code and email=:email
	`
)

type coachingProgramRepository struct {
	db *sqlx.DB
}

type CoachingProgramRepository interface {
	GetAllMasterCoachingProgram(filter string) (*[]model.MasterCoachingProgram, error)
	GetAllCoachingProgram(filter string) (*[]model.CoachingProgram, error)

	InsertCoachingProgram(data model.CoachingProgram) (bool, error)
	ConfirmCoachingProgram(cp model.CoachingProgram) (bool, error)
}

func InitCoachingProgramRepository(db *sqlx.DB) CoachingProgramRepository {
	return &coachingProgramRepository{
		db,
	}
}

func (r *coachingProgramRepository) GetAllMasterCoachingProgram(filter string) (*[]model.MasterCoachingProgram, error) {
	var result []model.MasterCoachingProgram
	query := fmt.Sprintf(_getAllMasterSql, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "coachingProgramRepository.GetAllMasterCoachingPorgram :  error get")
	}

	return &result, nil
}

func (r *coachingProgramRepository) GetAllCoachingProgram(filter string) (*[]model.CoachingProgram, error) {
	var result []model.CoachingProgram
	query := fmt.Sprintf(_getAllCpSql, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "coachingProgramRepository.GetAllCoachingPorgram : ")
	}

	return &result, nil
}

func (r *coachingProgramRepository) InsertCoachingProgram(data model.CoachingProgram) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("coachingProgramRepository.InsertCoachingProgram: error begin transaction")
	}

	_, err = tx.NamedExec(_insertCpSql, data)
	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "coachingProgramRepository: InsertProgressTaskHeader: error insert")
	}

	tx.Commit()
	return err == nil, nil
}

func (r *coachingProgramRepository) ConfirmCoachingProgram(data model.CoachingProgram) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("coachingProgramRepository.ConfirmCoachingProgram: error begin transaction")
	}

	_, err = tx.NamedExec(_confirmCpSql, data)
	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "coachingProgramRepository.ConfirmCoachingProgram: error update")
	}

	tx.Commit()
	return err == nil, nil
}
