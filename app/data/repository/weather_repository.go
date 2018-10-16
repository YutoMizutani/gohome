package repository

import (
	"os"

	"github.com/YutoMizutani/gohome/app/data/datastore"
	"github.com/YutoMizutani/gohome/app/data/entity"
)

type WeatherRepository struct {
	DataStore datastore.DarkSkyDataStore
}

func (repository *WeatherRepository) Fetch() (weatherEntity *entity.WeatherEntity, err error) {
	latitude := os.Getenv("WEATHER_LATITUDE")
	longitude := os.Getenv("WEATHER_LONGITUDE")
	f, err := repository.DataStore.Fetch(latitude, longitude)
	if err != nil {
		return nil, err
	}

	weatherEntity.Timezone = f.Timezone
	weatherEntity.Summary = f.Currently.Summary
	length := len(f.Daily.Data)
	summaryAll := make([]string, length)
	for i, e := range f.Daily.Data {
		summaryAll[i] = e.Summary
	}
	weatherEntity.SummaryAll = summaryAll
	weatherEntity.Temperature = f.Currently.Temperature
	weatherEntity.TemperatureMax = f.Currently.TemperatureMax
	weatherEntity.TemperatureMin = f.Currently.TemperatureMin
	weatherEntity.Humidity = f.Currently.Humidity
	return weatherEntity, nil
}
