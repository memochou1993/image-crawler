package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/memochou1993/image-crawler/crawler"
)

// Index func
func Index(w http.ResponseWriter, r *http.Request) {
	render(w, "index")
}

// Preview func
func Preview(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	gallery := crawler.Gallery{}
	gallery.Query(r.URL.Query().Get("links"))
	gallery.Fetch()

	response(w, http.StatusOK, gallery.Format())
}

// Download func
func Download(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	gallery := crawler.Gallery{}
	gallery.Query(r.URL.Query().Get("links"))
	gallery.Fetch()

	download(w, "images", gallery.Compress())
}

func render(w http.ResponseWriter, name string) {
	var tmpl = template.Must(template.ParseFiles("public/" + name + ".html"))

	tmpl.Execute(w, nil)
}

func response(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func download(w http.ResponseWriter, filename string, payload []byte) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", filename))
	w.Write(payload)
}
