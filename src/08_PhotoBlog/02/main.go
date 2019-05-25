package main

import (
	"html/template"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	http.HandleFunc("/index", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	// Append strings
	c = appendValue(w, c)
	// get strings in an array
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
		http.SetCookie(w, c)

	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	//Get Yor images, in this example we will just keep hardcoed image name
	i1 := "star.jpg"
	i2 := "disover.jpg"
	i3 := "amberroad.jpg"

	if !strings.Contains(c.Value, i1) {
		c.Value = c.Value + "|" + i1
	}
	if !strings.Contains(c.Value, i2) {
		c.Value = c.Value + "|" + i2
	}
	if !strings.Contains(c.Value, i3) {
		c.Value = c.Value + "|" + i3
	}

	http.SetCookie(w, c)
	return c
}
