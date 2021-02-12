package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type promotionRepository struct {
	db *sqlx.DB
}

type PromotionRepository interface {
	GetAllPromotion(skip, take int, filter string) ([]model.Promotion, error)
	GetPromotion(filter string) (model.Promotion, error)
	GetAllPromotionCount(filter string) (int, error)
}

func InitPromotionRepository(db *sqlx.DB) PromotionRepository {
	return &promotionRepository{
		db,
	}
}

func (promotionRepository *promotionRepository) GetAllPromotion(skip, take int, filter string) ([]model.Promotion, error) {
	var promotionList []model.Promotion
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		class_code,
		title,
		description,
		promo_code,
		discount,
		banner_image,
		header_image,
		expired_date,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM 
		v_m_promotion
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := promotionRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPromotion => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var isActive bool
			var createdDate time.Time
			var discount sql.NullFloat64
			var modifiedDate, expiredDate, deletedDate sql.NullTime
			var code, classCode, title, promoCode, createdBy string
			var bannerImage, headerImage, description, modifiedBy, deletedBy sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&classCode,
				&title,
				&description,
				&promoCode,
				&discount,
				&bannerImage,
				&headerImage,
				&expiredDate,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
			)
			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllPromotion => ", sqlError)
			} else {
				promotionList = append(
					promotionList,
					model.Promotion{
						ID:           id,
						Code:         code,
						ClassCode:    classCode,
						Title:        title,
						Description:  description,
						PromoCode:    promoCode,
						Discount:     discount,
						BannerImage:  bannerImage,
						HeaderImage:  headerImage,
						ExpiredDate:  expiredDate,
						IsActive:     isActive,
						CreatedBy:    createdBy,
						CreatedDate:  createdDate,
						ModifiedBy:   modifiedBy,
						ModifiedDate: modifiedDate,
						DeletedBy:    deletedBy,
						DeletedDate:  deletedDate,
					},
				)
			}
		}
	}
	return promotionList, sqlError
}

func (promotionRepository *promotionRepository) GetPromotion(filter string) (model.Promotion, error) {
	var promotionRow model.Promotion
	row := promotionRepository.db.QueryRow(`
	SELECT
		id,
		code,
		class_code,
		title,
		description,
		promo_code,
		discount,
		banner_image,
		header_image,
		expired_date,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM 
		v_m_promotion
	WHERE 
		deleted_by IS NULL AND
		is_active=true AND
		(code = $1 OR promo_code = $1) `, filter)

	var id int
	var isActive bool
	var createdDate time.Time
	var discount sql.NullFloat64
	var modifiedDate, expiredDate, deletedDate sql.NullTime
	var code, classCode, title, promoCode, createdBy string
	var bannerImage, headerImage, description, modifiedBy, deletedBy sql.NullString

	sqlError := row.Scan(
		&id,
		&code,
		&classCode,
		&title,
		&description,
		&promoCode,
		&discount,
		&bannerImage,
		&headerImage,
		&expiredDate,
		&isActive,
		&createdBy,
		&createdDate,
		&modifiedBy,
		&modifiedDate,
		&deletedBy,
		&deletedDate,
	)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetPromotion => ", sqlError)
		return model.Promotion{}, nil
	} else {
		promotionRow = model.Promotion{
			ID:           id,
			Code:         code,
			ClassCode:    classCode,
			Title:        title,
			Description:  description,
			PromoCode:    promoCode,
			Discount:     discount,
			BannerImage:  bannerImage,
			HeaderImage:  headerImage,
			ExpiredDate:  expiredDate,
			IsActive:     isActive,
			CreatedBy:    createdBy,
			CreatedDate:  createdDate,
			ModifiedBy:   modifiedBy,
			ModifiedDate: modifiedDate,
			DeletedBy:    deletedBy,
			DeletedDate:  deletedDate,
		}
		return promotionRow, sqlError
	}
}

func (promotionRepository *promotionRepository) GetAllPromotionCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		v_m_promotion  
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	`, filter)

	row := promotionRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPromotionCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}
