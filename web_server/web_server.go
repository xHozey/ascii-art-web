package ascii_art_web

import (
	"html/template"
	"net/http"

	art "ascii_art_web/ascii_art"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if CheckUrl(w, r) {
		return
	}
	tmpl, err := template.ParseFiles("template/index.html")
	if checkServer(w, err) {
		return
	}
	tmpl.Execute(w, nil)
}

func Generate(w http.ResponseWriter, r *http.Request) {
	if CheckMethod(w, r) {
		return
	}
	tmpl, err := template.ParseFiles("template/ascii-art.html")
	if checkServer(w, err) {
		return
	}
	input := r.FormValue("text")
	filtredInput := CheckValidInput(input)
	banner := r.FormValue("banner")
	if CheckBanner(w, banner) {
		return
	}

	bannerPath := "ascii_art/templates/" + banner + ".txt"

	generated, err := art.ArgProcessor(filtredInput, bannerPath)
	if checkServer(w, err) {
		return
	}

	tmpl.Execute(w, generated)
}

func CheckValidInput(input string) string {
	filtredInput := ""
	for _, char := range input {
		if char == '\r' || char == '\n' {
			filtredInput += string(char)
		} else if int(char) < 32 || int(char) > 126 {
			continue
		} else {
			filtredInput += string(char)
		}
	}
	return filtredInput
}
