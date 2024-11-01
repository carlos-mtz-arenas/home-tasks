package main

import (
  "flag"
  "log"
  "fmt"

  "hometasks/api"
)

func main() {
  address := flag.String("listenaddress", ":8080", "Server address")
  flag.Parse()

  fmt.Println("Starting service")

  // TODO: set up the DB connection
  // TODO: share the db connection with models or services

  server := api.NewServer(*address)
  log.Fatal(server.Start())
}
