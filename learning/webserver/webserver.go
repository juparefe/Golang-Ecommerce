package webserver

import (
	"net/http"
)

func MyWebServer() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./learning/webserver/index.html")
}
