package usecase

import "github.com/YutoMizutani/gohome/app/domain/model"

type WeatherUsecase interface {
	Fetch() (*model.WeatherModel, error)
}
