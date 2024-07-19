package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RestaurantRepo struct {
	DB *redis.Client
}

func NewRestaurantRepo(redisClient *redis.Client) *RestaurantRepo {
	return &RestaurantRepo{
		DB: redisClient,
	}
}

// Create
func (r *RestaurantRepo) Create(ctx context.Context, id, restaurant string) (error) {
	return r.DB.HSet(ctx, "restaurants", id, restaurant).Err()
}

// // Get all 
// func (r *RestaurantRepo) () () {}

// // Get by Id
// func (r *RestaurantRepo) () () {}

// // Update 
// func (r *RestaurantRepo) () () {}

// // Delete 
// func (r *RestaurantRepo) () () {}
