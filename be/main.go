package main

import (
	"log"
	"net/http"

	"be/app"
	"be/app/helper"
)

func main() {
	app.MCRMLogInit()

	if helper.EnvGet("MCRM__PORT") == "" {
		app.MCRMLog("Error: MCRM__PORT is not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		apiInit(w, r)
	})

	log.Println("Listening on http://localhost:" + helper.EnvGet("MCRM__PORT"))

	helper.OpenBrowser("http://localhost:" + helper.EnvGet("MCRM__PORT"))

	app.MCRMLog(http.ListenAndServe(":"+helper.EnvGet("MCRM__PORT"), nil))

}

func apiInit(w http.ResponseWriter, r *http.Request) {
	app.ApiCORS(w, r)

	app.MCRMLog("Remote IP: " + r.RemoteAddr)

	if r.Method == "GET" {
		app.ApiGet(w, r)
	} else if r.Method == "POST" {
		app.ApiPost(w, r)
	}
}
