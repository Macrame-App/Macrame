package main

import (
	"log"
	"mime"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// log.Println(r.URL.Path)

		file := "../public" + r.URL.Path

		if r.URL.Path == "/" {
			file = "../public/index.html"
			// } else if strings.HasSuffix(r.URL.Path, ".js") {
			// 	log.Println("js")
			// 	w.Header().Set("Content-Type", "application/javascript")
			// } else if strings.HasSuffix(r.URL.Path, ".css") {
			// 	log.Println("css")
			// 	w.Header().Set("Content-Type", "text/css")
			// }
		} else {
			contentType := mime.TypeByExtension(filepath.Ext(file))
			if contentType != "" {
				w.Header().Set("Content-Type", contentType)
			}
		}

		log.Println(file)
		log.Println("-------------")
		http.ServeFile(w, r, file)
	})

	log.Fatal(http.ListenAndServe(":6970", nil))
}
