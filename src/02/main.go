package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", gopher)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func gopher(w http.ResponseWriter, r *http.Request) {
	var s string
	if r.Method == http.MethodPost {
		f, _, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// Read file
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	// Write Html from Go
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method = "POST" enctype = "multipart/form-data">
	<input type ="file" name ="q">
	<input type = "submit">
	</form>
	<br>
	`+s)
}
