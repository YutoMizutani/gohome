package controller

import (
	"strconv"

	"github.com/YutoMizutani/gohome/app/domain/entity"
	"github.com/YutoMizutani/gohome/app/presenter/usecase"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	UseCase usecase.TodoUseCase
}

func (controller *TodoController) GetAll(c *gin.Context) {
	entities, err := controller.UseCase.GetAll()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, entities)
}

func (controller *TodoController) Get(c *gin.Context) {
	idString := c.Param("id")
	id64, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.Status(404)
		return
	}
	id := uint(id64)

	entity, err := controller.UseCase.Get(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, entity)
}

func (controller *TodoController) Create(c *gin.Context) {
	var todo *entity.Todo
	err := c.BindJSON(todo)
	if err != nil {
		c.JSON(500, err)
		return
	}
	err = controller.UseCase.Add(todo)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, todo)
}

func (controller *TodoController) Update(c *gin.Context) {
	var todo *entity.Todo
	err := c.BindJSON(todo)
	if err != nil {
		c.JSON(500, err)
		return
	}
	todo, err = controller.UseCase.Update(todo)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, todo)
}

func (controller *TodoController) UpdateDone(c *gin.Context, isDone bool) {
	idString := c.Param("id")
	id64, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.Status(404)
		return
	}
	id := uint(id64)

	todo, err := controller.UseCase.UpdateDoneState(id, isDone)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, todo)
}

func (controller *TodoController) Delete(c *gin.Context) {
	idString := c.Param("id")
	id64, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.Status(404)
		return
	}
	id := uint(id64)

	err = controller.UseCase.Delete(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.Status(200)
}
