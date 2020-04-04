package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/memochou1993/image-crawler/crawler"
)

// Handler func
func Handler(w http.ResponseWriter, r *http.Request) {
	collection := crawler.Collection{}

	links := []string{
		"https://risu.io/",
		"https://www.104.com.tw/jobs/main/",
		"https://www.google.com/?1",
		// "https://www.google.com/?2",
		// "https://www.google.com/?3",
		// "https://www.google.com/?4",
		// "https://www.google.com/?5",
		// "https://www.google.com/?6",
		// "https://www.google.com/?7",
		// "https://www.google.com/?8",
		// "https://www.google.com/?9",
		// "https://www.google.com/?10",
	}

	collection.Fetch(links)

	images := collection.Format()

	fmt.Println("images", images)

	renderTemplate(w)
}

func renderTemplate(w http.ResponseWriter) {
	var tmpl = template.Must(template.ParseFiles("views/index.html"))

	tmpl.Execute(w, struct {
		Langs []string
	}{
		[]string{"Python", "Ruby", "PHP", "Java", "Golang"},
	})
}
