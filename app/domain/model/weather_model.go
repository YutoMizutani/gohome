package model

type WeatherModel struct {
	Timezone       string  `json:"timezone"`
	Summary        string  `json:"summary"`
	Temperature    float64 `json:"temperature"`
	TemperatureMax float64 `json:"temperature_max"`
	TemperatureMin float64 `json:"temperature_min"`
	Humidity       float64 `json:"humidity"`
}
