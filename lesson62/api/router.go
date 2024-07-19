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
	
	// meal := r.Group("/meal")

	return r
}
