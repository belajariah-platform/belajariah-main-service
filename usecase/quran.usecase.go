package usecase

import (
	"belajariah-main-service/model"
	"belajariah-main-service/repository"
	"belajariah-main-service/shape"
	"belajariah-main-service/utils"
)

type quranUsecase struct {
	quranRepository repository.QuranRepository
}

type QuranUsecase interface {
	GetAllQuran(query model.Query) ([]shape.Quran, int, error)
	GetAllAyatQuran(query model.Query) ([]shape.Quran, int, error)
	GetAllQuranView(query model.Query) ([]shape.Quran, int, error)
}

func InitQuranUsecase(quranRepository repository.QuranRepository) QuranUsecase {
	return &quranUsecase{
		quranRepository,
	}
}

func (quranUsecase *quranUsecase) GetAllQuran(query model.Query) ([]shape.Quran, int, error) {
	var filterQuery string
	var qurans []model.Quran
	var quranResult []shape.Quran

	filterQuery = utils.GetFilterHandler(query.Filters)

	qurans, err := quranUsecase.quranRepository.GetAllQuran(query.Skip, query.Take, filterQuery)
	count, errCount := quranUsecase.quranRepository.GetAllQuranCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range qurans {
			quranResult = append(quranResult, shape.Quran{
				ID:              value.ID,
				Code:            value.Code,
				Surat_Name:      value.SuratName,
				Surat_Text:      value.SuratText,
				Surat_Translate: value.SuratTranslate,
				Count_Ayat:      value.CountAyat,
				Is_Active:       value.IsActive,
				Created_By:      value.CreatedBy,
				Created_Date:    value.CreatedDate,
				Modified_By:     value.ModifiedBy.String,
				Modified_Date:   value.ModifiedDate.Time,
				Deleted_By:      value.DeletedBy.String,
				Deleted_Date:    value.DeletedDate.Time,
			})
		}
	}
	quranEmpty := make([]shape.Quran, 0)
	if len(quranResult) == 0 {
		return quranEmpty, count, err
	}
	return quranResult, count, err
}

func (quranUsecase *quranUsecase) GetAllAyatQuran(query model.Query) ([]shape.Quran, int, error) {
	var filterQuery string
	var qurans []model.Quran
	var quranResult []shape.Quran

	filterQuery = utils.GetFilterHandler(query.Filters)

	qurans, err := quranUsecase.quranRepository.GetAllAyatQuran(query.Skip, query.Take, filterQuery)
	count, errCount := quranUsecase.quranRepository.GetAllAyatQuranCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range qurans {
			quranResult = append(quranResult, shape.Quran{
				ID:             value.ID,
				Code:           value.Code,
				Surat_Code:     value.SuratCode,
				Ayat_Number:    value.AyatNumber,
				Ayat_Text:      value.AyatText,
				Ayat_Translate: value.AyatTranslate,
				Juz_Number:     value.JuzNumber,
				Page_Number:    value.PageNumber,
				Is_Active:      value.IsActive,
				Created_By:     value.CreatedBy,
				Created_Date:   value.CreatedDate,
				Modified_By:    value.ModifiedBy.String,
				Modified_Date:  value.ModifiedDate.Time,
				Deleted_By:     value.DeletedBy.String,
				Deleted_Date:   value.DeletedDate.Time,
			})
		}
	}
	quranEmpty := make([]shape.Quran, 0)
	if len(quranResult) == 0 {
		return quranEmpty, count, err
	}
	return quranResult, count, err
}

func (quranUsecase *quranUsecase) GetAllQuranView(query model.Query) ([]shape.Quran, int, error) {
	var filterQuery string
	var qurans []model.Quran
	var quranResult []shape.Quran

	filterQuery = utils.GetFilterHandler(query.Filters)

	qurans, err := quranUsecase.quranRepository.GetAllQuranView(query.Skip, query.Take, filterQuery)
	count, errCount := quranUsecase.quranRepository.GetAllQuranViewCount(filterQuery)

	if err == nil && errCount == nil {
		for _, value := range qurans {
			quranResult = append(quranResult, shape.Quran{
				ID:              value.ID,
				Code:            value.Code,
				Surat_Name:      value.SuratName,
				Surat_Text:      value.SuratText,
				Surat_Translate: value.SuratTranslate,
				Count_Ayat:      value.CountAyat,
				Ayat_Number:     value.AyatNumber,
				Ayat_Text:       value.AyatText,
				Ayat_Translate:  value.AyatTranslate,
				Juz_Number:      value.JuzNumber,
				Page_Number:     value.PageNumber,
				Is_Active:       value.IsActive,
				Created_By:      value.CreatedBy,
				Created_Date:    value.CreatedDate,
				Modified_By:     value.ModifiedBy.String,
				Modified_Date:   value.ModifiedDate.Time,
				Deleted_By:      value.DeletedBy.String,
				Deleted_Date:    value.DeletedDate.Time,
			})
		}
	}
	quranEmpty := make([]shape.Quran, 0)
	if len(quranResult) == 0 {
		return quranEmpty, count, err
	}
	return quranResult, count, err
}
