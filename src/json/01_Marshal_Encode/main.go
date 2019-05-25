package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name      string
	Age       int
	Companies []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	s := `
	<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title> User data</title>
</head>


<body>
<h1> You are in foo</h1>
</body>


</html>

`
	w.Write([]byte(s))
}

func mshl(w http.ResponseWriter, r *http.Request) {
	v := person{
		Name:      "Sanjeev Choubey",
		Age:       34,
		Companies: []string{"Mindtree", "Manhattan", "Galaxe", "Amber Road"},
	}
	j1, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
	}

	w.Write(j1)

}

func encd(w http.ResponseWriter, r *http.Request) {

	v := person{
		Name:      "Sanjeev Choubey",
		Age:       34,
		Companies: []string{"Mindtree", "Manhattan", "Galaxe", "Amber Road"},
	}
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Println(err)
	}

}
