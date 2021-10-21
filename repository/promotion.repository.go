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
	query := fmt.Sprintf(`
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
	`, filter, skip, take)

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
	row := promotionRepository.db.QueryRow(`
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
		(code = $1 OR promo_code = $1) `, filter)

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
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		master.v_m_promotion  
	WHERE 
		is_deleted=false AND
		is_active=true
	%s
	`, filter)

	row := promotionRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllPromotionCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}

func (promotionRepository *promotionRepository) UpdatePromotionActivated(promotion model.Promotion) (bool, error) {
	var err error
	var result bool

	tx, errTx := promotionRepository.db.Begin()
	if errTx != nil {
		utils.PushLogf("error in UpdatePromotionActivated", errTx)
	} else {
		err = updatePromotionActivated(tx, promotion)
		if err != nil {
			utils.PushLogf("err in promotion---", err.Error())
		}
	}

	if err == nil {
		result = true
		tx.Commit()
	} else {
		result = false
		tx.Rollback()
		utils.PushLogf("failed to updatePromotionActivated", err.Error())
	}

	return result, err
}

func updatePromotionActivated(tx *sql.Tx, promotion model.Promotion) error {
	_, err := tx.Exec(`
	UPDATE
		master_promotion
	 SET
	 	is_active=false,
		modified_by=$1,
		modified_date=$2
 	WHERE
 		id=$3
	`,
		promotion.ModifiedBy.String,
		promotion.ModifiedDate.Time,
		promotion.ID,
	)
	return err
}

func (promotionRepository *promotionRepository) CheckAllPromotionExpired() ([]model.Promotion, error) {
	var promotionList []model.Promotion
	rows, sqlError := promotionRepository.db.Query(`
	SELECT
		id
	FROM 
		v_m_promotion
	WHERE  
		deleted_by IS NULL AND
		is_active=true AND
		promo_code <> 'BLJEXP' AND
		expired_date <= now() OR 
		expired_date = now() OR 
		quota_used >= quota_user
	;
	`)
	if sqlError != nil {
		utils.PushLogf("SQL error on CheckAllPromotionExpired => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int

			sqlError := rows.Scan(&id)
			if sqlError != nil {
				utils.PushLogf("SQL error on CheckAllPromotionExpired => ", sqlError.Error())
			} else {
				promotionList = append(promotionList, model.Promotion{
					ID: id,
				})
			}
		}
	}
	return promotionList, sqlError
}
