package postgres

import (
	"database/sql"
	"fmt"
	"leetcode/model"
	"time"
)

type Submissions struct {
	DB *sql.DB
}

func NewSubmissionsRepo(db *sql.DB) *Submissions {
	return &Submissions{DB: db}
}

// Create
func (s *Submissions) CreateSubmission(submission *model.Submission) error {
	tr, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	query := `insert into submissions(user_id,problem_id,language,
		submission_status,code) values($1,$2,$3,$4,$5)`
	_, err = tr.Exec(query, submission.UserId, submission.ProblemId,
		submission.Language, submission.SubmissionStatus,
		submission.Code)
	return err
}

// Read All
func (s *Submissions) GetSubmissions(filter *model.FilterSubmission) ([]model.Submission, error) {
	submissions := []model.Submission{}
	argums := []interface{}{}
	query := `select * from submissions where deleted_at is null`
	num := 1
	if filter.UserId != nil {
		query += fmt.Sprintf(` and user_id=$%d`, num)
		argums = append(argums, filter.UserId)
		num++
	}
	if filter.ProblemId != nil {
		query += fmt.Sprintf(` and problem_id=$%d`, num)
		argums = append(argums, filter.ProblemId)
		num++
	}
	if filter.Language != nil {
		query += fmt.Sprintf(` and language=$%d`, num)
		argums = append(argums, filter.Language)
		num++
	}
	if filter.SubmissionStatus != nil {
		query += fmt.Sprintf(` and submission_status=$%d`, num)
		argums = append(argums, filter.SubmissionStatus)
		num++
	}
	if filter.Code != nil {
		query += fmt.Sprintf(` and code=$%d`, num)
		argums = append(argums, filter.Code)
	}

	rows, err := s.DB.Query(query, argums...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		submission := model.Submission{}
		err := rows.Scan(&submission.Id, &submission.UserId,
			&submission.ProblemId, &submission.Language,
			&submission.SubmissionStatus, &submission.Code,
			&submission.SubmissionDate, &submission.Created_at,
			&submission.Updated_at, &submission.Deleted_at)
		if err != nil {
			return nil, err
		}
		submissions = append(submissions, submission)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return submissions, nil
}

// Read by Id
func (s *Submissions) GetSubmissionById(submissionId string) (model.Submission, error) {
	submission := model.Submission{}

	query := `select * from submissions where id=$1 and deleted_at is null`
	row := s.DB.QueryRow(query, submissionId)
	err := row.Scan(&submission.Id, &submission.UserId,
		&submission.ProblemId, &submission.Language,
		&submission.SubmissionStatus, &submission.Code,
		&submission.SubmissionDate, &submission.Created_at,
		&submission.Updated_at, &submission.Deleted_at)

	return submission, err
}

// Update
func (s *Submissions) UpdateSubmission(submission *model.Submission) error {
	tr, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update submissions set user_id=$1, problem_id=$2, language=$3, 
	submission_status=$4 code=$5 submission_date=$6 updated_at=$7
	where id=$8 and deleted_at is null`
	_, err = tr.Exec(query, submission.UserId,
		submission.ProblemId, submission.Language,
		submission.SubmissionStatus, submission.Code,
		submission.SubmissionDate, time.Now(), submission.Id)

	return err
}

// Delete
func (s *Submissions) DeleteSubmission(submissionId string) error {
	tr, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update submissions set deleted_at=$1 where id=$2 and 
	deleted_at is null`
	_, err = tr.Exec(query, time.Now(), submissionId)

	return err
}
