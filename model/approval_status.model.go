package model

import "database/sql"

type ApprovalStatus struct {
	ID             int
	Code           string
	CurrentStatus  string
	ApprovedStatus sql.NullString
	RejectStatus   sql.NullString
}
