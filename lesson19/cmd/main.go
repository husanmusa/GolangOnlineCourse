package main

import (
	"log"

	"lesson19/model"
	"lesson19/postgres"
	"lesson19/repository"
)

func main() {
	// course, student, course_student. 
	// Tutor, group

	// API: CRUD course, student, Tutor, group 

	// additional: best student by course, best student in group, the biggest group, the biggest course, the oldest student, the youngest student, 
	// List of female/male students by course, group, tutor, 

	// additional for transaction: change tutor of group and course. change course/group of student.

	// each method should add own object.

	db, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	student := model.CreateStudent{
		Name:     "John",
		LastName: "Doe",
		Phone:    "1234567890",
		Age:      20,
		Grade:    "A",
		Gender:   "M",
		CourseId: "d8139e0a-d029-43cd-95b7-6d51b6b3f10d",
	}

	err = repository.CreateStudent(db, &student)
	if err != nil {
		log.Fatal(err)
	}
}
