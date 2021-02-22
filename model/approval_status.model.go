package model

import "database/sql"

type ApprovalStatus struct {
	ID                  int
	Code                string
	CurrentStatus       string
	CurrentStatusValue  string
	ApprovedStatus      sql.NullString
	ApprovedStatusValue sql.NullString
	RejectStatus        sql.NullString
	RejectStatusValue   sql.NullString
	ReviseStatus        sql.NullString
	ReviseStatusValue   sql.NullString
}
