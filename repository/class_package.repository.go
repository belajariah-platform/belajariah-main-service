package repository

import "github.com/jmoiron/sqlx"

type packageRepository struct {
	db *sqlx.DB
}

type PackageRepository interface{}

func InitPackageRepository(db *sqlx.DB) PackageRepository {
	return &packageRepository{
		db,
	}
}
