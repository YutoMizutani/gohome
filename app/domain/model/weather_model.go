package model

type WeatherModels []WeatherModel

type WeatherModel struct {
	Timezone       string
	Summary        string
	SummaryAll     []string
	Temperature    float64
	TemperatureMax float64
	TemperatureMin float64
	Humidity       float64
}
