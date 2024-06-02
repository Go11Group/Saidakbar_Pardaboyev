package main

import (
	"users_and_produncts/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// users := postgres.NewUsersRepo(db)

	// // Randomly Generate users
	// usersGroup := generator.GenerateUsers(15)
	// for _, user := range usersGroup {
	// 	err := users.CreateUser(user)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// // Search users with giving filter for it
	// id := 5
	// username := "ljohnson45"
	// filter := model.Filter{
	// 	Id:       &id,
	// 	Username: &username,
	// }
	// usersGroup, err := users.GetUser(filter)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, user := range *usersGroup {
	// 	fmt.Printf("%+v\n", user)
	// }

	// // update a user's information
	// usersGroup, err := users.GetUser(model.Filter{})
	// if err != nil {
	// 	panic(err)
	// }
	// user := (*usersGroup)[0]
	// user.Username = "Saidakbar"
	// err = users.UpdateUser(&user)
	// if err != nil {
	// 	panic(err)
	// }

	// // Delete user from users table
	// usersGroup, err := users.GetUser(model.Filter{})
	// if err != nil {
	// 	panic(err)
	// }
	// user := (*usersGroup)[1]
	// err = users.DeleteUser(user)
	// if err != nil {
	// 	panic(err)
	// }

	// products := postgres.NewProdectsRepo(db)

	// Randomly Generate products
	// ProductsGroup := generator.GenerateProducts(30)
	// for _, product := range ProductsGroup {
	// 	err := products.CreateProduct(product)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// // Search products table with giving filter for it
	// id := 15
	// name := "Smartwatch"
	// filter := model.FilterProduct{
	// 	Id:   &id,
	// 	Name: &name,
	// }
	// productsGroup, err := products.GetProduct(filter)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, produnct := range *productsGroup {
	// 	fmt.Printf("%+v\n", produnct)
	// }

	// // update a product's information
	// productsGroup, err := products.GetProduct(model.FilterProduct{})
	// if err != nil {
	// 	panic(err)
	// }
	// product := (*productsGroup)[0]
	// product.Name = "Laptop L"
	// err = products.UpdateProduct(&product)
	// if err != nil {
	// 	panic(err)
	// }

	// // Delete user from users table
	// productsGroup, err := products.GetProduct(model.FilterProduct{})
	// if err != nil {
	// 	panic(err)
	// }
	// product := (*productsGroup)[1]
	// err = products.DeleteProduct(&product)
	// if err != nil {
	// 	panic(err)
	// }
}

// userni productslarida create, update, delete qilinganda,
// transactiondan foydalanilsin. Siklda aylantirgan holda.

// Crudlar:
// user_products:(tableni o'zingiz yozing)
// Bu tableda UserProduct ma'lumotlari saqlanadi.
// Aynan Userni productlarini create, read update qilishda transactionlardan foydalaning.
