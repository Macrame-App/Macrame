package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"be/app"
	"be/app/helper"
)

func main() {
	app.MCRMLogInit()

	switchToBeDir()

	if helper.EnvGet("MCRM__PORT") == "" {
		app.MCRMLog("Error: MCRM__PORT is not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		apiInit(w, r)
	})

	helper.OpenBrowser("http://localhost:" + helper.EnvGet("MCRM__PORT"))

	app.MCRMLog("Listening on http://localhost:" + helper.EnvGet("MCRM__PORT"))

	app.MCRMLog(http.ListenAndServe(":"+helper.EnvGet("MCRM__PORT"), nil))
}

func switchToBeDir() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if !strings.HasSuffix(cwd, "be") {
		err := os.Chdir("be")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func apiInit(w http.ResponseWriter, r *http.Request) {
	app.ApiCORS(w, r)

	if r.Method == "GET" {
		app.ApiGet(w, r)
	} else if r.Method == "POST" {
		app.ApiPost(w, r)
	}
}
