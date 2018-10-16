package translator

import (
	"github.com/YutoMizutani/gohome/app/data/entity"
	"github.com/YutoMizutani/gohome/app/domain/model"
)

type WeatherTranslator struct {
}

func (translator *WeatherTranslator) Translate(weatherEntity *entity.WeatherEntity) (weatherModel *model.WeatherModel) {
	weatherModel.Timezone = weatherEntity.Timezone
	weatherModel.Summary = weatherEntity.Summary
	weatherModel.Temperature = weatherEntity.Temperature
	weatherModel.TemperatureMax = weatherEntity.TemperatureMax
	weatherModel.TemperatureMin = weatherEntity.TemperatureMin
	weatherModel.Humidity = weatherEntity.Humidity
	return
}
