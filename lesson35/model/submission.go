package model

import "time"

type Submission struct {
	Id               string
	UserId           string
	ProblemId        string
	Language         string
	SubmissionStatus string
	Code             string
	SubmissionDate   time.Time
	Time
}

type FilterSubmission struct {
	UserId           *string
	ProblemId        *string
	Language         *string
	SubmissionStatus *string
	Code             *string
}
