package repository

import "github.com/jmoiron/sqlx"

type testRepository struct {
	db *sqlx.DB
}

type TestRepository interface{}

func InitTestRepository(db *sqlx.DB) TestRepository {
	return &testRepository{
		db,
	}
}
