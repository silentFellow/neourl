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
	subRouter := http.NewServeMux()
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", subRouter))

	subRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	middlewareStack := logginMiddleware(router)

	s.server = &http.Server{
		Addr:    port,
		Handler: middlewareStack,
	}

	return s.server.ListenAndServe()
}
