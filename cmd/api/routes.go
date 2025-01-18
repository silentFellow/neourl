package api

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/silentFellow/neourl/internal/urlcoder"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/not-found.html"))
	tmpl.Execute(w, nil)
}

func HandleUrlShorten(storage *urlcoder.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			originalURL := r.FormValue("url")
			shortenedURL := fmt.Sprintf(
				"%v/r/%v",
				getCurrentURL(r),
				storage.EncodeURL(originalURL),
			)

			response := fmt.Sprintf(
				"<a href='%s' target='_blank' data-clipboard='%s'>%s</a>",
				shortenedURL,
				shortenedURL,
				shortenedURL,
			)
			w.Write([]byte(response))
		}
	}
}

func HandleUrlRedirection(storage *urlcoder.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		encoded := r.PathValue("encoded")

		decoded, err := storage.DecodeURL(encoded)
		if err != nil {
			http.Redirect(w, r, "/not-found", http.StatusFound)
			return
		}

		http.Redirect(w, r, decoded, http.StatusTemporaryRedirect)
	}
}
