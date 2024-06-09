package mockdatagenerator

import (
	"encoding/json"
	"leetcode/model"
	"leetcode/storage/postgres"
	"os"
)

func ProblemsGenerator(p *postgres.Problems) error {
	f, err := os.OpenFile("/home/saidakbar/Desktop/Go11Group/Saidakbar/lesson35/MockDataGenerator/problemsJSON.json", os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	var problems []model.Problem
	err = json.NewDecoder(f).Decode(&problems)
	if err != nil {
		return err
	}

	for _, problem := range problems {
		err := p.CreateProblem(&problem)
		if err != nil {
			return err
		}
	}
	return nil
}
