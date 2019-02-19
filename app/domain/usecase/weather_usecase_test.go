package usecase_test

import (
	"log"
	"testing"

	"github.com/YutoMizutani/gohome/app/domain/entity"
	"github.com/YutoMizutani/gohome/app/domain/usecase"
)

type weatherRepositoryMock struct {
}

func (repository *weatherRepositoryMock) Fetch() (*entity.Weather, error) {
	weatherEntity := &entity.Weather{}
	weatherEntity.Timezone = "Timezone"
	weatherEntity.Summary = "Summary"
	weatherEntity.Temperature = 0.0
	weatherEntity.TemperatureMax = 0.0
	weatherEntity.TemperatureMin = 0.0
	weatherEntity.Humidity = 0.0
	return weatherEntity, nil
}

func TestWeatherUseCaseFetch(t *testing.T) {
	mock := new(weatherRepositoryMock)

	weatherUseCase := usecase.WeatherUseCase{}
	weatherUseCase.Repository = mock

	entity, err := weatherUseCase.Fetch()
	if err != nil {
		log.Fatal("Error entity returns nil")
	}

	if entity.Timezone != "Timezone" {
		log.Fatal("Error Timezone value")
	}
	if entity.Summary != "Summary" {
		log.Fatal("Error Summary value")
	}
	if entity.Temperature != 0.0 {
		log.Fatal("Error Temperature value")
	}
	if entity.TemperatureMax != 0.0 {
		log.Fatal("Error TemperatureMax value")
	}
	if entity.TemperatureMin != 0.0 {
		log.Fatal("Error TemperatureMin value")
	}
	if entity.Humidity != 0.0 {
		log.Fatal("Error Humidity value")
	}
}
