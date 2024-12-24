package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"lesson22/model"
	"lesson22/postgres"
	"lesson22/repository"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	studentRepo := repository.NewStudentRepository(db)
	handler := NewHandler(studentRepo)

	http.HandleFunc("/student", handler.createStudent)

	http.ListenAndServe(":8080", nil)

}

type Handler struct {
	studentRepo *repository.StudentRepository
}

func NewHandler(studentRepo *repository.StudentRepository) *Handler {
	return &Handler{studentRepo: studentRepo}
}

func (h *Handler) createStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	student := model.CreateStudent{}

	err = json.Unmarshal(body, &student)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = h.studentRepo.CreateStudent(&student)
	if err != nil {
		log.Fatal(err)
	}
}
