package postgres

import (
	"database/sql"
	"leetcode/model"
	"time"
)

type Languages struct {
	DB *sql.DB
}

func NewLanguagesRepo(db *sql.DB) *Languages {
	return &Languages{DB: db}
}

// Create
func (l *Languages) CreateLanguage(language *model.Language) error {
	tr, err := l.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	query := `insert into languages(name) values($1)`
	_, err = tr.Exec(query, language.Name)
	return err
}

// Read All
func (l *Languages) GetLanguage(filter *model.FilterLanguage) ([]model.Language, error) {
	languages := []model.Language{}
	argums := []interface{}{}
	query := `select * from languages where deleted_at is null`
	if filter.Name != nil {
		query += ` and name=$1`
		argums = append(argums, filter.Name)
	}

	rows, err := l.DB.Query(query, argums...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		language := model.Language{}
		err := rows.Scan(&language.Id, &language.Name,
			&language.Created_at, &language.Updated_at,
			&language.Deleted_at)
		if err != nil {
			return nil, err
		}
		languages = append(languages, language)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return languages, nil
}

// Read by Id
func (l *Languages) GetLanguageById(languageId string) (model.Language, error) {
	language := model.Language{}
	query := `select * from languages where id=$1 and deleted_at is null`
	row := l.DB.QueryRow(query, languageId)
	err := row.Scan(&language.Id, &language.Name,
		&language.Created_at, &language.Updated_at,
		&language.Deleted_at)
	return language, err
}

// Update
func (l *Languages) UpdateLanguage(language *model.Language) error {
	tr, err := l.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update languages set name=$1, updated_at=$2
	where id=$3 and deleted_at is null`
	_, err = tr.Exec(query, language.Name, time.Now(), language.Id)

	return err
}

// Delete
func (l *Languages) DeleteLanguage(languageId string) error {
	tr, err := l.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update languages set deleted_at=$1 where id=$2 and deleted_at is null`
	_, err = tr.Exec(query, time.Now(), languageId)

	return err
}
