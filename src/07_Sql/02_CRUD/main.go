package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/go_schema")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)
	http.HandleFunc("/", index)
	http.HandleFunc("/person", persons)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func drop(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "Customer Table Dropped")
}

func delete(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DELETE from customer WHERE name = "GOOGLE";`)
	check(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)

}

func update(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer Set  name = "GOOGLE" WHERE name = "MICROSOFT";`)
	check(err)
	defer stmt.Close()
	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)
	fmt.Fprintln(w, "Number of Rows Updated:", n)

}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()
	var name string

	for rows.Next() {
		err := rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "Retrived records", name)
	}
}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES("MICROSOFT");`)
	check(err)
	defer stmt.Close()
	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Rows Inserted: ", n)
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()
	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)
	fmt.Fprintln(w, "CREATED TABLE customer", n)

}
func persons(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT FirstName FROM persons;`)
	check(err)
	defer rows.Close()
	// Read Records
	var s, name string
	s = "Retrived Records are:\n"
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"

	}
	fmt.Fprintln(w, s)
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Successfully Completed Db Connections")
	check(err)

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
