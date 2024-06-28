package handler

import (
	pb "api_gateway/genproto/weather"
	"api_gateway/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) GetCurrentWeather(ctx *gin.Context) {
	city := ctx.Query("city")
	if city == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: not found city in URL query",
		})
		log.Printf("Error: not found city in URL query")
		return
	}
	ctx2, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	resp, err := h.Weather.GetCurrentWeather(ctx2, &pb.RequestCurrentWeather{City: city})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with getting data from db: %s", err),
		})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}

func (h *Handler) GetWeatherForecast(ctx *gin.Context) {
	city := ctx.Query("city")
	days := ctx.Query("days")
	if city == "" || days == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: not found city or days in URL query",
		})
		log.Printf("Error: not found city or days in URL query")
		return
	}
	daysInt, err := strconv.Atoi(days)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: days type not int",
		})
		log.Printf("Error: days type not int")
		return
	}
	ctx2, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	resp, err := h.Weather.GetWeatherForecast(ctx2, &pb.RequestWeatherForecast{City: city, Days: int32(daysInt)})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error with date of weather report",
		})
		log.Printf("Error with date of weather report")
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}

func (h *Handler) ReportWeatherCondition(ctx *gin.Context) {

	newW := models.Weather{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&newW)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error with decoding body of API",
		})
		log.Printf("Error with decoding body of API")
		return
	}

	date, err := time.Parse("2006-01-02 15-04-05", newW.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error with date of weather report",
		})
		log.Printf("Error with date of weather report")
		return
	}

	newWeather := pb.Weather{
		City:        newW.City,
		Date:        timestamppb.New(date),
		Temperature: float32(newW.Temperature),
		Humidity:    int32(newW.Humidity),
		Wind:        float32(newW.Wind),
	}

	ctx2, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	resp, err := h.Weather.ReportWeatherCondition(ctx2, &newWeather)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusAccepted, resp.Success)
}
