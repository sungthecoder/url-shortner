package main

import (
  "urlShort/web"
  "urlShort/web/db"

  "log"
  "net/http"
  "os"
)

const (
  PORT = ":8080"
)

func main() {

  web.App.SetRoute()

  if len(os.Args) > 1 && os.Args[1] == "--create-db" {
    db.Create(web.App.Db)
  }

  // Start serving
  log.Println("Starting server on:", PORT)

  http.Handle("/", web.App)
  log.Fatal(http.ListenAndServe(PORT, web.App))
}
