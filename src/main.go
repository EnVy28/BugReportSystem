package main

import (
	"html/template"
	"net/http"
)

func main() {
	//Specify that localhost/css is where the css files can be accessed, files coming from app/css
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("app/css"))))
	//Specify that localhost/js is where the js files can be accessed, files coming from app/js
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("app/js"))))

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("app/login.html")
	t.Execute(w, r)
}
