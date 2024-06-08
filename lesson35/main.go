package main

import (
	"fmt"
	mockdatagenerator "leetcode/MockDataGenerator"
	"leetcode/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	mockdatagenerator.Mockdatagenerator()
	users := postgres.NewUsersRepo(db)
	topics := postgres.NewTopicsRepo(db)
	problems := postgres.NewProblemsRepo(db)
	problemsTopics := postgres.NewProblemsTopicsRepo(db)
	languages := postgres.NewLanguagesRepo(db)
	submissions := postgres.NewSubmissionsRepo(db)

	fmt.Println("users: ", users)
	fmt.Println("topics: ", topics)
	fmt.Println("problems: ", problems)
	fmt.Println("problemsTopics: ", problemsTopics)
	fmt.Println("languages: ", languages)
	fmt.Println("submissions: ", submissions)
}
