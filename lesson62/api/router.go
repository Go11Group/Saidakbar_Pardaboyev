package api

import (
	"lesson62/api/handler"
	"lesson62/models"

	"github.com/gin-gonic/gin"
)

func NewRouter(sysconf *models.SystemConfig) *gin.Engine {
	r := gin.Default()

	h := handler.NewHandler(sysconf)

	restaurant := r.Group("/restaurant")
	restaurant.POST("/create", h.CreateRestaurant)
	restaurant.GET("/getall", h.GetAllRestaurants)
	restaurant.GET("/:id", h.GetById)
	restaurant.PUT("/update", h.UpdateRestaurant)
	restaurant.DELETE("/:id", h.DeleteRestaurant)

	// meal := r.Group("/meal")

	return r
}
