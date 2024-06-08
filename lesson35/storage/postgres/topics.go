package postgres

import (
	"database/sql"
	"leetcode/model"
	"time"
)

type Topics struct {
	DB *sql.DB
}

func NewTopicsRepo(db *sql.DB) *Topics {
	return &Topics{DB: db}
}

// Create
func (t *Topics) CreateTopic(topic *model.Topic) error {
	tr, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	query := `insert into topics(name) values($1)`
	_, err = tr.Exec(query, topic.Name)
	return err
}

// Read All
func (t *Topics) GetTopics(filter *model.FilterTopic) ([]model.Topic, error) {
	topics := []model.Topic{}
	argums := []interface{}{}
	query := `select * from topics where deleted_at is null`
	if filter.Name != nil {
		query += ` and name=$1`
		argums = append(argums, filter.Name)
	}

	rows, err := t.DB.Query(query, argums...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		topic := model.Topic{}
		err := rows.Scan(&topic.Id, &topic.Name,
			&topic.Created_at, &topic.Updated_at,
			&topic.Deleted_at)
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return topics, nil
}

// Read by Id
func (t *Topics) GetTopicById(topicId string) (model.Topic, error) {
	topic := model.Topic{}
	query := `select * from topics where id=$1 and deleted_at is null`
	row := t.DB.QueryRow(query, topicId)
	err := row.Scan(&topic.Id, &topic.Name,
		&topic.Created_at, &topic.Updated_at,
		&topic.Deleted_at)
	return topic, err
}

// Update
func (t *Topics) UpdateTopic(topic *model.Topic) error {
	tr, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update topics set name=$1, updated_at=$2
	where id=$3 and deleted_at is null`
	_, err = tr.Exec(query, topic.Name, time.Now(), topic.Id)

	return err
}

// Delete
func (t *Topics) DeleteTopic(topicId string) error {
	tr, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update topics set deleted_at=$1 where id=$2 and deleted_at is null`
	_, err = tr.Exec(query, time.Now(), topicId)

	return err
}
