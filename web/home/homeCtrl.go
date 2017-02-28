package home

import (
  "urlShort/web/session"
  "urlShort/web/url"
  "fmt"
  "net/http"
  "html/template"
)

func HandleRequest(res http.ResponseWriter, req *http.Request) {
  switch req.Method {
    case "GET":  index(res, req)

    default:    index(res, req)
  }
}

type UserVM struct {
  Username string
  URLs     []url.Url
}

func index(res http.ResponseWriter, req *http.Request) {
  fmt.Println("Home index")
  tmpl := "web/home/index.html"
  userId, username := session.CurrentUser(req)

  if username == "" {
    http.Redirect(res, req, "/login", 302)
  }

  urls := url.FindAll(userId)
  vm := new(UserVM)
  vm.Username = username
  vm.URLs     = urls

  t, _ := template.ParseFiles(tmpl)
  t.Execute(res, vm)
}

