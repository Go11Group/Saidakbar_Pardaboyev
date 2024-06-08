package model

type Problem struct {
	Id          string
	Problem_num int
	Title       string
	Status      string
	Description string
	Examples    []string
	Constraints []string
	Time
}

type FilterProduct struct {
	Problem_num *int
	Title       *string
	Status      *string
}
