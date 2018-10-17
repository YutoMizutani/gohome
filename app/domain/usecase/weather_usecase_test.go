package usecase_test

import (
	"log"
	"testing"

	"github.com/YutoMizutani/gohome/app/data/entity"
	"github.com/YutoMizutani/gohome/app/domain/translator"
	"github.com/YutoMizutani/gohome/app/domain/usecase"
)

type weatherRepositoryMock struct {
}

func (repository *weatherRepositoryMock) Fetch() (weatherEntity *entity.WeatherEntity, err error) {
	weatherEntity = &entity.WeatherEntity{}
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
	weatherUseCase.Translator = translator.WeatherTranslator{}

	model, err := weatherUseCase.Fetch()
	if err != nil {
		log.Fatal("Error entity returns nil")
	}

	if model.Timezone != "Timezone" {
		log.Fatal("Error Timezone value")
	}
	if model.Summary != "Summary" {
		log.Fatal("Error Summary value")
	}
	if model.Temperature != 0.0 {
		log.Fatal("Error Temperature value")
	}
	if model.TemperatureMax != 0.0 {
		log.Fatal("Error TemperatureMax value")
	}
	if model.TemperatureMin != 0.0 {
		log.Fatal("Error TemperatureMin value")
	}
	if model.Humidity != 0.0 {
		log.Fatal("Error Humidity value")
	}
}
