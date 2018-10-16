package builder

import (
	"github.com/YutoMizutani/gohome/app/data/datastore"
	"github.com/YutoMizutani/gohome/app/data/repository"
	"github.com/YutoMizutani/gohome/app/domain/translator"
	"github.com/YutoMizutani/gohome/app/domain/usecase"
	"github.com/YutoMizutani/gohome/app/presenter/controller"
)

type WeatherBuilder struct {
}

func (builder *WeatherBuilder) Build() *controller.WeatherController {
	return &controller.WeatherController{
		Usecase: &usecase.WeatherUsecase{
			Repository: &repository.WeatherRepository{
				DataStore: datastore.DarkSkyDataStore{},
			},
			Translator: translator.WeatherTranslator{},
		},
	}
}
