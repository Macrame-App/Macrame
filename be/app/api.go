package app

import (
	"be/app/helper"
	"log"
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
	file := "" // base directory

	if r.URL.Path != "/" {
		file = "../public" + r.URL.Path // request
	}
	contentType := mime.TypeByExtension(filepath.Ext(file)) // get content type

	if contentType != "" {
		w.Header().Set("Content-Type", contentType) // set content type header
	}

	if contentType == "" {
		file = "../public/index.html" // default
	}

	// log.Println("GET:", file)

	http.ServeFile(w, r, file) // serve file
}

func ApiPost(w http.ResponseWriter, r *http.Request) {
	access, data := helper.EndpointAccess(w, r)

	if !access {
		return
	}

	log.Println("api post", data == "")
	if data != "" {
		ApiAuth(data, w, r)
		return
	}

	switch r.URL.Path {
	case "/macro/record":
		SaveMacro(w, r)
	case "/macro/list":
		ListMacros(w, r)
	case "/macro/delete":
		DeleteMacro(w, r)
	case "/macro/play":
		PlayMacro("", w, r)
	case "/device/server/ip":
		GetServerIP(w, r)
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
	log.Println("apiauth", data != "")
	switch r.URL.Path {
	case "/macro/play":
		PlayMacro(data, w, r)
	case "/device/link/remove":
		RemoveLink(data, w, r)
	case "/panel/get":
		GetPanel(data, w, r)
	}
}
