package model

type WeatherModel struct {
	Timezone       string
	Summary        string
	Temperature    float64
	TemperatureMax float64
	TemperatureMin float64
	Humidity       float64
}
