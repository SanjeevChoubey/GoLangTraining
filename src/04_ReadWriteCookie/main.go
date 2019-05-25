package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/write", write)
	http.HandleFunc("/read", read)
	http.ListenAndServe(":8080", nil)
}

func write(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some-value",
	})
	fmt.Fprintln(w, "Cookies writen in to browser- path- in chrome goto dev tools/application/cookies")
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintln(w, "Recived Cookie", c)
}
