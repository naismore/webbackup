package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	Title         string
	FeaturedPosts []featuredPostData
}

type featuredPostData struct {
	Title    string
	Subtitle string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	data := indexPage{
		Title:         "Escape",
		FeaturedPosts: featuredPosts(),
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func featuredPosts() []featuredPostData {
	return []featuredPostData{
		{
			Title:    "The Road Ahead",
			Subtitle: "The road ahead might be paved - it might not be.",
		},
		{
			Title:    "From Top Down",
			Subtitle: "Once a year.",
		},
	}
}
