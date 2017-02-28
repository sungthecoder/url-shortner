package session

import (
  "net/http"
  "github.com/gorilla/securecookie"
  "fmt"
  "strconv"
)

var cookieHandler = securecookie.New(
  securecookie.GenerateRandomKey(64),
  securecookie.GenerateRandomKey(32),
)

func SetSession(username string, id int, res http.ResponseWriter) {
  value := map[string]string {
    "username": username,
    "id": fmt.Sprintf("%d", id),
  }

  if encoded, err := cookieHandler.Encode("session", value); err == nil {
    cookie := &http.Cookie{
      Name: "session",
      Value: encoded,
      Path: "/",
    }
    http.SetCookie(res, cookie)
  }
}

func ClearSession(res http.ResponseWriter) {
  cookie := &http.Cookie{
    Name: "session",
    Value: "",
    Path: "/",
    MaxAge: -1,
  }
  http.SetCookie(res, cookie)
}

func CurrentUser(req *http.Request) (int64, string) {
  var username string
  var id       int64
  if cookie, err := req.Cookie("session"); err == nil {
    cookieValue := make(map[string] string)
    if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
      username = cookieValue["username"]
      id, _    = strconv.ParseInt(cookieValue["id"], 10, 64)
    }
  }
  return id, username
}
