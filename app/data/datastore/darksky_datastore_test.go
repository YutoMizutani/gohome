package datastore_test

import (
	"log"
	"testing"

	"github.com/YutoMizutani/gohome/app/application/util"
	"github.com/YutoMizutani/gohome/app/data/datastore"

	"github.com/joho/godotenv"
)

func TestFetch(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping e2e tests in short mode.")
	}

	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := util.Getenv("DARK_SKY_API_KEY")
	latitude := util.Getenv("WEATHER_LATITUDE")
	longitude := util.Getenv("WEATHER_LONGITUDE")
	empty := "nil"
	if apiKey == empty || latitude == empty || longitude == empty {
		log.Fatal("Error not defined required env keys")
	}

	darkSkyDataStore := new(datastore.DarkSkyDataStore)
	_, err = darkSkyDataStore.Fetch(apiKey, latitude, longitude)
	if err != nil {
		log.Fatal(err)
	}
}
