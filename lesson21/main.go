package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))
	fmt.Println(r.RemoteAddr)
	fmt.Println(r.Method)
	fmt.Println(r.URL, r.URL.Path)
	fmt.Println(r.Header)

	fmt.Fprintf(w, "Hello, World!")
}

type Student struct {
	Name   string
	Age    int
	Group  int
	Gender string
}

func craeteStudent(w http.ResponseWriter, r *http.Request) {
	s := Student{}
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(s)


	fmt.Fprintf(w, "Student created!")
}

func getStudentByGender(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.PathValue("id"))
	
	fmt.Print(r.URL.Query())
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/create-student", craeteStudent)
	http.HandleFunc("/get-student/{id}", getStudentByGender)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
