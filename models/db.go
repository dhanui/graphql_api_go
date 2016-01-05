package models

import (
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() error {
  var err error
  db, err = sql.Open("mysql", "root:yeahyeah@/api_test?parseTime=true")
  if err != nil {
    return err
  }

  err = db.Ping()
  if err != nil {
    return err
  }

  return err
}
