package handler

import (
	"lesson22/repository"
	"net/http"
)

type Handler struct {
	studentRepo *repository.StudentRepository
	courseRepo  *repository.CourseRepository
}

func NewHandler(studentRepo *repository.StudentRepository, courseRepo *repository.CourseRepository) *Handler {
	return &Handler{
		studentRepo: studentRepo,
		courseRepo:  courseRepo,
	}
}

func Run(handler *Handler) *http.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /student", handler.CreateStudent)
	mux.HandleFunc("GET /student/{id}", handler.GetStudent)
	//mux.HandleFunc("PUT /student/{id}", handler.UpdateStudent)

	mux.HandleFunc("POST /course", handler.CreateCourse)

	server := &http.Server{Addr: ":8080", Handler: mux}
	return server

}
