package model

type ProblemTopic struct {
	Id        string
	ProblemId string
	TopicId   string
	Time
}

type FilterProductTopic struct {
	ProblemId *string
	TopicId   *string
}
