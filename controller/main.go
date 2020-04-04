package controller

import (
	"net/http"

	"github.com/memochou1993/image-crawler/crawler"
)

// Handler func
func Handler(w http.ResponseWriter, r *http.Request) {
	links := []string{
		"https://www.google.com/?1",
		"https://www.google.com/?2",
		"https://www.google.com/?3",
		"https://www.google.com/?4",
		"https://www.google.com/?5",
		"https://www.google.com/?6",
		"https://www.google.com/?7",
		"https://www.google.com/?8",
		"https://www.google.com/?9",
		"https://www.google.com/?10",
	}

	crawler.Initialize(links)
}
