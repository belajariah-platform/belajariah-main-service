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
		story_category_code,
		image_banner_story,
		image_header_story,
		video_story,
		title,
		content,
		source,
		is_active,
		created_by,
		created_date,
		modified_by,
		modified_date,
		is_deleted
	FROM 
		master.master_story
	WHERE 
		is_deleted=false AND
		is_active=true
	%s
	OFFSET %d
	LIMIT %d
	`, filter, skip, take)

	rows, sqlError := storyRepository.db.Query(query)

	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllStory => ", sqlError.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var createdDate time.Time
			var isActive, isDeleted bool
			var modifiedDate sql.NullTime
			var code, categoryCode, title, content, createdBy string
			var source, headerImg, bannerImg, videoCode, modifiedBy sql.NullString

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
				&isDeleted,
			)

			if sqlError != nil {
				utils.PushLogf("SQL error on GetAllStory => ", sqlError.Error())
			} else {
				storyList = append(
					storyList,
					model.Story{
						ID:                id,
						Code:              code,
						StoryCategoryCode: categoryCode,
						ImageHeaderStory:  headerImg,
						ImageBannerStory:  bannerImg,
						VideoStory:        videoCode,
						Title:             title,
						Content:           content,
						Source:            source,
						IsActive:          isActive,
						CreatedBy:         createdBy,
						CreatedDate:       createdDate,
						ModifiedBy:        modifiedBy,
						ModifiedDate:      modifiedDate,
						IsDeleted:         isDeleted,
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
		master.master_story  
	WHERE 
		is_deleted=false AND
		is_active=true
	%s
	`, filter)

	row := storyRepository.db.QueryRow(query)
	sqlError := row.Scan(&count)
	if sqlError != nil {
		utils.PushLogf("SQL error on GetAllStoryCount => ", sqlError.Error())
		count = 0
	}
	return count, sqlError
}
