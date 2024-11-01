package handlers

import (
  "net/http"
  "fmt"

  "hometasks/internal/services"
)

func AuthDevice(w http.ResponseWriter, r *http.Request) {
  // TODO: how to validate data types from the path?
  device := r.PathValue("device")
  _, err := services.AuthDevice("hello", "secret")

  if err != nil {
    fmt.Fprintf(w, "error for device %s, %s", device, err)
    return
  }

  fmt.Fprintf(w, "authing device %s", device)
}

func ListTasksForDevice(w http.ResponseWriter, r *http.Request) {
  // TODO: create a middleware to get the token and validate it
  location := r.PathValue("device")
  fmt.Fprintf(w, "tasks for location %s", location)
}
