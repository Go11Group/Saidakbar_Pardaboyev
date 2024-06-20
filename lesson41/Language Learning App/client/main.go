package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	requestForUsers()
}

func requestForUsers() {
	// // create user
	// client := &http.Client{}
	// newUser := `{
	// 	"name": "Iskandar",
	// 	"email": "sdfghj@gmail.com",
	// 	"birthday": "2005-06-11T15:04:05Z",
	// 	"password": "12345678"
	// }`

	// req, err := http.NewRequest("POST", "http://localhost:8080/exam/users/create", bytes.NewBuffer([]byte(newUser)))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// req.Header.Set("Content-Type", "application/json")

	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(body))

	// // get Users by filter
	// client := &http.Client{}
	// getReq, err := http.NewRequest("GET", "http://localhost:8080/exam/users/getall", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// q := getReq.URL.Query()
	// q.Add("limit", "1")
	// getReq.URL.RawQuery = q.Encode()

	// resp, err := client.Do(getReq)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()

	// users, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(users))

	// // get by id
	// client := &http.Client{}

	// url := "http://localhost:8080/exam/users/3bda9776-bad4-4f6c-b295-eb95b266c9e9"
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// user, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(user))

	// // update user
	// client := &http.Client{}

	// userInfoToUpdate := `{
	// 	"name": "Iskandar",
	// 	"email": "sdfghj@gmail.com",
	// 	"birthday": "2005-06-11T15:04:05Z",
	// 	"password": "87654321"
	// }`

	// url := "http://localhost:8080/exam/users/update/0177b2d8-c995-494d-a504-a0f5606c4c3c"
	// res, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(userInfoToUpdate)))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// resp, err := client.Do(res)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(body))

	// delete from users
	client := &http.Client{}

	url := "http://localhost:8080/exam/users/delete/0177b2d8-c995-494d-a504-a0f5606c4c3c"

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal()
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(string(body))
}
