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

package app

import (
	"be/app/helper"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

func ApiCORS(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	origin := r.Header.Get("Origin")

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	if strings.HasPrefix(r.Host, "192.168.") {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Accept-Language, Accept-Encoding")

	return w, r
}

func ApiGet(w http.ResponseWriter, r *http.Request) {
	root, err := filepath.Abs("../public")
	if err != nil {
		MCRMLog("ApiGet Abs Error: ", err)
		return
	}

	var file string

	if strings.Contains(r.URL.Path, "/config.js") {
		file = filepath.Join(root, "config.js")
		w.Header().Set("Content-Type", "text/javascript") // set content type header
	} else if r.URL.Path != "/" {
		file = filepath.Join(root, r.URL.Path)
	}

	contentType := mime.TypeByExtension(filepath.Ext(file)) // get content type

	if contentType == "" {
		file = filepath.Join(root, "index.html")
	}

	http.ServeFile(w, r, file)
}

func ApiPost(w http.ResponseWriter, r *http.Request) {
	access, data, err := helper.EndpointAccess(w, r)

	if !access || err != nil {
		MCRMLog("ApiPost EndPointAccess Error: ", err)
		return
	}

	if data != "" {
		ApiAuth(data, w, r)
		return
	}

	switch r.URL.Path {
	case "/macro/check":
		CheckMacro(w, r)
	case "/macro/record":
		SaveMacro(w, r)
	case "/macro/list":
		ListMacros(w, r)
	case "/macro/open":
		OpenMacro(w, r)
	case "/macro/delete":
		DeleteMacro(w, r)
	case "/macro/play":
		PlayMacro("", w, r)
	case "/device/server/ip":
		GetServerIP(w)
	case "/device/list":
		DeviceList(w, r)
	case "/device/access/check":
		DeviceAccessCheck(w, r)
	case "/device/access/request":
		DeviceAccessRequest(w, r)
	case "/device/link/ping":
		PingLink(w, r)
	case "/device/link/start":
		StartLink(w, r)
	case "/device/link/poll":
		PollLink(w, r)
	case "/device/link/remove":
		RemoveLink("", w, r)
	case "/device/handshake":
		Handshake(w, r)
	case "/panel/list":
		PanelList(w, r)
	case "/panel/get":
		GetPanel("", w, r)
	case "/panel/save/json":
		SavePanelJSON(w, r)
	}
}

func ApiAuth(data string, w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/macro/play":
		PlayMacro(data, w, r)
	case "/device/link/remove":
		RemoveLink(data, w, r)
	case "/panel/list":
		MCRMLog("Authenticated Panellist")
		PanelList(w, r)
	case "/panel/get":
		GetPanel(data, w, r)
	}
}
