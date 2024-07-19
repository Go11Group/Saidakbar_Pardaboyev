package systemconfig

import (
	"lesson62/models"
	"lesson62/pkg/logger"
	"lesson62/storage/redis"
)

func NewSystemConfig() (*models.SystemConfig, error) {
	logger, err := logger.New("debug", "app.log")
	if err != nil {
		return nil, err
	}

	db, err := redis.ConnectDB()
	if err != nil {
		return nil, err
	}

	return &models.SystemConfig{
		Logger: logger,
		DB:     db,
	}, nil
}
