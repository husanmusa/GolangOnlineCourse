package repository

import (
	"database/sql"

	"lesson19/model"

	"github.com/google/uuid"
)

func CreateStudent(db *sql.DB, student *model.CreateStudent) error {
	id := uuid.NewString()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	_, err = tx.Exec("insert into student (student_id, name, lastname, phone, age, grade, gender) values ($1, $2, $3, $4, $5, $6, $7)",
		id, student.Name, student.LastName, student.Phone, student.Age, student.Grade, student.Gender)
	if err != nil {
		return err
	}

	_, err = tx.Exec("insert into student_course(student_id, course_id) values ($1, $2)", id, student.CourseId)
	if err != nil {
		return err
	}

	_, err = tx.Exec("update course set number=number+1 where course_id=$1", student.CourseId)
	if err != nil {
		return err
	}

	return nil
}
