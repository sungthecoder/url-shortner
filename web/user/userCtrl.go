package user

import (
  "fmt"
  "net/http"
  "html/template"
  "strings"
)

func HandleRequest(res http.ResponseWriter, req *http.Request) {
  paths := strings.Split(req.URL.Path, "/")
  lastPath := paths[len(paths)-1]

  if lastPath == "new" {
    newUser(res, req)
  } else {
    switch req.Method {
      case "POST":  create(res, req)
      default:      redirect(res, req)
    }
  }
}

func newUser(res http.ResponseWriter, req *http.Request) {
  fmt.Println("New User")
  tmpl := "web/user/new.html"

  t, _ := template.ParseFiles(tmpl)
  t.Execute(res, nil)
}

func create(res http.ResponseWriter, req *http.Request) {
  username := req.FormValue("username")
  password := req.FormValue("password")

  Create(username, password)
  redirect(res, req)
}

func redirect(res http.ResponseWriter, req *http.Request) {
  http.Redirect(res, req, "/login", 302)
}

