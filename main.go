package main

import (
	"net/http"

	web "ascii_art_web/web_server"
)

func main() {
	http.HandleFunc("/", web.Index)
	http.HandleFunc("/ascii-art", web.Generate)
	http.ListenAndServe(":8080", nil)
}
