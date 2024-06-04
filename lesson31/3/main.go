package main

import (
	"database/sql"
	"fmt"
	"math"
	"strconv"

	_ "github.com/go-faker/faker/v3"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "employees"
	password = "root"
)

func main() {
	conn := fmt.Sprintf(`host=%s port=%d user=%s dbname=%s password=%s sslmode=disable`, host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// query := `insert into employees(firstname, surname, age, salary, experience)
	// values ($1, $2, $3, $4, $5)`
	// for i := 0; i < 1000000; i++ {
	// 	_, err := db.Exec(query, faker.FirstName(), faker.FirstName(), i/100,
	// 		i, i%10)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if i%1000 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	min := math.MaxFloat64
	max := float64(-15152151212)

	for i := 0; i < 1000; i++ {
		rows, err := db.Query(`explain (analyse)
		select * from employees
		where salary = 999999 and 
		id = '1c236221-7c8f-4e5d-b1a3-1b822de281ab'
		and firstname = 'Jedediah';`)
		if err != nil {
			panic(err)
		}
		row := ""
		for rows.Next() {
			err = rows.Scan(&row)
			if err != nil {
				panic(err)
			}
		}
		timeTook, err := strconv.ParseFloat(row[16:21], 64)
		if err != nil {
			panic(err)
		}
		if min > timeTook {
			min = timeTook
		}
		if max < timeTook {
			max = timeTook
		}
	}
	fmt.Println(min)
	fmt.Println(max)
}
