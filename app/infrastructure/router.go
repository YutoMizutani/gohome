package infrastructure

import (
	"github.com/YutoMizutani/gohome/app/application/builder"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	menuController := new(builder.MenuBuilder).Build()
	animalController := new(builder.AnimalBuilder).Build()
	weatherController := new(builder.WeatherBuilder).Build()
	todoController := new(builder.TodoBuilder).Build()

	router.GET("/menu", func(c *gin.Context) { menuController.Get(c) })
	router.GET("/animal", func(c *gin.Context) { animalController.Fetch(c) })
	router.GET("/forecast", func(c *gin.Context) { weatherController.Fetch(c) })

	// Todo list
	router.GET("/todos", func(c *gin.Context) { todoController.GetAll(c) })
	router.POST("/todos", func(c *gin.Context) { todoController.Create(c) })
	router.GET("/todos/:id", func(c *gin.Context) { todoController.Get(c) })
	router.PUT("/todos/:id", func(c *gin.Context) { todoController.Update(c) })
	router.DELETE("/todos/:id", func(c *gin.Context) { todoController.Delete(c) })
	router.PUT("/todos/:id/done", func(c *gin.Context) { todoController.UpdateDone(c, true) })
	router.DELETE("/todos/:id/done", func(c *gin.Context) { todoController.UpdateDone(c, false) })

	Router = router
}
