package datastore

import (
	forecast "github.com/mlbright/forecast/v2"
)

type DarkSkyDataStore struct {
}

func (dataStore *DarkSkyDataStore) Fetch(apiKey string, latitude string, longitude string) (res *forecast.Forecast, err error) {
	res, err = forecast.Get(apiKey, latitude, longitude, "now", forecast.SI, "ja")
	return
}
