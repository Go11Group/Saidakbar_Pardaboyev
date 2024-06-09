package mockdatagenerator

import (
	"leetcode/model"
	"leetcode/storage/postgres"
)

func LanguagesGenerator(languages *postgres.Languages) error {
	newLanguages := []string{"C", "C++", "Java", "Python", "Python 3", "C#",
		"JavaScript (Node.js)", "Ruby", "Swift", "Go", "Scala", "Kotlin",
		"Rust", "TypeScript"}

	for _, language := range newLanguages {
		newLanguage := model.Language{
			Name: language,
		}
		err := languages.CreateLanguage(&newLanguage)
		if err != nil {
			return err
		}
	}
	return nil
}
