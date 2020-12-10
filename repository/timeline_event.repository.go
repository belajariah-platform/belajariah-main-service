package repository

import "github.com/jmoiron/sqlx"

type eventRepository struct {
	db *sqlx.DB
}

type EventRepository interface{}

func InitEventRepository(db *sqlx.DB) EventRepository {
	return &eventRepository{
		db,
	}
}
