package postgres

import (
	"database/sql"
	"fmt"
	"leetcode/model"
	"time"
)

type ProblemsTopics struct {
	DB *sql.DB
}

func NewProblemsTopicsRepo(db *sql.DB) *ProblemsTopics {
	return &ProblemsTopics{DB: db}
}

// Create
func (p *ProblemsTopics) CreateProblemTopic(problemTopic *model.ProblemTopic) error {
	tr, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	query := `insert into problems_topics(problem_id,topic_id) 
		values($1,$2)`
	_, err = tr.Exec(query, problemTopic.ProblemId, problemTopic.TopicId)
	return err
}

// Read All
func (p *ProblemsTopics) GetProblemsTopics(filter *model.FilterProductTopic) ([]model.ProblemTopic, error) {
	problemTopics := []model.ProblemTopic{}
	argums := []interface{}{}
	query := `select * from problems_topics where deleted_at is null`
	num := 1
	if filter.ProblemId != nil {
		query += fmt.Sprintf(` and problem_id=$%d`, num)
		argums = append(argums, filter.ProblemId)
		num++
	}
	if filter.TopicId != nil {
		query += fmt.Sprintf(` and topic_id=$%d`, num)
		argums = append(argums, filter.TopicId)
	}

	rows, err := p.DB.Query(query, argums...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		problemTopic := model.ProblemTopic{}
		err := rows.Scan(&problemTopic.Id, &problemTopic.ProblemId,
			&problemTopic.TopicId, &problemTopic.Created_at,
			&problemTopic.Updated_at, &problemTopic.Deleted_at)
		if err != nil {
			return nil, err
		}
		problemTopics = append(problemTopics, problemTopic)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return problemTopics, nil
}

// Read by Id
func (p *ProblemsTopics) GetProblemTopicsById(problemTopicId string) (model.ProblemTopic, error) {
	problemTopic := model.ProblemTopic{}

	query := `select * from problems_topics where id=$1 and deleted_at is null`

	row := p.DB.QueryRow(query, problemTopicId)
	err := row.Scan(&problemTopic.Id, &problemTopic.ProblemId,
		&problemTopic.TopicId, &problemTopic.Created_at,
		&problemTopic.Updated_at, &problemTopic.Deleted_at)

	return problemTopic, err
}

// Update
func (p *ProblemsTopics) UpdateProblemTopic(problemTopic *model.ProblemTopic) error {
	tr, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update problems_topics set problem_id=$1, topic_id=$2,
		updated_at=$3 where id=$4 and deleted_at is null`
	_, err = tr.Exec(query, problemTopic.ProblemId, problemTopic.TopicId,
		time.Now(), problemTopic.Id)

	return err
}

// Delete
func (p *ProblemsTopics) DeleteProblemTopic(problemTopicId string) error {
	tr, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update problems_topics set deleted_at=$1 where id=$2 and 
	deleted_at is null`
	_, err = tr.Exec(query, time.Now(), problemTopicId)

	return err
}
