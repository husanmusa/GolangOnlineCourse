package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Student struct {
	Name   string
	Age    int
	Group  int
	Gender string
}

func main() {
	s := Student{"John Doe", 25, 5344, "male"}

	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/create-student", "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	r, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(r))
}
