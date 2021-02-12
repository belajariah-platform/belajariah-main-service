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
		approve_status,
		reject_status
	FROM 
		master_approval_status
	WHERE
		current_status = $1
	;
	`, statusCode)
	var id int
	var code, currentStatus string
	var approvedStatus, rejectStatus sql.NullString

	err := row.Scan(
		&id,
		&code,
		&currentStatus,
		&approvedStatus,
		&rejectStatus,
	)
	if err != nil {
		utils.PushLogf("SQL error on GetApprovalStatusByID => ", err)
		return model.ApprovalStatus{}, nil
	} else {
		approvalStatusRow = model.ApprovalStatus{
			ID:             id,
			Code:           code,
			CurrentStatus:  currentStatus,
			ApprovedStatus: approvedStatus,
			RejectStatus:   rejectStatus,
		}
		return approvalStatusRow, err
	}
}
