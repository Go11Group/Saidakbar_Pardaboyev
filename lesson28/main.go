package main

import (
	"students_and_courses/storage/postgres"

	_ "github.com/lib/pq"
)

func main() {
	db, err := postgres.CreateDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// courses := postgres.NewCousesRepo(db)
	// students := postgres.NewStudentsRepo(db)

	// // Add New Course into courses table
	// newId := uuid.NewString()
	// name := "Python Backend"
	// description := "Learning Python Backend for earnging money"
	// newCourse := &models.Course{
	// 	Id:          newId,
	// 	Name:        name,
	// 	Description: description}
	// err = courses.CreateCourse(newCourse)
	// if err != nil {
	// 	panic(err)
	// }

	// // Get all Courses
	// coursesGroup, err := courses.GetAllCourses()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, val := range *coursesGroup {
	// 	fmt.Println("ID =", val.Id)
	// 	fmt.Println("Name =", val.Name)
	// 	fmt.Println("Description =", val.Description)
	// 	fmt.Println("-----------------------------------------------\n")
	// }

	// // Update course
	// newName := "NodeJS backend"
	// newCourse := &models.Course{
	// 	Id:          "6762ca14-ece0-49fc-93b5-c946ffb6305e", // not change
	// 	Name:        newName,
	// 	Description: "Learning Python Backend for earnging money", // not change
	// }
	// err = courses.UpdateCourse(newCourse)
	// if err != nil {
	// 	panic(err)
	// }

	// // Delete from courses table
	// id := "82080a0e-fffd-4746-9460-2af661b26b8d"
	// err = courses.DeleteCourse(id)
	// if err != nil {
	// 	panic(err)
	// }

	// // Add new student into students table
	// newId := uuid.NewString()
	// name := "Umar"
	// surname := "Xatob"
	// age := 45
	// NewStudent := &models.Student{
	// 	Id:      newId,
	// 	Name:    name,
	// 	Surname: surname,
	// 	Age:     age,
	// }
	// err = students.CreateStudent(NewStudent)
	// if err != nil {
	// 	panic(err)
	// }

	// // Get All Students in Students table
	// studentsGroup, err := students.GetAllStudents()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("-------------Students--------------")
	// for _, val := range *studentsGroup {
	// 	fmt.Println("ID =", val.Id)
	// 	fmt.Println("Name =", val.Name)
	// 	fmt.Println("Surname =", val.Surname)
	// 	fmt.Println("Age =", val.Age)
	// 	fmt.Println("-----------------------------------------------")
	// }

	// // Update student
	// newName := "Xadicha"
	// newAge := 30
	// newStudent := &models.Student{
	// 	Id:      "3b1b20d8-45e5-4707-b195-753222c3f431", // not change
	// 	Name:    newName,
	// 	Surname: "Sharapateva", // not change
	// 	Age:     newAge,
	// }
	// err = students.UpdateStudent(newStudent)
	// if err != nil {
	// 	panic(err)
	// }

	// Delete Student from students table
	// id := "d4811510-1e67-45b9-94ab-736979cd239a"
	// err = students.DeleteStudent(id)
	// if err != nil {
	// 	panic(err)
	// }
}
