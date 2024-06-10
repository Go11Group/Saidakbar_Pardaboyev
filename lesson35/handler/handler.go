package handler

import (
	"leetcode/storage/postgres"
)

type Handler struct {
	Users          *postgres.Users
	Topics         *postgres.Topics
	Problems       *postgres.Problems
	ProblemsTopics *postgres.ProblemsTopics
	Languages      *postgres.Languages
	Submissions    *postgres.Submissions
}

func NewHandlerRepo(users *postgres.Users, topics *postgres.Topics,
	problems *postgres.Problems, problemsTopics *postgres.ProblemsTopics,
	languages *postgres.Languages, submissions *postgres.Submissions) *Handler {

	return &Handler{
		Users:          users,
		Topics:         topics,
		Problems:       problems,
		ProblemsTopics: problemsTopics,
		Languages:      languages,
		Submissions:    submissions,
	}
}
