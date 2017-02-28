package db

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "log"
)

var Db *sql.DB

func Open() *sql.DB {
  if Db == nil {
    db, err := sql.Open("sqlite3", "./app.db")
    if err != nil {
       log.Fatal(err)
    }
    //defer db.Close()
    Db = db
  }
  return Db
}

func Create(database *sql.DB) {
   log.Println("Creating DB...")

   sqlStmt := `
   create table if not exists user (id integer not null primary key, username text, password text);
   create table if not exists url  (id integer not null primary key,
      userId integer not null,
      name text not null,
      longUrl text not null,
      shortUrl text not null,
      createdAt datetime default current_timestamp,
      lastAccessed datetime default current_timestamp,
      FOREIGN KEY(userId) REFERENCES user(id)
    )
   `

   _, err := database.Exec(sqlStmt)
   if err != nil {
     log.Printf("%q: %s\n", err, sqlStmt)
     return
   }
}
