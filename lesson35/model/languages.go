package model

type Language struct {
	Id   string
	Name string
	Time
}

type FilterLanguage struct {
	Name *string
}
