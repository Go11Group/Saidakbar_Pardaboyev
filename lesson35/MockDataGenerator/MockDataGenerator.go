package mockdatagenerator

import (
	"fmt"
	"leetcode/model"
	"leetcode/storage/postgres"
)

type Repos struct {
	Users          *postgres.Users
	Topics         *postgres.Topics
	Problems       *postgres.Problems
	ProblemsTopics *postgres.ProblemsTopics
	Languages      *postgres.Languages
	Submissions    *postgres.Submissions
}

func (r *Repos) Mockdatagenerator() error {

	// err := UsersGenerator(r.Users)
	// if err != nil {
	// 	return err
	// }

	// err2 := TopicsGenerator(r.Topics)
	// if err2 != nil {
	// 	return err2
	// }

	// err3 := ProblemsGenerator(r.Problems)
	// if err3 != nil {
	// 	return err3
	// }

	// // Add Topics for mock problems
	// topics, err := r.Topics.GetTopics(&model.FilterTopic{})
	// if err != nil {
	// 	return err
	// }
	// problems, err := r.Problems.GetProblems(&model.FilterProduct{})
	// if err != nil {
	// 	return err
	// }
	// err4 := ProblemsTopicsGenerator(r.ProblemsTopics, problems, topics)
	// if err4 != nil {
	// 	return err4
	// }

	// err5 := LanguagesGenerator(r.Languages)
	// if err5 != nil {
	// 	return err5
	// }

	// Adding mock submissions
	users, err := r.Users.GetUsers(&model.FilterUser{})
	if err != nil {
		return err
	}

	problems, err := r.Problems.GetProblems(&model.FilterProduct{})
	if err != nil {
		return err
	}
	languages, err := r.Languages.GetLanguages(&model.FilterLanguage{})
	if err != nil {
		return err
	}

	err6 := SubmissionGenerator(r.Submissions, users, problems, languages)
	if err6 != nil {
		return err6
	}

	fmt.Println("Mock data muvaffaqiyatli qo'shildi")
	return nil
}
