package repository

import (
	"errors"

	"github.com/YutoMizutani/gohome/app/application/util"
	"github.com/YutoMizutani/gohome/app/data/datastore"
	"github.com/YutoMizutani/gohome/app/data/entity"
)

type WeatherRepository struct {
	DataStore datastore.DarkSkyDataStore
}

func (repository *WeatherRepository) Fetch() (*entity.WeatherEntity, error) {
	apiKey := util.Getenv("DARK_SKY_API_KEY")
	latitude := util.Getenv("WEATHER_LATITUDE")
	longitude := util.Getenv("WEATHER_LONGITUDE")
	empty := "nil"
	if apiKey == empty || latitude == empty || longitude == empty {
		err := errors.New("Error env required values are empty")
		return nil, err
	}

	f, err := repository.DataStore.Fetch(apiKey, latitude, longitude)
	if err != nil {
		return nil, err
	}

	weatherEntity := &entity.WeatherEntity{}
	weatherEntity.Timezone = f.Timezone
	weatherEntity.Summary = f.Currently.Summary
	weatherEntity.Temperature = f.Currently.Temperature
	weatherEntity.TemperatureMax = f.Currently.TemperatureMax
	weatherEntity.TemperatureMin = f.Currently.TemperatureMin
	weatherEntity.Humidity = f.Currently.Humidity
	return weatherEntity, nil
}
