package url

import (
  "urlShort/web/db"
  "fmt"
  _ "github.com/mattn/go-sqlite3"
  "log"
  "time"
  "net/url"
  "math/rand"
)

type Url struct {
  Name         string
  ShortUrl     string
  LongUrl      string
  CreatedAt    *time.Time
  LastAccessed *time.Time
}

func Create(url string, userId int64) error {
  fmt.Printf("Creating a short url: %s\n", url)

  shortUrl := shortenUrl(8)
  name     := getName(url)

  database := db.Open()
  tx, err := database.Begin()
  if err != nil {
    log.Fatal(err)
  }

  stmt, err := tx.Prepare("insert into url(userId, name, longUrl, shortUrl, createdAt) values(?, ?, ?, ?, ?)")
  if err != nil {
     log.Fatal(err)
  }
  defer stmt.Close()

  _, err = stmt.Exec(userId, name,  url, shortUrl, time.Now())
  if err != nil {
     log.Fatal(err)
  }

  return tx.Commit()
}

func shortenUrl(size int) string{

  rand.Seed(time.Now().UTC().UnixNano())
  const alphabet = "1234567890abcdefghjiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ$-_.+!*'(),'"
  const base     = len(alphabet)
  result := make([]byte, size)
  for i := 0; i < size; i++ {
     result[i] = alphabet[rand.Intn(base)]
  }
  return string(result)
}

func getName(longUrl string) string{
  u, err := url.Parse(longUrl)
  if err != nil {
     log.Fatal(err)
  }

  return u.Host
}

func FindUrl(shortUrl string) (int, string){
  var (
    id int
    longUrl string
  )

  database := db.Open()
  rows, err := database.Query("select id, longUrl from url where shortUrl=?", shortUrl)
  if err != nil {
     log.Fatal(err)
  }
  defer rows.Close()

  if rows.Next() {
    err := rows.Scan(&id, &longUrl)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(longUrl)
  }
  err = rows.Err()
  if err != nil {
     log.Fatal(err)
  }

  return id, longUrl
}

func FindAll(userId int64) []Url {
  var urls []Url
  database := db.Open()
  //rows, err := database.Query(
    //"select name, shortUrl, longUrl,createdAt, lastAccessed from url where userId = ?", userId)
  rows, err := database.Query(
    "select name, shortUrl, longUrl, createdAt, lastAccessed from url where userId = ?", userId)
  if err != nil {
     log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
    var u Url
    //err := rows.Scan(&u.name, &u.shortUrl, &u.longUrl, &u.createdAt, &u.lastAccessed)
    err := rows.Scan(&u.Name, &u.ShortUrl, &u.LongUrl, &u.CreatedAt, &u.LastAccessed)
    if err != nil {
      log.Fatal(err)
    }
    urls = append(urls, u)
  }

  err = rows.Err()
  if err != nil {
     log.Fatal(err)
  }

  return urls
}

func UpdateTimeStamp(id int) {
  database := db.Open()

  _, err := database.Exec("update url SET lastAccessed = ? WHERE id = ?", time.Now(), id)
  if err != nil {
    log.Fatal(err)
  }
}
