package service

import (
	"context"
	pb "weather_and_bus/weather/genproto/weather"
	"weather_and_bus/weather/storage/postgres"
)

type Weather struct {
	pb.UnimplementedWeatherServiceServer
	WeatherRepo *postgres.WeatherRepo
}

func NewWeather(weatherRepo *postgres.WeatherRepo) *Weather {
	return &Weather{WeatherRepo: weatherRepo}
}

func (w *Weather) GetCurrentWeather(ctx context.Context, req *pb.RequestCurrentWeather) (*pb.RequestCurrentWeather, error) {

}

func (w *Weather) GetWeatherForecast(ctx context.Context, req *pb.RequestWeatherForecast) (*pb.ResponceWeatherForecast, error) {

}

func (w *Weather) ReportWeatherCondition(ctx context.Context, req *pb.ResponceWeatherCondition) (*pb.ResponceWeatherCondition, error) {

}
