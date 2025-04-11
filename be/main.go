package main

import (
	"net/http"

	"be/app"
	"be/app/helper"
)

func main() {
	app.MCRMLogInit()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		apiInit(w, r)
	})

	app.MCRMLog(http.ListenAndServe(":"+helper.EnvGet("MCRM__PORT"), nil))
}

func apiInit(w http.ResponseWriter, r *http.Request) {
	app.ApiCORS(w, r)

	if r.Method == "GET" {
		app.ApiGet(w, r)
	} else if r.Method == "POST" {
		app.ApiPost(w, r)
	}
}
