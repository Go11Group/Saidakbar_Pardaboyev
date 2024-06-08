package model

type User struct {
	Id       string
	Fullname string
	Username string
	Bio      string
	Time
}

type FilterUser struct {
	Fullname *string
	Username *string
}
