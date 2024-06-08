package postgres

import (
	"database/sql"
	"fmt"
	"leetcode/model"
	"time"
)

type Problems struct {
	DB *sql.DB
}

func NewProblemsRepo(db *sql.DB) *Problems {
	return &Problems{DB: db}
}

// Create
func (p *Problems) CreateProblem(problem *model.Problem) error {
	tr, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	query := `insert into problems(title,status,description,
		examples,constraints) values($1,$2,$3,$4,$5)`
	_, err = tr.Exec(query, problem.Title, problem.Status, problem.Description,
		problem.Examples, problem.Constraints)
	return err
}

// Read All
func (p *Problems) GetProblems(filter *model.FilterProduct) ([]model.Problem, error) {
	problems := []model.Problem{}
	argums := []interface{}{}
	query := `select * from problems where deleted_at is null`
	num := 1
	if filter.Problem_num != nil {
		query += fmt.Sprintf(` and problem_num=$%d`, num)
		argums = append(argums, filter.Problem_num)
		num++
	}
	if filter.Title != nil {
		query += fmt.Sprintf(` and title=$%d`, num)
		argums = append(argums, filter.Title)
		num++
	}
	if filter.Status != nil {
		query += fmt.Sprintf(` and status=$%d`, num)
		argums = append(argums, filter.Status)
	}

	rows, err := p.DB.Query(query, argums...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		problem := model.Problem{}
		err := rows.Scan(&problem.Id, &problem.Problem_num, &problem.Title,
			&problem.Status, &problem.Description, &problem.Examples,
			&problem.Constraints, &problem.Created_at, &problem.Updated_at,
			&problem.Deleted_at)
		if err != nil {
			return nil, err
		}
		problems = append(problems, problem)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return problems, nil
}

// Read by Id
func (p *Problems) GetProblemById(ProblemId string) (model.Problem, error) {
	problem := model.Problem{}
	query := `select * from problems where id=$1 and deleted_at is null`
	row := p.DB.QueryRow(query, ProblemId)
	err := row.Scan(&problem.Id, &problem.Problem_num, &problem.Title,
		&problem.Status, &problem.Description, &problem.Examples,
		&problem.Constraints, &problem.Created_at, &problem.Updated_at,
		&problem.Deleted_at)
	return problem, err
}

// Update
func (p *Problems) UpdateProblem(problem *model.Problem) error {
	tr, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update problems set problem_num=$1, title=$2, status=$3, 
	description=$4 examples=$5 constraints=$6 updated_at=$7
	where id=$8 and deleted_at is null`
	_, err = tr.Exec(query, problem.Problem_num, problem.Title,
		problem.Status, problem.Description, problem.Examples,
		problem.Constraints, time.Now(), problem.Id)

	return err
}

// Delete
func (p *Problems) DeleteProblem(ProblemId string) error {
	tr, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update problems set deleted_at=$1 where id=$2 and deleted_at is null`
	_, err = tr.Exec(query, time.Now(), ProblemId)

	return err
}
