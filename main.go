package main

import (
	"log"
	"net/http"

	"github.com/memochou1993/image-crawler/controller"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public/assets"))))
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/api", controller.Handle)

	log.Fatal(http.ListenAndServe(":8084", nil))
}
