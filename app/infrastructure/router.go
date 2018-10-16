package infrastructure

import (
	"github.com/YutoMizutani/gohome/app/application/builder"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	animalController := new(builder.AnimalBuilder).Build()
	weatherController := new(builder.WeatherBuilder).Build()

	router.GET("/animal", func(c *gin.Context) { animalController.Fetch(c) })
	router.GET("/forecast", func(c *gin.Context) { weatherController.Fetch(c) })

	Router = router
}
