package api

import (
	"fmt"
	"html/template"
	"net/http"
)

type server struct {
	addr   string
	server *http.Server
}

func NewServer(addr string) *server {
	return &server{
		addr: addr,
	}
}

func (s *server) Run() error {
	port := fmt.Sprintf(":%v", s.addr)

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/v1/", http.StatusMovedPermanently)
	})

	subRouter := http.NewServeMux()
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", subRouter))

	subRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
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
