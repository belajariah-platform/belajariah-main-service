package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type approvalStatusRepository struct {
	db *sqlx.DB
}

type ApprovalStatusRepository interface {
	GetApprovalStatus(statusCode string) (model.ApprovalStatus, error)
}

func InitApprovalStatusRepository(db *sqlx.DB) ApprovalStatusRepository {
	return &approvalStatusRepository{
		db,
	}
}

func (approvalStatusRepository *approvalStatusRepository) GetApprovalStatus(statusCode string) (model.ApprovalStatus, error) {
	var approvalStatusRow model.ApprovalStatus
	row := approvalStatusRepository.db.QueryRow(`
	SELECT 
		id,
		code, 
		current_status,
		current_status_value,
		approve_status,
		approve_status_value,
		reject_status,
		reject_status_value,
		revise_status,
		revise_status_value
	FROM 
		v_m_approval_status
	WHERE
		current_status = $1
	;
	`, statusCode)
	var id int
	var code, currentStatus, currentStatusValue string
	var approvedStatus, approvedStatusValue, rejectStatus, rejectStatusValue, reviseStatus, reviseStatusValue sql.NullString

	err := row.Scan(
		&id,
		&code,
		&currentStatus,
		&currentStatusValue,
		&approvedStatus,
		&approvedStatusValue,
		&rejectStatus,
		&rejectStatusValue,
		&reviseStatus,
		&reviseStatusValue,
	)
	if err != nil {
		utils.PushLogf("SQL error on GetApprovalStatusByID => ", err)
		return model.ApprovalStatus{}, nil
	} else {
		approvalStatusRow = model.ApprovalStatus{
			ID:                  id,
			Code:                code,
			CurrentStatus:       currentStatus,
			CurrentStatusValue:  currentStatusValue,
			ApprovedStatus:      approvedStatus,
			ApprovedStatusValue: approvedStatusValue,
			RejectStatus:        rejectStatus,
			RejectStatusValue:   rejectStatusValue,
			ReviseStatus:        reviseStatus,
			ReviseStatusValue:   reviseStatusValue,
		}
		return approvalStatusRow, err
	}
}
