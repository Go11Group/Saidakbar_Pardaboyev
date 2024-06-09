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

	users := postgres.NewUsersRepo(db)
	topics := postgres.NewTopicsRepo(db)
	problems := postgres.NewProblemsRepo(db)
	problemsTopics := postgres.NewProblemsTopicsRepo(db)
	languages := postgres.NewLanguagesRepo(db)
	submissions := postgres.NewSubmissionsRepo(db)

	repos := mockdatagenerator.Repos{
		Users:          users,
		Topics:         topics,
		Problems:       problems,
		ProblemsTopics: problemsTopics,
		Languages:      languages,
		Submissions:    submissions,
	}

	err = repos.Mockdatagenerator()
	if err != nil {
		panic(err)
	}

	// pro, err := problems.GetProblems(&model.FilterProduct{})
	// if err != nil {
	// 	panic(err)
	// }
	// for _, p := range pro {
	// 	fmt.Println(p.Problem_num, p.Title, p.Status, p.Description, string(p.Examples), string(p.Constraints))
	// 	fmt.Println("------------------------------------")
	// }

	fmt.Println("users: ", users)
	fmt.Println("topics: ", topics)
	fmt.Println("problems: ", problems)
	fmt.Println("problemsTopics: ", problemsTopics)
	fmt.Println("languages: ", languages)
	fmt.Println("submissions: ", submissions)
}
