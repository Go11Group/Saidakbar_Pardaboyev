package handler

import (
	"lesson62/models"
	rd "lesson62/storage/redis"

	"go.uber.org/zap"
)

type Handler struct {
	Restaurant *rd.RestaurantRepo
	Meal       *rd.MealRepo
	Logger     *zap.Logger
}

func NewHandler(sysconf *models.SystemConfig) *Handler {
	meal := rd.NewMealClient(sysconf.DB)
	restaurant := rd.NewRestaurantRepo(sysconf.DB)
	return &Handler{
		Restaurant: restaurant,
		Meal:       meal,
		Logger:     sysconf.Logger,
	}
}
