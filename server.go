package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", handler)
	r.HandleFunc("/thala", handler)
	r.HandleFunc("/thala/verdict", verdict)

	port = ":" + port
	http.ListenAndServe(port, r)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./resources/templates/index.html"))
	tmpl.Execute(writer, "")
}

func verdict(writer http.ResponseWriter, request *http.Request) {
	str := request.FormValue("inputThala")
	tmpl := template.Must(template.ParseFiles("./resources/templates/blocks.html"))
	if len(str) != 7 {
		tmpl.ExecuteTemplate(writer, "moye-moye", nil)
	} else {
		tmpl.ExecuteTemplate(writer, "bole-koyal", nil)
	}
}
