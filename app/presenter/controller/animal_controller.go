package controller

import (
	"github.com/YutoMizutani/gohome/app/presenter/usecase"
	"github.com/gin-gonic/gin"
)

type AnimalController struct {
	UseCase usecase.AnimalUseCase
}

func (controller *AnimalController) Fetch(c *gin.Context) {
	animals, err := controller.UseCase.Fetch()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, animals)
}
