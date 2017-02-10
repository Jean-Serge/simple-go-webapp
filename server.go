package main

import (
  "fmt"
  "html"
  "log"
  "net/http"
  "github.com/jean-serge/simple-webapp/db"
)


func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Tmp(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
  const PORT = 8080
  db.ConnectDB()
  router := NewRouter()

  log.Printf("Server listening on port %d", PORT)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), router))
}
