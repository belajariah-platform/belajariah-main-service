package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
)

type storyUsecase struct {
	storyRepository repository.StoryRepository
}

type StoryUsecase interface {
	GetAllStory(query model.Query) ([]shape.Story, int, error)
}

func InitStoryUsecase(storyRepository repository.StoryRepository) StoryUsecase {
	return &storyUsecase{
		storyRepository,
	}
}

func (storyUsecase *storyUsecase) GetAllStory(query model.Query) ([]shape.Story, int, error) {
	var filterQuery string
	var stories []model.Story
	var storyResult []shape.Story

	filterQuery = utils.GetFilterHandler(query.Filters)

	stories, err := storyUsecase.storyRepository.GetAllStory(query.Skip, query.Take, filterQuery)
	count, errCount := storyUsecase.storyRepository.GetAllStoryCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range stories {
			storyResult = append(storyResult, shape.Story{
				ID:                  value.ID,
				Code:                value.Code,
				Story_Category_Code: value.StoryCategoryCode,
				Image_Header_Story:  value.ImageHeaderStory.String,
				Image_Banner_Story:  value.ImageBannerStory.String,
				Video_Story:         value.VideoStory.String,
				Title:               value.Title,
				Content:             value.Content,
				Source:              value.Source.String,
				Is_Active:           value.IsActive,
				Created_By:          value.CreatedBy,
				Created_Date:        value.CreatedDate,
				Modified_By:         value.ModifiedBy.String,
				Modified_Date:       value.ModifiedDate.Time,
				Is_Deleted:          value.IsDeleted,
			})
		}
	}
	ratingEmpty := make([]shape.Story, 0)
	if len(storyResult) == 0 {
		return ratingEmpty, count, err
	}
	return storyResult, count, err
}
