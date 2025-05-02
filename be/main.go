/*
Macrame is a program that enables the user to create keyboard macros and button panels.
The macros are saved as simple JSON files and can be linked to the button panels. The panels can
be created with HTML and CSS.

Copyright (C) 2025 Jesse Malotaux

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

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

	app.InitSystray()

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
