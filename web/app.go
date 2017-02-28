package web

import (
  "urlShort/web/db"
  "urlShort/web/home"
  "urlShort/web/login"
  "urlShort/web/user"
  "urlShort/web/url"
  "urlShort/web/forwarding"

  "net/http"
  "github.com/gorilla/mux"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

type WebApp struct {
  Router      *mux.Router
  Db          *sql.DB
}

func (app *WebApp) SetRoute() {
  app.Router.HandleFunc("/",      home.HandleRequest)
  app.Router.HandleFunc("/login",    login.HandleRequest).Methods("GET", "POST")
  app.Router.HandleFunc("/logout",   login.HandleRequest)
  app.Router.HandleFunc("/users",     user.HandleRequest)
  app.Router.HandleFunc("/users/new", user.HandleRequest)
  app.Router.HandleFunc("/url",       url.HandleRequest).Methods("POST")
  app.Router.HandleFunc("/g/{shortUrl}", forwarding.HandleRequest).Methods("GET")
  app.Router.PathPrefix("/assets/").Handler(http.FileServer(http.Dir("./")))
}

func (app *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  app.Router.ServeHTTP(w,r)
}

var App = func() *WebApp {
  app := &WebApp {
    Router:      mux.NewRouter(),
    Db:          db.Open(),
  }

  return app
}()
