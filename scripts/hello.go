package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(index))
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, e := template.ParseFiles("pages/index.html")
	if e != nil {
		http.Error(w, "Internal Server Errror", 500)
		log.Println(e.Error())
		return
	}

	data := struct {
		Title    string
		Subtitle string
	}{
		Title:    "My Page",
		Subtitle: "Subtitle for my page",
	}
	e = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(e.Error())
		return
	}
}
