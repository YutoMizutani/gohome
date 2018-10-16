package controller

import "github.com/YutoMizutani/gohome/app/presenter/usecase"

type WeatherController struct {
	Usecase usecase.WeatherUsecase
}

func (controller *WeatherController) Fetch(c Context) {
	weather, err := controller.Usecase.Fetch()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, weather)
}
