package repository

import "github.com/YutoMizutani/gohome/app/domain/entity"

type WeatherRepository interface {
	Fetch() (*entity.Weather, error)
}
