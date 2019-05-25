package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbSessions = map[string]string{} // Key- Session Id  and value is user id
var dbUsers = map[string]user{}      // Key- User id , value- user

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure : true,
			//HttpOnly: true,
		}
		http.SetCookie(w, c)
	}
	// if user aleready exist then get user detail
	var u user
	// Get user id from map seesion using session id as key
	// if un, ok := dbSessions[c.Value]; ok {
	// 	u = dbUsers[un] // get user using user id as key
	// }

	// If user already not exist then get from form, when method is post
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u

	}
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	// Get Cookie
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	user1 := dbUsers[u]
	tpl.ExecuteTemplate(w, "bar.gohtml", user1)
}
