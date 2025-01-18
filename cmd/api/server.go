package api

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/silentFellow/neourl/internal/urlcoder"
)

type server struct {
	addr       string
	server     *http.Server
	urlStorage *urlcoder.Storage
}

func NewServer(addr string, urlStorage *urlcoder.Storage) *server {
	return &server{
		addr:       addr,
		urlStorage: urlStorage,
	}
}

func (s *server) Run() error {
	port := fmt.Sprintf(":%v", s.addr)

	router := http.NewServeMux()

	// Serve static files with the correct handler
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/", http.StripPrefix("/static", fs))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	router.HandleFunc("/shorten-url", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			originalURL := r.FormValue("url")
			shortenedURL := fmt.Sprintf(
				"%v/r/%v",
				getCurrentURL(r),
				s.urlStorage.EncodeURL(originalURL),
			)

			response := fmt.Sprintf(
				"<a href='%s' target='_blank' data-clipboard='%s'>%s</a>",
				shortenedURL,
				shortenedURL,
				shortenedURL,
			)
			w.Write([]byte(response))
		}
	})

	router.HandleFunc("/r/{encoded}", func(w http.ResponseWriter, r *http.Request) {
		encoded := r.PathValue("encoded")

		decoded, err := s.urlStorage.DecodeURL(encoded)
		if err != nil {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, decoded, http.StatusTemporaryRedirect)
	})

	middlewareStack := createMiddlewareStack(
		loggingMiddleware,
		corsMiddleware,
	)

	s.server = &http.Server{
		Addr:    port,
		Handler: middlewareStack(router),
	}

	return s.server.ListenAndServe()
}

func getCurrentURL(r *http.Request) string {
	var currentURL strings.Builder

	if r.TLS != nil {
		currentURL.WriteString("https://")
	} else {
		currentURL.WriteString("http://")
	}

	currentURL.WriteString(r.Host)
	return currentURL.String()
}
