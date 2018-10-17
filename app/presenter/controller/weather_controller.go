package controller

import "github.com/YutoMizutani/gohome/app/presenter/usecase"

type WeatherController struct {
	UseCase usecase.WeatherUseCase
}

func (controller *WeatherController) Fetch(c Context) {
	weather, err := controller.UseCase.Fetch()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, weather)
}
