package repository

import (
	"belajariah-main-service/model"
	"belajariah-main-service/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type storyRepository struct {
	db *sqlx.DB
}

type StoryRepository interface {
	GetAllStory(skip, take int, filter string) ([]model.Story, error)
	GetAllStoryCount(filter string) (int, error)
}

func InitStoryRepository(db *sqlx.DB) StoryRepository {
	return &storyRepository{
		db,
	}
}

func (storyRepository *storyRepository) GetAllStory(skip, take int, filter string) ([]model.Story, error) {
	var storyList []model.Story
	query := fmt.Sprintf(`
	SELECT
		id,
		code,
		category_code,
		header_image,
		banner_image,
		video_code,
		title,
		content,
		source,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		deleted_by,
		deleted_date
	FROM 
		master_story
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := storyRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllStory => ", sqlError)
	} else {
		defer rows.Close()
		for rows.Next() {
			var isActive bool
			var id int
			var createdDate time.Time
			var modifiedDate, deletedDate sql.NullTime
			var code, categoryCode, title, content, createdBy string
			var source, headerImg, bannerImg, videoCode, modifiedBy, deletedBy sql.NullString

			sqlError := rows.Scan(
				&id,
				&code,
				&categoryCode,
				&headerImg,
				&bannerImg,
				&videoCode,
				&title,
				&content,
				&source,
				&isActive,
				&createdBy,
				&createdDate,
				&modifiedBy,
				&modifiedDate,
				&deletedBy,
				&deletedDate,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllStory => ", sqlError)
			} else {
				storyList = append(
					storyList,
					model.Story{
						ID:           id,
						Code:         code,
						CategoryCode: categoryCode,
						HeaderImage:  headerImg,
						BannerImage:  bannerImg,
						VideoCode:    videoCode,
						Title:        title,
						Content:      content,
						Source:       source,
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
	return storyList, sqlError
}

func (storyRepository *storyRepository) GetAllStoryCount(filter string) (int, error) {
	var count int
	query := fmt.Sprintf(`
	SELECT COUNT(*) FROM 
		master_story  
	WHERE 
		deleted_by IS NULL AND
		is_active=true
	%s
	`, filter)

	row := storyRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllStoryCount => ", sqlError)
		count = 0
	}
	return count, sqlError
}
