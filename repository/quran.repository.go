package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type quranRepository struct {
	db *sqlx.DB
}

type QuranRepository interface {
	GetAllQuran(skip, take int, filter string) ([]model.Quran, error)
	GetAllQuranCount(filter string) (int, error)
	GetAllAyatQuran(skip, take int, filter string) ([]model.Quran, error)
	GetAllAyatQuranCount(filter string) (int, error)
	GetAllQuranView(skip, take int, filter string) ([]model.Quran, error)
	GetAllQuranViewCount(filter string) (int, error)
}

func InitQuranRepository(db *sqlx.DB) QuranRepository {
	return &quranRepository{
		db,
	}
}

func (quranRepository *quranRepository) GetAllQuran(skip, take int, filter string) ([]model.Quran, error) {
	var quranList []model.Quran
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		surat_name,
		surat_text,
		surat_translate,
		count_ayat,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM 
		v_m_alquran
	WHERE 
		deleted_by IS NULL
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := quranRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllQuran => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var isActive bool
			var id, countAyat int
			var createdDate time.Time
			var modifiedDate, deletedDate sql.NullTime
			var modifiedBy, deletedBy sql.NullString
			var code, suratName, suratText, suratTranslate, createdBy string

			sqlError := rows.Scan(
				&id,
				&code,
				&suratName,
				&suratText,
				&suratTranslate,
				&countAyat,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllQuran => ", sqlError)
			} else {
				quranList = append(
					quranList,
					model.Quran{
						ID:             id,
						Code:           code,
						SuratName:      suratName,
						SuratText:      suratText,
						SuratTranslate: suratTranslate,
						CountAyat:      countAyat,
						IsActive:       isActive,
						CreatedBy:      createdBy,
						CreatedDate:    createdDate,
						ModifiedBy:     modifiedBy,
						ModifiedDate:   modifiedDate,
						DeletedBy:      deletedBy,
						DeletedDate:    deletedDate,
					},
				)
			}
		}
	}
	return quranList, sqlError
}

func (quranRepository *quranRepository) GetAllQuranCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		v_m_alquran  
	WHERE 
		deleted_by IS NULL
	%s
	`, filter)

	row := quranRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllQuranCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}

func (quranRepository *quranRepository) GetAllAyatQuran(skip, take int, filter string) ([]model.Quran, error) {
	var quranList []model.Quran
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		ayat_number,
		ayat_text,
		ayat_translate,
		juz_number,
		page_number,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM 
		master_alquran
	WHERE 
		deleted_by IS NULL
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := quranRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllAyatQuran => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var isActive bool
			var createdDate time.Time
			var modifiedBy, deletedBy sql.NullString
			var modifiedDate, deletedDate sql.NullTime
			var id, juzNUmber, pageNumber, ayatNumber int
			var code, ayatText, ayatTranslate, createdBy string

			sqlError := rows.Scan(
				&id,
				&code,
				&ayatNumber,
				&ayatText,
				&ayatTranslate,
				&juzNUmber,
				&pageNumber,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllAyatQuran => ", sqlError)
			} else {
				quranList = append(
					quranList,
					model.Quran{
						ID:            id,
						Code:          code,
						AyatNumber:    ayatNumber,
						AyatText:      ayatText,
						AyatTranslate: ayatTranslate,
						JuzNumber:     juzNUmber,
						PageNumber:    pageNumber,
						IsActive:      isActive,
						CreatedBy:     createdBy,
						CreatedDate:   createdDate,
						ModifiedBy:    modifiedBy,
						ModifiedDate:  modifiedDate,
						DeletedBy:     deletedBy,
						DeletedDate:   deletedDate,
					},
				)
			}
		}
	}
	return quranList, sqlError
}

func (quranRepository *quranRepository) GetAllAyatQuranCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		master_alquran  
	WHERE 
		deleted_by IS NULL
	%s
	`, filter)

	row := quranRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllAyatQuranCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}

func (quranRepository *quranRepository) GetAllQuranView(skip, take int, filter string) ([]model.Quran, error) {
	var quranList []model.Quran
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		surat_name,
		surat_text,
		surat_translate,
		count_ayat,
		ayat_number,
		ayat_text,
		ayat_translate,
		juz_number,
		page_number,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM 
		master_alquran
	WHERE 
		deleted_by IS NULL
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := quranRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllQuranView => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var isActive bool
			var createdDate time.Time
			var modifiedBy, deletedBy sql.NullString
			var modifiedDate, deletedDate sql.NullTime
			var id, juzNUmber, pageNumber, ayatNumber, countAyat int
			var code, suratName, suratText, suratTranslate, ayatText, ayatTranslate, createdBy string

			sqlError := rows.Scan(
				&id,
				&code,
				&suratName,
				&suratText,
				&suratTranslate,
				&countAyat,
				&ayatNumber,
				&ayatText,
				&ayatTranslate,
				&juzNUmber,
				&pageNumber,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllQuranView => ", sqlError)
			} else {
				quranList = append(
					quranList,
					model.Quran{
						ID:             id,
						Code:           code,
						SuratName:      suratName,
						SuratText:      suratText,
						SuratTranslate: suratTranslate,
						CountAyat:      countAyat,
						AyatNumber:     ayatNumber,
						AyatText:       ayatText,
						AyatTranslate:  ayatTranslate,
						JuzNumber:      juzNUmber,
						PageNumber:     pageNumber,
						IsActive:       isActive,
						CreatedBy:      createdBy,
						CreatedDate:    createdDate,
						ModifiedBy:     modifiedBy,
						ModifiedDate:   modifiedDate,
						DeletedBy:      deletedBy,
						DeletedDate:    deletedDate,
					},
				)
			}
		}
	}
	return quranList, sqlError
}

func (quranRepository *quranRepository) GetAllQuranViewCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		master_alquran  
	WHERE 
		deleted_by IS NULL
	%s
	`, filter)

	row := quranRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllQuranViewCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}
