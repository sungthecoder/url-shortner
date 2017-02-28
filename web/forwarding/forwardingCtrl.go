package forwarding

import (
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
  "urlShort/web/url"
)

func HandleRequest(res http.ResponseWriter, req *http.Request) {
  redirect(res, req)
}

func redirect(res http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  shortUrl := vars["shortUrl"]
  id, longUrl  := url.FindUrl(shortUrl)

  fmt.Println(longUrl)

  if longUrl == "" {
    // 404
    fmt.Println("404")
    http.NotFound(res, req)
  } else {
    url.UpdateTimeStamp(id)
    http.Redirect(res, req, longUrl, 301)
  }
}
