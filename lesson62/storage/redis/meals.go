package redis

import (
	"github.com/redis/go-redis/v9"
)

type MealRepo struct {
	DB *redis.Client
}

func NewMealClient(redisClient *redis.Client) *MealRepo {
	return &MealRepo{
		DB: redisClient,
	}
}

// // Create
// func (m *MealRepo) () () {}

// // Get all 
// func (m *MealRepo) () () {}

// // Get by Id
// func (m *MealRepo) () () {}

// // Update 
// func (m *MealRepo) () () {}

// // Delete 
// func (m *MealRepo) () () {}