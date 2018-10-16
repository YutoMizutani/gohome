package repository

import (
	"os"

	"github.com/YutoMizutani/gohome/app/data/datastore"
)

type WeatherRepository struct {
	DataStore datastore.DarkSkyDataStore
}

func (repository *WeatherRepository) Fetch() {
	latitude := os.Getenv("WEATHER_LATITUDE")
	longitude := os.Getenv("WEATHER_LONGITUDE")
	repository.DataStore.Fetch(latitude, longitude)
	return
}
