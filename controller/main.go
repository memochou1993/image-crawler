package controller

import (
	"net/http"

	"github.com/memochou1993/image-crawler/crawler"
)

// Handler func
func Handler(w http.ResponseWriter, r *http.Request) {
	links := []string{
		// "https://www.104.com.tw/jobs/main/",
		"https://www.google.com/?1",
		"https://www.google.com/?2",
		"https://www.google.com/?3",
		"https://www.google.com/?4",
		"https://www.google.com/?5",
		"https://www.google.com/?6",
	}

	crawler.Initialize(links)
}
