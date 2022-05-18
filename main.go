package main

import (
	"fmt"
	"net/http"
	"text/template"
)

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

	fmt.Println(name, email, mobile)

	fmt.Fprintf(w, `Received`) //response
}
