package models

import (
  "encoding/json"
  "os"
  "fmt"
  "database/sql"

  "github.com/mattes/migrate/migrate"
  _ "github.com/go-sql-driver/mysql"
  _ "github.com/mattes/migrate/driver/mysql"
)

type DBConf struct {
  Host string `json:"host"`
  Port string `json:"port"`
  User string `json:"user"`
  Pass string `json:"pass"`
  Name string `json:"name"`
}

var db *sql.DB

func loadConnectionString() (connString string, err error) {
  fmt.Printf("Loading database configuration...\n")
  file, err := os.Open("config/database.json")
  if err != nil {
    return
  }
  decoder := json.NewDecoder(file)
  dbConf := DBConf{}
  err = decoder.Decode(&dbConf)
  if err != nil {
    return
  } else {
    fmt.Printf("Configuration loaded\n")
  }
  return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.Name), nil
}

func InitDBConnection() (err error) {
  connString, err := loadConnectionString()
  if err != nil {
    return
  }
  fmt.Printf("Testing database connection...\n")
  db, err = sql.Open("mysql", connString)
  if err != nil {
    return
  }
  err = db.Ping()
  if err != nil {
    return
  }
  fmt.Printf("Connection successful\n")
  return
}

func Migrate() (errors []error) {
  connString, err := loadConnectionString()
  if err != nil {
    return append(errors, err)
  }
  url := fmt.Sprintf("mysql://%s", connString)
  errors, _ = migrate.UpSync(url, "./migrations")
  return
}

func Rollback() (errors []error) {
  connString, err := loadConnectionString()
  if err != nil {
    return append(errors, err)
  }
  url := fmt.Sprintf("mysql://%s", connString)
  errors, _ = migrate.DownSync(url, "./migrations")
  return
}
