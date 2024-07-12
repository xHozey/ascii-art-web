package ascii_art_web

import (
	"net/http"
)

func CheckBanner(w http.ResponseWriter, s string) bool {
	if s != "standard" && s != "shadow" && s != "thinkertoy" && s != "graffiti" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return true
	}
	return false
}

func CheckUrl(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return true
	}
	return false
}

func CheckMethod(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "POST" {
		http.Error(w, "Methos Not Allowed", http.StatusMethodNotAllowed)
		return true
	}
	return false
}

func checkServer(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, " Internal Server Error", http.StatusInternalServerError)
		return true
	}
	return false
}
