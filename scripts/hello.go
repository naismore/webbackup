package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	const port = ":3000"

	mux := http.NewServeMux()
	mux.HandleFunc("/home", index)

	log.Println("Start server " + port)
	e := http.ListenAndServe(port, mux)
	if e != nil {
		log.Fatal(e)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, e := template.ParseFiles("index.html")
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
	if e != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(e.Error())
		return
	}
}
