package postgres

import (
	"database/sql"
	"leetcode/model"
	"time"
)

type Users struct {
	DB *sql.DB
}

func NewUsersRepo(db *sql.DB) *Users {
	return &Users{DB: db}
}

// Create
func (u *Users) CreateUser(user *model.User) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	query := `insert into users(fullname,username,bio) values($1,$2,$3)`
	_, err = tr.Exec(query, user.Fullname, user.Username, user.Bio)
	return err
}

// Read All
func (u *Users) GetUsers(filter *model.FilterUser) ([]model.User, error) {
	users := []model.User{}
	argums := []interface{}{}
	query := `select * from users where deleted_at is null`
	num := "1"
	if filter.Fullname != nil {
		query += ` and fullname=$1`
		num = "2"
		argums = append(argums, filter.Fullname)
	}
	if filter.Username != nil {
		query += ` and username=$` + num
		argums = append(argums, filter.Username)
	}

	rows, err := u.DB.Query(query, argums...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Fullname, &user.Username, &user.Bio,
			&user.Created_at, &user.Updated_at, &user.Deleted_at)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return users, nil
}

// Read by Id
func (u *Users) GetUserById(userId string) (model.User, error) {
	user := model.User{}
	query := `select * from users where id=$1 and deleted_at is null`
	row := u.DB.QueryRow(query, userId)
	err := row.Scan(&user.Id, &user.Fullname, &user.Username, &user.Bio,
		&user.Created_at, &user.Updated_at, &user.Deleted_at)
	return user, err
}

// Update
func (u *Users) UpdateUser(user *model.User) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update users set fullname=$1, username=$2, bio=$3, updated_at=$4
	where id=$5 and deleted_at is null`
	_, err = tr.Exec(query, user.Fullname, user.Username, user.Bio, time.Now(),
		user.Id)

	return err
}

// Delete
func (u *Users) DeleteUser(userId string) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update users set deleted_at=$1 where id=$2 and deleted_at is null`
	_, err = tr.Exec(query, time.Now(), userId)

	return err
}
