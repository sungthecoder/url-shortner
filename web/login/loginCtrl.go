package login

import (
  "urlShort/web/session"
  "urlShort/web/user"

  "fmt"
  "net/http"
  "html/template"
)

func HandleRequest(res http.ResponseWriter, req *http.Request) {
  if req.URL.Path == "/logout" {
    destroy(res, req)

  } else {
    switch req.Method {
      case "GET":  index(res, req)
      case "POST": create(res, req)

      default:    index(res, req)
    }
  }
}

func index(res http.ResponseWriter, req *http.Request) {
  fmt.Println("Login index")
  tmpl := "web/login/index.html"
  var x map[string] string

  t, _ := template.ParseFiles(tmpl)
  t.Execute(res, x)
}

func create(res http.ResponseWriter, req *http.Request) {
  username := req.FormValue("username")
  password := req.FormValue("password")
  fmt.Printf("Login create: %s(%s)\n", username, password)

  redirectTarget := "/login"
  id := user.FindUserId(username, password)
  if id > -1 {
     session.SetSession(username, id, res)
     redirectTarget = "/"
  }
  http.Redirect(res, req, redirectTarget, 302)
}

func destroy(res http.ResponseWriter, req *http.Request) {
  session.ClearSession(res)
  http.Redirect(res, req, "/login", 302)
}
