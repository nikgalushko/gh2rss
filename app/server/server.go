package server

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

type Server struct{}

func New() Server {
	return Server{}
}

func (s Server) Run(port int) error {
	router, err := s.routes()
	if err != nil {
		return err
	}

	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	return nil
}

func (s Server) routes() (http.Handler, error) {
	mux := http.NewServeMux()

	subFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		return nil, fmt.Errorf("failed to create sub filesystem: %w", err)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if r.URL.Path == "/" {
			http.ServeFileFS(w, r, subFS, "index.html")
			return
		}

		s.handler(w, r)
	})

	return mux, nil
}

func (s Server) handler(w http.ResponseWriter, r *http.Request) {}
