package repository

import (
  "testing"
  "os"
  "fmt"
  "flag"
)

var dbConfPath = flag.String("C", "../config/database.json", "path to database config file")
var migrationPath = flag.String("m", "../migrations", "path to migration directory")

func TestMain(m *testing.M) {
  flag.Parse()
  errors := Migrate(*dbConfPath, *migrationPath, true)
  if len(errors) > 0 {
    fmt.Printf("Errors:\n")
    for i := 0; i < len(errors); i++ {
      fmt.Printf("* %s\n", errors[i].Error())
    }
    os.Exit(1)
  }
  err := InitDbConnection(*dbConfPath, true)
  if err != nil {
    fmt.Printf("Error: %s\n", err.Error())
    os.Exit(1)
  }
  res := m.Run()
  errors = Rollback(*dbConfPath, *migrationPath, true)
  if len(errors) > 0 {
    fmt.Printf("Errors:\n")
    for i := 0; i < len(errors); i++ {
      fmt.Printf("* %s\n", errors[i].Error())
    }
    os.Exit(1)
  }
  os.Exit(res)
}
