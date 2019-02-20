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

	router.GET("/todo", func(c *gin.Context) { todoController.Get(c) })
	router.POST("/todo/create", func(c *gin.Context) { todoController.Create(c) })

	Router = router
}
