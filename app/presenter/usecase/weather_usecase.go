package usecase

import "github.com/YutoMizutani/gohome/app/domain/model"

type WeatherUseCase interface {
	Fetch() (*model.WeatherModel, error)
}
