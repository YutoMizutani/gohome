package repository

import (
	"os"

	"github.com/YutoMizutani/gohome/app/data/datastore"
)

type WeatherRepository struct {
	DataStore datastore.DarkSkyDataStore
}

func (repository *WeatherRepository) Fetch() (err error) {
	latitude := os.Getenv("WEATHER_LATITUDE")
	longitude := os.Getenv("WEATHER_LONGITUDE")
	_, err := repository.DataStore.Fetch(latitude, longitude)
	return err
}
