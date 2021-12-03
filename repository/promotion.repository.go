package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	_getAllPromotion = `
		SELECT
			id,
			code,
			class_code,
			title,
			coalesce(description, '') as description,
			promo_code,
			coalesce(promo_type_code, '') as promo_type_code,
			coalesce(promo_type, '') as promo_type,
			coalesce(discount, 0) as discount,
			coalesce(image_banner, '') as image_banner,
			coalesce(image_header, '') as image_header,
			coalesce(expired_date, to_timestamp(0)) as expired_date,
			coalesce(quota_user, 0) as quota_user,
			coalesce(quota_used, 0) as quota_used,
			is_active,
			created_by,
			created_date,
			coalesce(modified_by, '') as modified_by,
			coalesce(modified_date, to_timestamp(0)) as modified_date,
			is_deleted
		FROM 
			master.v_m_promotion
		%s
	`
	_getAllPromotionHeader = `
		SELECT
			id,
			code,
			title,
			coalesce(description, '') as description,
			promo_code,
			coalesce(promo_type_code, '') as promo_type_code,
			coalesce(promo_type, '') as promo_type,
			coalesce(discount, 0) as discount,
			coalesce(image_banner, '') as image_banner,
			coalesce(image_header, '') as image_header,
			coalesce(expired_date, to_timestamp(0)) as expired_date,
			coalesce(quota_user, 0) as quota_user,
			coalesce(quota_used, 0) as quota_used,
			is_active,
			created_by,
			created_date,
			coalesce(modified_by, '') as modified_by,
			coalesce(modified_date, to_timestamp(0)) as modified_date,
			is_deleted
		FROM 
			master.v_m_promotion_header
		%s
	`
	_getAllPromotionHeaderCount = `
		SELECT COUNT(*) FROM 
			master.v_m_promotion_header  
		%s
	`
	_checkAllPromotionExpired = `
		SELECT
			code
		FROM 
			master.v_m_promotion
		WHERE  
			is_deleted = false AND
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
	GetAllPromotions(filter string) (*[]model.Promotion, error)
	GetAllPromotionHeader(filter string) (*[]model.Promotion, error)
	GetAllPromotionHeaderCount(filter string) (int, error)

	CheckAllPromotionExpired() ([]model.Promotion, error)
	UpdatePromotionActivated(payment model.Promotion) (bool, error)
}

func InitPromotionRepository(db *sqlx.DB) PromotionRepository {
	return &promotionRepository{
		db,
	}
}

func (r *promotionRepository) GetAllPromotions(filter string) (*[]model.Promotion, error) {
	var result []model.Promotion
	query := fmt.Sprintf(_getAllPromotion, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "promotionRepository.GetAllPromotions :  error get")
	}

	return &result, nil
}

func (r *promotionRepository) GetAllPromotionHeader(filter string) (*[]model.Promotion, error) {
	var result []model.Promotion
	query := fmt.Sprintf(_getAllPromotionHeader, filter)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, utils.WrapError(err, "promotionRepository.GetAllPromotionHeader :  error get")
	}

	return &result, nil
}

func (r *promotionRepository) GetAllPromotionHeaderCount(filter string) (int, error) {
	var count int

	query := fmt.Sprintf(_getAllPromotionHeaderCount, filter)

	row := r.db.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return 0, utils.WrapError(err, "promotionRepository: GetCount: error query row")
	}

	return count, err
}

func (r *promotionRepository) UpdatePromotionActivated(promotion model.Promotion) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, errors.New("promotionRepository: UpdatePromotionActivated: error begin transaction")
	}

	_, err = tx.Exec(_updatePromotionActivated,
		promotion.ModifiedBy,
		promotion.ModifiedDate,
		promotion.Code,
	)

	if err != nil {
		tx.Rollback()
		return false, utils.WrapError(err, "promotionRepository: UpdatePromotionActivated: error update")
	}

	tx.Commit()
	return err == nil, nil
}

func (r *promotionRepository) CheckAllPromotionExpired() ([]model.Promotion, error) {
	var promotionList []model.Promotion
	var code string

	rows, err := r.db.Query(_checkAllPromotionExpired)
	if err != nil {
		utils.PushLogf("promotionRepository.CheckAllPromotionExpired : err query ", err.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&code)
			if err != nil {
				utils.PushLogf("promotionRepository.CheckAllPromotionExpired : err scan ", err.Error())
			} else {
				promotionList = append(promotionList, model.Promotion{
					Code: code,
				})
			}
		}
	}
	return promotionList, err
}
