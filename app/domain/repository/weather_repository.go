package repository

import "github.com/YutoMizutani/gohome/app/data/entity"

type WeatherRepository interface {
	Fetch() (*entity.WeatherEntity, error)
}
