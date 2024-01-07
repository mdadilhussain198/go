package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

    r.HandleFunc("/", handler)
	r.HandleFunc("/thala", handler)
	r.HandleFunc("/thala/verdict", verdict)

	http.ListenAndServe(":8081", r)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./resources/templates/index.html"))
	tmpl.Execute(writer, "")
}

func verdict(writer http.ResponseWriter, request *http.Request) {
	str := request.FormValue("inputThala")
	fmt.Fprintf(writer, "%s\n", str)
}
