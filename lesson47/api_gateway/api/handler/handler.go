package handler

import (
	pbT "api_gateway/genproto/transport"
	pbW "api_gateway/genproto/weather"
	"api_gateway/pkg"
)

type Handler struct {
	Transport *pbT.TransportServerClient
	Weather   *pbW.WeatherServiceClient
}

func NewHandler() *Handler {
	transportClient := pkg.CreateTransportServerClient()
	weatherClient := pkg.CreateWeatherServerClient()

	return &Handler{
		Transport: transportClient,
		Weather:   weatherClient,
	}
}
