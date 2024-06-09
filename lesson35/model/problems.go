package model

type Problem struct {
	Id          string
	Problem_num int
	Title       string
	Status      string
	Description string
	Examples    []uint8
	Constraints []uint8
	Time
}

type FilterProduct struct {
	Problem_num *int
	Title       *string
	Status      *string
}
