package postgres

import (
	"database/sql"
	"strconv"
	"time"
	"users_and_produncts/model"
)

type User_products struct {
	DB *sql.DB
}

func NewUser_ProductsRepo(db *sql.DB) *User_products {
	return &User_products{DB: db}
}

func (up *User_products) CreateUserProduct(user model.User, product model.Product) error {
	tr, err := up.DB.Begin()
	if err != nil {
		return err
	}

	defer tr.Commit()

	query := `insert into user_products(user_id, product_id) values ($1, $2)`
	_, err = tr.Exec(query, user.Id, product.Id)
	return err
}

func (up *User_products) GetUserProduct(filter model.FilterUserProduct) (*[]model.User_product, error) {
	query := `select * from user_products`

	isEmptyParams := true
	sch := 1
	params := []interface{}{}
	if filter.Id != nil {
		query += ` where`
		query += ` id=$` + strconv.Itoa(sch)
		params = append(params, filter.Id)
		isEmptyParams = false
		sch++
	}
	if filter.User_id != nil {
		if !isEmptyParams {
			query += ` and`
		} else {
			query += ` where`
		}
		query += ` user_id=$` + strconv.Itoa(sch)
		params = append(params, filter.User_id)
		isEmptyParams = false
		sch++
	}
	if filter.Product_id != nil {
		if !isEmptyParams {
			query += ` and`
		} else {
			query += ` where`
		}
		query += ` product_id=$` + strconv.Itoa(sch)
		params = append(params, filter.Product_id)
		isEmptyParams = false
		sch++
	}
	if isEmptyParams {
		query += ` where deletedat is null`
	} else {
		query += ` and deletedat is null`
	}

	userProducts := []model.User_product{}
	rows, err := up.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		userProduct := model.User_product{}
		err = rows.Scan(&userProduct.Id, &userProduct.User_id, &userProduct.Product_id, &userProduct.CreatedAt, &userProduct.UpdatedAt, &userProduct.DeletedAt)
		if err != nil {
			return nil, err
		}
		userProducts = append(userProducts, userProduct)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &userProducts, err
}

func (up *User_products) UpdateUserProduct(userProduct *model.User_product) error {
	tr, err := up.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update user_products set user_id=$1, product_id=$2, updatedat=$3 where id=$4`
	_, err = tr.Exec(query, userProduct.User_id, userProduct.Product_id, time.Now(), userProduct.Id)
	if err != nil {
		return err
	}
	return nil
}

func (up *User_products) DeleteUserProduct(userProduct *model.User_product) error {
	tr, err := up.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	query := `update user_products set user_id=$1, product_id=$2,
	deletedat=$3 where id=$4`
	_, err = tr.Exec(query, userProduct.User_id, userProduct.Product_id,
		time.Now(), userProduct.Id)
	if err != nil {
		return err
	}
	return nil
}
