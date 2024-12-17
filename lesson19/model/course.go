package model

import "time"

type Course struct {
	Id        string
	Name      string
	Number    int
	StartedAt time.Time
	Tutor     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
