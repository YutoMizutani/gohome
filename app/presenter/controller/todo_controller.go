package controller

import (
	"github.com/YutoMizutani/gohome/app/domain/entity"
	"github.com/YutoMizutani/gohome/app/presenter/usecase"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	UseCase usecase.TodoUseCase
}

func (controller *TodoController) Get(c *gin.Context) {
	entities, err := controller.UseCase.Get()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, entities)
}

func (controller *TodoController) Create(c *gin.Context) {
	var todo entity.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(500, err)
		return
	}
	err = controller.UseCase.Add(&todo)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, todo)
}
