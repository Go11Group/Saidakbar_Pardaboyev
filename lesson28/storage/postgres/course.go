package postgres

import (
	"database/sql"
	"fmt"
	"students_and_courses/models"
)

type Courses struct {
	db *sql.DB
}

func NewCousesRepo(db *sql.DB) *Courses {
	return &Courses{db: db}
}

func (c *Courses) CreateCourse(course *models.Course) error {
	query := fmt.Sprintf(`insert into courses(id, name, description)
	values ($1, $2, $3)`)
	_, err := c.db.Exec(query, course.Id, course.Name, course.Description)
	return err
}

func (c *Courses) GetAllCourses() (*[]models.Course, error) {
	courses := &[]models.Course{}
	rows, err := c.db.Query("select * from courses")
	if err != nil {
		return nil, err
	}
	var course models.Course
	for rows.Next() {
		err = rows.Scan(&course.Id, &course.Name, &course.Description)
		if err != nil {
			return nil, err
		}
		*courses = append(*courses, course)
	}
	return courses, rows.Err()
}

func (c *Courses) UpdateCourse(course *models.Course) error {
	query := fmt.Sprintf(`update courses set name=$1, description=$2
		where id=$3`)
	_, err := c.db.Exec(query, course.Name, course.Description, course.Id)
	return err
}

func (c *Courses) DeleteCourse(id string) error {
	query := fmt.Sprintf(`delete from courses where id=$1`)
	_, err := c.db.Exec(query, id)
	return err
}

// func (c *Courses) PrintAllCoursesInfo()             {}
// func (c *Courses) GetTopStudentsOfPerCourses()      {}
// func (c *Courses) GetAverageGradeOfPerCourses()     {}
// func (c *Courses) GetYoungestStudentsOfPerCourses() {}
// func (c *Courses) GetTheBestGroup()                 {}
