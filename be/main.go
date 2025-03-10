package main

import (
	"log"
	"mime"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		file := "../public" // base directory

		if r.URL.Path == "/" {
			file = file + "/index.html" // default
		} else {
			file = file + r.URL.Path // request
		}

		contentType := mime.TypeByExtension(filepath.Ext(file)) // get content type

		if contentType != "" {
			w.Header().Set("Content-Type", contentType) // set content type header
		}

		log.Println(file)
		log.Println("-------------")

		http.ServeFile(w, r, file) // serve file
	})

	log.Fatal(http.ListenAndServe(":6970", nil))
}
