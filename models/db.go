package models

import (
  "encoding/json"
  "os"
  "fmt"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
)

type DBConf struct {
  Host string `json:"host"`
  Port string `json:"port"`
  User string `json:"user"`
  Pass string `json:"pass"`
  Name string `json:"name"`
}

var db *sql.DB

func InitDB() error {
  file, err := os.Open("database.json")
  if err != nil {
    return err
  }
  decoder := json.NewDecoder(file)
  dbConf := DBConf{}
  err = decoder.Decode(&dbConf)
  if err != nil {
    return err
  }

  connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
    dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.Name)
  db, err = sql.Open("mysql", connString)
  if err != nil {
    return err
  }

  err = db.Ping()
  if err != nil {
    return err
  }

  return err
}
