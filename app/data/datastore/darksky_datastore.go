package datastore

import (
	"time"

	forecast "github.com/mlbright/forecast/v2"
	cache "github.com/patrickmn/go-cache"
)

type DarkSkyDataStore struct {
	cache *cache.Cache
}

func NewDarkSky() *DarkSkyDataStore {
	dataStore := &DarkSkyDataStore{}
	dataStore.cache = cache.New(10*time.Minute, 20*time.Minute)
	return dataStore
}

func (dataStore *DarkSkyDataStore) Fetch(
	apiKey string,
	latitude string,
	longitude string) (res *forecast.Forecast, err error) {

	cacheKey := "DarkSkyDataStoreResponse"

	if data, found := dataStore.cache.Get(cacheKey); found {
		res = data.(*forecast.Forecast)
		return res, err
	}

	res, err = forecast.Get(apiKey, latitude, longitude, "now", forecast.SI, forecast.English)
	if res != nil && err == nil {
		dataStore.cache.Set(cacheKey, res, cache.DefaultExpiration)
	}

	return
}
