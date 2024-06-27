package postgres

import (
	"context"
	"database/sql"
)

type WeatherRepo struct {
	DB *sql.DB
}

func NewWeatherRepo(db *sql.DB) *WeatherRepo {
	return &WeatherRepo{DB: db}
}
