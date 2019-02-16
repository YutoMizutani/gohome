package usecase

import (
	"github.com/YutoMizutani/gohome/app/domain/model"
)

type MenuUseCase struct {
}

func (usecase *MenuUseCase) Get() (*model.MenuModels, error) {
	models := model.MenuModels{}

	// Life
	lifeModel := model.MenuModel{
		Title:            "Lives",
		Description:      "Daily lives contents",
		NumberOfContents: 1,
		ImageURL:         "",
	}
	models = append(models, lifeModel)

	return &models, nil
}
