package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	_getAllEventSql = `
	SELECT id,
		code,
		event_name,
		event_type,
		event_type_desc,
		event_image,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM master.v_m_event
	%s
`
	_getAllEventMappingFormSql = `
	SELECT id,
		code,
		event_code,
		event_name,
		event_type,
		event_form_code,
		question,
		question_type,
		choice,
		is_required,
		sort,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM master.v_m_event_mapping_form
	%s
`
	_insertFormClassIntensSql = `
	INSERT INTO "transaction".transact_class_intens_registrar 
	(
		event_code,
		event_form_code,
		user_code,
		question,
		answer,
		created_by,
		created_date,
		modified_by,
		modified_date
	)
	VALUES (
		:event_code,
		:event_form_code,
		:user_code,
		:question,
		:answer,
		:created_by,
		:created_date,
		:modified_by,
		:modified_date
		)
	`
)

type eventRepository struct {
	db *sqlx.DB
}

type EventRepository interface {
	GetAllEvent(filter string) (*[]model.Event, error)
	GetAllEventMappingForm(filter string) (*[]model.EventMappingForm, error)
	InsertFormClassIntens(data model.EventMappingForm) (bool, error)
}

func InitEventRepository(db *sqlx.DB) EventRepository {
	return &eventRepository{
		db,
	}
}

func (r *eventRepository) GetAllEvent(filter string) (*[]model.Event, error) {
	var result []model.Event
	query := fmt.Sprintf(_getAllEventSql, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "eventRepository.GetAllEvent :  error get")
	}

	return &result, nil
}

func (r *eventRepository) GetAllEventMappingForm(filter string) (*[]model.EventMappingForm, error) {
	var result []model.EventMappingForm
	query := fmt.Sprintf(_getAllEventMappingFormSql, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "eventRepository.GetAllEventMappingForm :  error get")
	}

	return &result, nil
}

func (r *eventRepository) InsertFormClassIntens(data model.EventMappingForm) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("eventRepository: InsertFormClassIntens: error begin transaction")
	}

	_, err = tx.NamedExec(_insertFormClassIntensSql, data)
	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "eventRepository: InsertFormClassIntens: error insert")
	}

	tx.Commit()
	return err == nil, nil
}
