package models

type Weather struct {
	City        string `json:"city"`
	Date        string `json:"date"`
	Temperature int    `json:"temperature"`
	Humidity    int    `json:"humidity"`
	Wind        int    `json:"wind"`
}
