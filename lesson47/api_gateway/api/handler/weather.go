package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCurrentWeather(g *gin.Context) {
	// // get current weather
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// resp, err := (*weatherClient).GetCurrentWeather(ctx, &pbWeather.RequestCurrentWeather{City: "Tashkent"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp)
}

func (h *Handler) GetWeatherForecast(g *gin.Context) {
	// // get weather forecast
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// resp, err := (*weatherClient).GetWeatherForecast(ctx, &pbWeather.RequestWeatherForecast{City: "Tashkent", Days: 7})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, val := range resp.Forecastdays {
	// 	fmt.Println(val)
	// }
}

func (h *Handler) ReportWeatherCondition(g *gin.Context) {
	// // report weather
	// newWeather := pbWeather.Weather{
	// 	City:        "Tashkent",
	// 	Date:        timestamppb.New(time.Now()),
	// 	Temperature: 42,
	// 	Humidity:    17,
	// 	Wind:        8,
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// resp, err := (*weatherClient).ReportWeatherCondition(ctx, &newWeather)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp.Success)
}
