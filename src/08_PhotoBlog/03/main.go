package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", fs))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)

	if r.Method == http.MethodPost {
		mf, fh, err := r.FormFile("myfile")
		if err != nil {
			log.Fatalln(err)
		}
		defer mf.Close()

		//Create sha1 for file name

		ext := strings.Split(fh.Filename, ".")[1] // get the file extension
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext // file name

		//Create new file
		wd, err := os.Getwd() // get working directory
		if err != nil {
			log.Fatalln(err)
		}
		path := filepath.Join(wd, "public", "pics", fname) // create file path

		newfile, err := os.Create(path) // create new file
		if err != nil {
			log.Fatalln(err)
		}
		defer newfile.Close()
		mf.Seek(0, 0)
		// Append strings
		c = appendValue(w, c, fname)
	}

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

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {

	if !strings.Contains(c.Value, fname) {
		c.Value = c.Value + "|" + fname
	}

	http.SetCookie(w, c)
	return c
}
