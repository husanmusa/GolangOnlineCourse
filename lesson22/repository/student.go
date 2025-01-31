package repository

import (
	"database/sql"

	"lesson22/model"

	"github.com/google/uuid"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) CreateStudent(student *model.CreateStudent) error {
	id := uuid.NewString()

	tx, err := r.db.Begin()
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

	_, err = tx.Exec("update course set number=number+1 where id=$1", student.CourseId)
	if err != nil {
		return err
	}

	return nil
}

func (r *StudentRepository) GetStudent(id string) (*model.GetStudentResp, error) {
	student := &model.GetStudentResp{}
	err := r.db.QueryRow(`
	select s.name, s.lastname, s.phone, s.age, s.grade, s.gender, c.name, c.number, c.tutor
	from student s
	 left join student_course sc on s.student_id = sc.student_id 
	 left join course c on sc.course_id = c.id
	where s.student_id = $1`, id).Scan(
		&student.Name,
		&student.LastName,
		&student.Phone,
		&student.Age,
		&student.Grade,
		&student.Gender,
		&student.Course.Name,
		&student.Course.Number,
		&student.Course.Tutor,
	)
	if err != nil {
		return nil, err
	}
	return student, nil
}
