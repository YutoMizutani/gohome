package controller

import (
	"github.com/YutoMizutani/gohome/app/presenter/usecase"
	"github.com/gin-gonic/gin"
)

type WeatherController struct {
	UseCase usecase.WeatherUseCase
}

func (controller *WeatherController) Fetch(c *gin.Context) {
	weather, err := controller.UseCase.Fetch()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, weather)
}
