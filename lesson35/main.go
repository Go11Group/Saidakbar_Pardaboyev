package main

import (
	"leetcode/handler"
	"leetcode/router"
	"leetcode/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	users := postgres.NewUsersRepo(db)
	topics := postgres.NewTopicsRepo(db)
	problems := postgres.NewProblemsRepo(db)
	problemsTopics := postgres.NewProblemsTopicsRepo(db)
	languages := postgres.NewLanguagesRepo(db)
	submissions := postgres.NewSubmissionsRepo(db)

	// // Generate mock data
	// repos := mockdatagenerator.Repos{
	// 	Users:          users,
	// 	Topics:         topics,
	// 	Problems:       problems,
	// 	ProblemsTopics: problemsTopics,
	// 	Languages:      languages,
	// 	Submissions:    submissions,
	// }
	// err = repos.Mockdatagenerator()
	// if err != nil {
	// 	panic(err)
	// }

	h := handler.NewHandlerRepo(users, topics, problems, problemsTopics,
		languages, submissions)
	server := router.CreateNewServer(h)
	server.ListenAndServe()
}
