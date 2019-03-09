package usecase

import "github.com/YutoMizutani/gohome/app/domain/entity"

type WeatherUseCase interface {
	Fetch() (*entity.Weather, error)
}
