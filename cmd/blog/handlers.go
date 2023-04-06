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
	Preview     string
	Title       string
	Subtitle    string
	AuthorImage string
	Author      string
	PublishDate string
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
			Preview:     "featured-posts__first-featured-post_back-image",
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			AuthorImage: "../static/img/Mat-Vogels.jpg",
			Author:      "Mat Vogels",
			PublishDate: "September 25, 2015",
		},
		{
			Preview:     "featured-posts__second-featured-post_back-image",
			Title:       "From Top Down",
			Subtitle:    "Once a year.",
			AuthorImage: "../static/img/William-Wong.jpg",
			Author:      "William Wong",
			PublishDate: "September 25, 2015",
		},
	}
}
