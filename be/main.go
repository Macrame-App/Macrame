package main

import (
	"log"
	"net/http"

	"be/app"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		app.ApiCORS(w, r)

		if r.Method == "GET" {
			app.ApiGet(w, r)
		} else if r.Method == "POST" {
			app.ApiPost(w, r)
		}

	})

	// helper.OpenBrowser("http://localhost:6970")

	log.Fatal(http.ListenAndServe(":6970", nil))
}
