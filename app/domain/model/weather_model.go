package model

type WeatherModels []WeatherModel

type WeatherModel struct {
	Timezone       string
	Summary        string
	Temperature    float64
	TemperatureMax float64
	TemperatureMin float64
	Humidity       float64
}
