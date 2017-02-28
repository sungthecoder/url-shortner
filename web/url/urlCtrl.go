package url

import (
  "net/http"
  "urlShort/web/session"
)

func HandleRequest(res http.ResponseWriter, req *http.Request) {
  create(res, req)
}

func create(res http.ResponseWriter, req *http.Request) {
  id, username := session.CurrentUser(req)
  if username == "" {
    http.Redirect(res, req, "/login", 302)
  }

  longURL := req.FormValue("url")

  Create(longURL, id)
  http.Redirect(res, req, "/", 302)
}
