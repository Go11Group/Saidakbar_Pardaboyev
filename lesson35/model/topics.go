package model

type Topic struct {
	Id   string
	Name string
	Time
}

type FilterTopic struct {
	Name *string
}