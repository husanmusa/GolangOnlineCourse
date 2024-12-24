package handler

import "net/http"

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
