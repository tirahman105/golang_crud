package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err = sql.Open("mysql", "root:drtareq@tcp(127.0.0.1:3306)/gocrud_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	//defer db.Close()

	fmt.Println("DB connection success")
}

func main() {
	fmt.Println(" Local server is running on port 8888")
	http.HandleFunc("/", home)
	http.HandleFunc("/new_member", new_member)
	// http.HandleFunc("/add", add)
	// http.HandleFunc("/dashboard", dashboard)
	// http.HandleFunc("/mlist", mlist)
	//built in file server
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	// fmt.Fprint(w, `welcome`)
}

// func add(w http.ResponseWriter, r *http.Request) {
// 	ptmp, err := template.ParseFiles("template/base.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	ptmp, err = template.ParseFiles("wpage/add.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	ptmp.Execute(w, nil)
// 	// fmt.Fprint(w, `welcome`)
// }

// func dashboard(w http.ResponseWriter, r *http.Request) {
// 	ptmp, err := template.ParseFiles("template/base.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	ptmp.Execute(w, nil)
// 	// fmt.Fprint(w, `welcome`)
// }

// func mlist(w http.ResponseWriter, r *http.Request) {
// 	ptmp, err := template.ParseFiles("template/base.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	ptmp.Execute(w, nil)
// 	// fmt.Fprint(w, `welcome`)
// }

func new_member(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	email := r.FormValue("email")
	mobile := r.FormValue("mobile")

	//fmt.Println(name, email, mobile)

	fmt.Fprintf(w, `Data inserted successfully`) //response

	// func new_member(w http.ResponseWriter, r *http.Request) {
	// 	r.ParseForm()

	// 	for key, val := range r.Form {
	// 		fmt.Println(key, val)
	// 	}

	qs := "INSERT INTO `members` (`id`, `name`, `email`, `mobile`, `status`) VALUES (NULL, '%s', '%s', '%s', '1');"

	sql := fmt.Sprintf(qs, name, email, mobile)

	//fmt.Println(sql)
	insert, err := db.Query(sql)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	//fmt.Fprintf(w, `ok`)

}
