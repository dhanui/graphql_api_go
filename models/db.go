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

type DbParams struct {
  Host string `json:"host"`
  Port string `json:"port"`
  User string `json:"user"`
  Pass string `json:"pass"`
  Name string `json:"name"`
}

type DbConf struct {
  Db DbParams `json:"db"`
  TestDb DbParams `json:"test_db"`
}

var db *sql.DB

func loadConnectionString(dbConfPath string, test bool) (connString string, err error) {
  file, err := os.Open(dbConfPath)
  if err != nil {
    return
  }
  dbConf := DbConf{}
  err = json.NewDecoder(file).Decode(&dbConf)
  if err != nil {
    return
  }
  var dbParams DbParams
  if test {
    dbParams = dbConf.TestDb
  } else {
    dbParams = dbConf.Db
  }
  return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbParams.User, dbParams.Pass, dbParams.Host, dbParams.Port, dbParams.Name), nil
}

func InitDbConnection(dbConfPath string, test bool) (err error) {
  connString, err := loadConnectionString(dbConfPath, test)
  if err != nil {
    return
  }
  db, err = sql.Open("mysql", connString)
  if err != nil {
    return
  }
  err = db.Ping()
  if err != nil {
    return
  }
  return
}

func Migrate(dbConfPath string, migrationPath string, test bool) (errors []error) {
  connString, err := loadConnectionString(dbConfPath, test)
  if err != nil {
    return append(errors, err)
  }
  url := fmt.Sprintf("mysql://%s", connString)
  errors, _ = migrate.UpSync(url, migrationPath)
  return
}

func Rollback(dbConfPath string, migrationPath string, test bool) (errors []error) {
  connString, err := loadConnectionString(dbConfPath, test)
  if err != nil {
    return append(errors, err)
  }
  url := fmt.Sprintf("mysql://%s", connString)
  errors, _ = migrate.DownSync(url, migrationPath)
  return
}
