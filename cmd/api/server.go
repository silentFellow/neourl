package api

import (
	"fmt"
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

	router.HandleFunc("/", HandleIndex)
	router.HandleFunc("/shorten-url", HandleUrlShorten(s.urlStorage))
	router.HandleFunc("/r/{encoded}", HandleUrlRedirection(s.urlStorage))

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
