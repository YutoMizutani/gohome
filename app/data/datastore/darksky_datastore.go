package datastore

import (
	forecast "github.com/mlbright/forecast/v2"
)

type DarkSkyDataStore struct {
	APIKey string
}

func (dataStore *DarkSkyDataStore) Fetch(latitude string, longitude string) (res *forecast.Forecast, err error) {
	res, err = forecast.Get(dataStore.APIKey, latitude, longitude, "now", forecast.SI, "ja")
	return
}
