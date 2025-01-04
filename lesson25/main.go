package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func login(w http.ResponseWriter, r *http.Request) {
	// err := basicAuth(r.Header.Get("Authorization"))

	// if err != nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	w.Write([]byte("Unauthorized"))
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Authorized"))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User"))
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/user", basicAuth(getUser))

	http.ListenAndServe(":8080", nil)
}

func basicAuth(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("In Midlleware"))
		// f(w, r)

		key := r.Header.Get("Authorization")

		slc := strings.Split(key, " ")
		if len(slc) != 2 && slc[0] != "Basic" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized 1"))
			return
		}

		buffer, err := base64.StdEncoding.DecodeString(slc[1])
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized 2"))
			return
		}

		slc = strings.Split(string(buffer), ":")
		if len(slc) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized 3"))
			return
		}

		if slc[0] != "foo" || slc[1] != "bar" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized 4"))
			return
		}

		f(w, r)
	}
}
