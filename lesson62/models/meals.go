package models

type MealCreate struct {
	Title        string
	Description  string
	Price        float32
	RestaurantId string
}

type MealGet struct {
	Id           string
	Title        string
	Description  string
	Price        float32
	RestaurantId string
	Time
}

type MealUpdate struct {
	Id           string
	Title        string
	Description  string
	Price        float32
	RestaurantId string
}
