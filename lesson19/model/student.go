package model

import "time"

type Student struct {
	Id        string
	Name      string
	LastName  string
	Phone     string
	Age       int
	Grade     string
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateStudent struct {
	Name      string
	LastName  string
	Phone     string
	Age       int
	Grade     string
	Gender    string
	CourseId  string
}
