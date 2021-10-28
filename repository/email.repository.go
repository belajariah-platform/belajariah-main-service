package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	_getAllEmailSql = `
		SELECT 
			id, 
			code, 
			type,
			body,
			subject,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM
			master.master_email
		WHERE type = '%s'
	`
)

type emailRepository struct {
	db *sqlx.DB
}

type EmailRepository interface {
	GetAllEmail(filter string) (*[]model.Email, error)
}

func InitEmailRepository(db *sqlx.DB) EmailRepository {
	return &emailRepository{
		db,
	}
}

func (r *emailRepository) GetAllEmail(filter string) (*[]model.Email, error) {
	var result []model.Email
	query := fmt.Sprintf(_getAllEmailSql, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "emailRepository.GetAllEmail :  error get")
	}

	return &result, nil
}
