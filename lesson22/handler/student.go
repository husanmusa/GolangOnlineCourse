package handler

import (
	"encoding/json"
	"io"
	"lesson22/model"
	"log"
	"net/http"
)

func (h *Handler) CreateStudent(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetStudent(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	resp, err := h.studentRepo.GetStudent(id)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
