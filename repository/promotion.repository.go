package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	_getPromotion = `
		SELECT
			id,
			code,
			class_code,
			title,
			description,
			promo_code,
			promo_type_code,
			promo_type,
			discount,
			image_banner,
			image_header,
			expired_date,
			quota_user,
			quota_used,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM 
			master.v_m_promotion
		WHERE 
			is_deleted=false AND
			is_active=true AND
			(code = $1 OR promo_code = $1) 
	`
	_getAllPromotion = `
		SELECT
			id,
			code,
			class_code,
			title,
			description,
			promo_code,
			promo_type_code,
			promo_type,
			discount,
			image_banner,
			image_header,
			expired_date,
			quota_user,
			quota_used,
			is_active,
			created_by,
			created_date,
			modified_by,
			modified_date,
			is_deleted
		FROM 
			master.v_m_promotion
		WHERE 
			is_deleted=false AND
			is_active=true
		%s
		OFFSET %d
	LIMIT %d
	`
	_getAllPromotionCount = `
		SELECT COUNT(*) FROM 
			master.v_m_promotion  
		WHERE 
			is_deleted=false AND
			is_active=true
		%s
	`
	_checkAllPromotionExpired = `
		SELECT
			code
		FROM 
			master.v_m_promotion
		WHERE  
			deleted_by IS NULL AND
			is_active=true AND
			promo_code <> 'BLJEXP' AND
			expired_date <= now() OR 
			expired_date = now() OR 
			quota_used >= quota_user
	`
	_updatePromotionActivated = `
		UPDATE
			master.master_promotion
		SET
			is_active=false,
			modified_by=$1,
			modified_date=$2
		WHERE
			code=$3
	`
)

type promotionRepository struct {
	db *sqlx.DB
}

type PromotionRepository interface {
	GetPromotion(filter string) (model.Promotion, error)

	GetAllPromotion(skip, take int, filter string) ([]model.Promotion, error)
	GetAllPromotionCount(filter string) (int, error)

	CheckAllPromotionExpired() ([]model.Promotion, error)
	UpdatePromotionActivated(payment model.Promotion) (bool, error)
}

func InitPromotionRepository(db *sqlx.DB) PromotionRepository {
	return &promotionRepository{
		db,
	}
}

func (promotionRepository *promotionRepository) GetAllPromotion(skip, take int, filter string) ([]model.Promotion, error) {
	var promotionList []model.Promotion
	query := fmt.Sprintf(_getAllPromotion, filter, skip, take)

	rows, sqlError := promotionRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPromotion => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var createdDate time.Time
			var isActive, isDeleted bool
			var discount sql.NullFloat64
			var quoteUser, quotaUsed sql.NullInt64
			var modifiedDate, expiredDate sql.NullTime
			var code, classCode, title, promoCode, createdBy string
			var promoType, promoTypeCode, bannerImage, headerImage, description, modifiedBy sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&classCode,
				&title,
				&description,
				&promoCode,
				&promoTypeCode,
				&promoType,
				&discount,
				&bannerImage,
				&headerImage,
				&expiredDate,
				&quoteUser,
				&quotaUsed,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&isDeleted,
			)
			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllPromotion => ", sqlError.Error())
			} else {
				promotionList = append(
					promotionList,
					model.Promotion{
						ID:            id,
						Code:          code,
						ClassCode:     classCode,
						Title:         title,
						Description:   description,
						PromoCode:     promoCode,
						PromoTypeCode: promoTypeCode,
						PromoType:     promoType,
						Discount:      discount,
						ImageBanner:   bannerImage,
						ImageHeader:   headerImage,
						ExpiredDate:   expiredDate,
						QuotaUser:     quoteUser,
						QuotaUsed:     quotaUsed,
						IsActive:      isActive,
						CreatedBy:     createdBy,
						CreatedDate:   createdDate,
						ModifiedBy:    modifiedBy,
						ModifiedDate:  modifiedDate,
						IsDeleted:     isDeleted,
					},
				)
			}
		}
	}
	return promotionList, sqlError
}

func (promotionRepository *promotionRepository) GetPromotion(filter string) (model.Promotion, error) {
	var promotionRow model.Promotion
	row := promotionRepository.db.QueryRow(_getPromotion, filter)

	var id int
	var createdDate time.Time
	var isActive, isDeleted bool
	var discount sql.NullFloat64
	var quoteUser, quotaUsed sql.NullInt64
	var modifiedDate, expiredDate sql.NullTime
	var code, classCode, title, promoCode, createdBy string
	var promoType, promoTypeCode, bannerImage, headerImage, description, modifiedBy sql.NullString

	sqlError := row.Scan(
		&id,
		&code,
		&classCode,
		&title,
		&description,
		&promoCode,
		&promoTypeCode,
		&promoType,
		&discount,
		&bannerImage,
		&headerImage,
		&expiredDate,
		&quoteUser,
		&quotaUsed,
		&isActive,
		&createdBy,
		&createdDate,
		&modifiedBy,
		&modifiedDate,
		&isDeleted,
	)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetPromotion => ", sqlError.Error())
		return model.Promotion{}, nil
	} else {
		promotionRow = model.Promotion{
			ID:            id,
			Code:          code,
			ClassCode:     classCode,
			Title:         title,
			Description:   description,
			PromoCode:     promoCode,
			PromoTypeCode: promoTypeCode,
			PromoType:     promoType,
			Discount:      discount,
			ImageBanner:   bannerImage,
			ImageHeader:   headerImage,
			ExpiredDate:   expiredDate,
			QuotaUser:     quoteUser,
			QuotaUsed:     quotaUsed,
			IsActive:      isActive,
			CreatedBy:     createdBy,
			CreatedDate:   createdDate,
			ModifiedBy:    modifiedBy,
			ModifiedDate:  modifiedDate,
			IsDeleted:     isDeleted,
		}
		return promotionRow, sqlError
	}
}

func (promotionRepository *promotionRepository) GetAllPromotionCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(_getAllPromotionCount, filter)

	row := promotionRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPromotionCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (r *promotionRepository) UpdatePromotionActivated(promotion model.Promotion) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("promotionRepository: UpdatePromotionActivated: error begin transaction")
	}

	_, err = tx.Exec(_updatePromotionActivated,
		promotion.ModifiedBy.String,
		promotion.ModifiedDate.Time,
		promotion.Code,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "promotionRepository: UpdatePromotionActivated: error update")
	}

	tx.Commit()
	return err == nil, nil
}

func (promotionRepository *promotionRepository) CheckAllPromotionExpired() ([]model.Promotion, error) {
	var promotionList []model.Promotion
	var code string

	rows, sqlError := promotionRepository.db.Query(_checkAllPromotionExpired)
	if sqlError != nil {
		utils.PushLogf("SQL error on CheckAllPromotionExpired => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			sqlError := rows.Scan(&code)
			if sqlError != nil {
				utils.PushLogf("SQL error on CheckAllPromotionExpired => ", sqlError.Error())
			} else {
				promotionList = append(promotionList, model.Promotion{
					Code: code,
				})
			}
		}
	}
	return promotionList, sqlError
}
