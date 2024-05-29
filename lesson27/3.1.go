package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "users"
	password = "root"
)

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("----------------------Printing with .QueryRow------------------")
	query := fmt.Sprintf("select u.full_name, a.name from users as u left join address as a on u.id = a.user_id offset $1")
	for i := 0; true; i++ {
		fullname := ""
		address := ""
		err := db.QueryRow(query, i).
			Scan(&fullname, &address)
		if err != nil {
			break
		}
		fmt.Println("Full name:", fullname, "\nAddress:", address, "\n")
	}

	fmt.Println("----------------------Printing with .Query------------------")
	query2 := fmt.Sprintf("select u.full_name, a.name from users as u left join address as a on u.id = a.user_id")
	rows, err := db.Query(query2)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		fullname := ""
		address := ""
		err := rows.Scan(&fullname, &address)
		if err != nil {
			panic(err)
		}
		fmt.Println("Full name:", fullname, "\nAddress:", address, "\n")
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	// insertQuery := "insert into users(id, full_name, age) values($1, $2, $3)"
	// _, err = db.Exec(insertQuery, uuid.NewString(), "Saidakbar Pardaboyev", 19)
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = db.Exec(insertQuery, uuid.NewString(), "Muhammadjon Ko'palov", 19)
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = db.Exec(insertQuery, uuid.NewString(), "Iskandar Iskandarov", 22)
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = db.Exec(insertQuery, uuid.NewString(), "Isroil Isroilov", 12)
	// if err != nil {
	// 	panic(err)
	// }

	// query := "insert into address(id, user_id, name) values($1, $2, $3)"
	// _, err = db.Exec(query, uuid.NewString(), "c46b1764-096d-4db4-a5ca-b166499b04d6", "Yunusobod 9 kvartil")
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = db.Exec(query, uuid.NewString(), "146f5adf-1150-47a7-960e-3817d6792af3", "Sergeli 2 kvartil")
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = db.Exec(query, uuid.NewString(), "b75ffa78-54fc-4849-a713-705b1074d448", "Chilonzor 5 kvartil")
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = db.Exec(query, uuid.NewString(), "3aa2555d-b278-419e-8ae9-f1a4e678b7dd", "Yunusobod 7 kvartil")
	// if err != nil {
	// 	panic(err)
	// }
}
