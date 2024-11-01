package api

import (
  "net/http"
  "fmt"

  "hometasks/internal/handlers"
)

type Server struct {
  address string
}

func NewServer(address string) *Server {
  return &Server {
    address: address,
  }
}

func (s *Server) Start() error {
  mux := http.NewServeMux()
  // fake endpoints :v
  mux.HandleFunc("GET /samples/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    fmt.Fprintf(w, "hello %s", id)
  })

  // REAL endpoints
  mux.HandleFunc("POST /devices/{device}/auth", handlers.AuthDevice)
  mux.HandleFunc("GET /devices/{device}/tasks", handlers.ListTasksForDevice)

  return http.ListenAndServe(s.address, mux)
}

