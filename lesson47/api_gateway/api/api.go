package api

import (
	"api_gateway/api/handler"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	h := handler.NewHandler()

	weatherRouter := r.Group("/weather")
	weatherRouter.GET("/current", h.GetCurrentWeather)
	weatherRouter.GET("/forecast", h.GetWeatherForecast)
	weatherRouter.POST("/report", h.ReportWeatherCondition)

	transportRouter := r.Group("/transport")
	transportRouter.GET("/bus_schedule", h.GetBusSchedule)
	transportRouter.GET("/bus_location", h.TrackBusLocation)
	transportRouter.POST("/report", h.ReportTrafficJam)

	return r
}
