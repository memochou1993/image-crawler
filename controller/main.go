package controller

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/memochou1993/image-crawler/crawler"
	"github.com/memochou1993/image-crawler/formatter"
)

// Index func
func Index(w http.ResponseWriter, r *http.Request) {
	render(w, "index")
}

// Handle func
func Handle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	gallery := crawler.Gallery{}
	gallery.Query(r, "links")
	gallery.Fetch()

	payload := formatter.Payload{}
	payload.Set(gallery.Format())

	response(w, http.StatusOK, payload)
}

func response(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func render(w http.ResponseWriter, name string) {
	var tmpl = template.Must(template.ParseFiles("public/" + name + ".html"))

	tmpl.Execute(w, nil)
}
