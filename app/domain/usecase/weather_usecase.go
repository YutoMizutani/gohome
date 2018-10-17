package usecase

import (
	"github.com/YutoMizutani/gohome/app/domain/model"
	"github.com/YutoMizutani/gohome/app/domain/repository"
	"github.com/YutoMizutani/gohome/app/domain/translator"
)

type WeatherUseCase struct {
	Repository repository.WeatherRepository
	Translator translator.WeatherTranslator
}

func (usecase *WeatherUseCase) Fetch() (*model.WeatherModel, error) {
	entity, err := usecase.Repository.Fetch()
	if err != nil {
		return nil, err
	}

	model := usecase.Translator.Translate(entity)
	return model, nil
}
