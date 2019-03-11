package usecase

import (
	"github.com/YutoMizutani/gohome/app/domain/entity"
	"github.com/YutoMizutani/gohome/app/domain/repository"
)

type WeatherUseCase struct {
	Repository repository.WeatherRepository
}

func (usecase *WeatherUseCase) Fetch() (*entity.Weather, error) {
	entity, err := usecase.Repository.Fetch()
	if err != nil {
		return nil, err
	}

	return entity, nil
}
