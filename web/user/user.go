package user

import (
  "urlShort/web/db"
  "fmt"
  _ "github.com/mattn/go-sqlite3"
  "log"
)

func Authenticate(username string, password string) bool{
  return username != "" && password != "" && FindUserId(username, password) > -1
}

func Create(username, password string) error {
  fmt.Printf("Creating a User: %s(%s)\n", username, password)

  database := db.Open()
  tx, err := database.Begin()
  if err != nil {
    log.Fatal(err)
  }

  stmt, err := tx.Prepare("insert into user(username, password) values(?, ?)")
  if err != nil {
     log.Fatal(err)
  }
  defer stmt.Close()

  _, err = stmt.Exec(username, password)
  if err != nil {
     log.Fatal(err)
  }

  return tx.Commit()
}

func FindUserId(username, password string) int{
  var id int = -1

  database := db.Open()
  rows, err := database.Query("select id from user where username = ? and password = ?", username, password)
  if err != nil {
     log.Fatal(err)
  }
  defer rows.Close()

  if rows.Next() {
    err := rows.Scan(&id)
    if err != nil {
      log.Fatal(err)
    }
  }
  err = rows.Err()
  if err != nil {
     log.Fatal(err)
  }
  return id
}
