package mockdatagenerator

import (
	"leetcode/model"
	"leetcode/storage/postgres"
	"math/rand"
)

func ProblemsTopicsGenerator(problemsTopics *postgres.ProblemsTopics,
	problems []model.Problem, topics []model.Topic) error {

	for _, problem := range problems {
		problemId := problem.Id
		for i := 0; i < rand.Intn(5); i++ {
			numRundom := rand.Intn(len(topics))
			topicId := topics[numRundom].Id
			problemTopic := model.ProblemTopic{
				ProblemId: problemId,
				TopicId:   topicId,
			}
			err := problemsTopics.CreateProblemTopic(&problemTopic)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
