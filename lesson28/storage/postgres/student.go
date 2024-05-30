package postgres

import (
	"database/sql"
	"fmt"
	"students_and_courses/models"
)

type Students struct {
	db *sql.DB
}

func NewStudentsRepo(db *sql.DB) *Students {
	return &Students{db: db}
}

func (s *Students) CreateStudent(student *models.Student) error {
	query := fmt.Sprintf(`insert into students(id, name, surname, age)
	values ($1, $2, $3, $4)`)
	_, err := s.db.Exec(query, student.Id, student.Name, student.Surname, student.Age)
	return err
}

func (s *Students) GetAllStudents() (*[]models.Student, error) {
	students := &[]models.Student{}
	rows, err := s.db.Query("select * from students")
	if err != nil {
		return nil, err
	}
	var student models.Student
	for rows.Next() {
		err = rows.Scan(&student.Id, &student.Name, &student.Surname, &student.Age)
		if err != nil {
			return nil, err
		}
		*students = append(*students, student)
	}
	return students, rows.Err()
}

func (s *Students) UpdateStudent(student *models.Student) error {
	query := fmt.Sprintf(`update students set name=$1, surname=$2, age=$3
		where id=$4`)
	_, err := s.db.Exec(query, student.Name, student.Surname, student.Age, student.Id)
	return err
}

func (s *Students) DeleteStudent(id string) error {
	query := fmt.Sprintf(`delete from students where id=$1`)
	_, err := s.db.Exec(query, id)
	return err
}

// func (c *Students) UpdateCourse()            {}
// func (c *Students) DeleteCourse()            {}
// func (c *Students) PrintAllStudentsInfo()    {}
// func (c *Students) GetStudentsAverageGrade() {}
// func (c *Students) GetCoursesOfStudent()     {}
