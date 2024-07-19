package handler

import (
	"encoding/json"
	"fmt"
	"lesson62/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (h *Handler) CreateRestaurant(ctx *gin.Context) {
	newres := models.RestaurantCreate{}
// "a1ebb162-92cb-4f72-b9b2-141f26031575"
	err := json.NewDecoder(ctx.Request.Body).Decode(&newres)
	if err != nil {
		h.Logger.Error("Error with decoding url body", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with decoding url body: %s", err))
		return
	}

	reswithFullInfo := models.RestaurantGet{
		Id:          uuid.NewString(),
		Title:       newres.Title,
		Description: newres.Description,
		Address:     newres.Address,
		ProneNumber: newres.ProneNumber,
		Time: models.Time{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	req, err := json.Marshal(reswithFullInfo)
	if err != nil {
		h.Logger.Error("Error with marshel restaulrant info", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with marshel restaulrant info: %s", err))
		return
	}

	err = h.Restaurant.Create(ctx, reswithFullInfo.Id, string(req))
	if err != nil {
		h.Logger.Error("Error with creating restaurant", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with creating restaurant: %s", err))
		return
	}

	ctx.JSON(http.StatusOK, reswithFullInfo)
}
