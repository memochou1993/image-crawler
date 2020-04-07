package main

import (
	"log"
	"net/http"

	"github.com/memochou1993/image-crawler/controller"
)

func main() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/api", controller.Handle)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public/assets"))))

	log.Fatal(http.ListenAndServe(":8084", nil))
}
