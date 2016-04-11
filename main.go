package main

import (
  "fmt"
  "flag"
  "net/http"
  "encoding/json"

  "github.com/dhanui/graphql_api_go/repository"
  "github.com/dhanui/graphql_api_go/schema"
)

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
  username, password, ok := r.BasicAuth()
  var user repository.User
  if ok {
    user, ok = repository.AuthenticateUser(username, password)
  }
  if !ok {
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
    return
  }
  if r.Method == "POST" {
    body := make([]byte, r.ContentLength)
    _, err := r.Body.Read(body)
    if err != nil {
      result := schema.ExecuteQuery(string(body[:]), user)
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(result)
    } else {
      http.Error(w, "Bad Request", http.StatusBadRequest)
    }
  } else {
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
  }
}

func printErrors(errors []error) {
  for i := 0; i < len(errors); i++ {
    fmt.Printf("* %s\n", errors[i].Error())
  }
}

var dbConfPath = flag.String("config", "./config/database.json", "path to database config file")
var migrationPath = flag.String("m", "./migrations", "path to migration directory")

func main() {
  flag.Parse()
  args := flag.Args()
  if len(args) == 1 {
    switch args[0] {
    case "migrate":
      errors := repository.Migrate(*dbConfPath, *migrationPath, false)
      if len(errors) > 0 {
        fmt.Println("Migration errors:")
        printErrors(errors)
      }
    case "rollback":
      errors := repository.Rollback(*dbConfPath, *migrationPath, false)
      if len(errors) > 0 {
        fmt.Println("Rollback errors:")
        printErrors(errors)
      }
    case "server":
      err := repository.InitDbConnection(*dbConfPath, false)
      if (err != nil) {
        fmt.Printf("Error initializing database connection: %s\n", err.Error())
        return
      }
      http.HandleFunc("/graphql", graphqlHandler)
      fmt.Println("HTTP server listening on port 8080...")
      err = http.ListenAndServe(":8080", nil)
      if err != nil {
        fmt.Printf("Error starting HTTP server: %s\n", err.Error())
      }
    default:
      fmt.Printf("Unknown command: %s\n", args[0])
    }
  } else {
    fmt.Println("Usage: graphql_api_go [flags] [command]")
    fmt.Println()
    fmt.Println("Flags:")
    fmt.Println("  -config string    path to database config file (default: ./config/database.json)")
    fmt.Println("  -m string         path to migration directory (default: ./migrations)")
    fmt.Println()
    fmt.Println("Commands:")
    fmt.Println("  migrate           migrate database")
    fmt.Println("  rollback          rollback database")
    fmt.Println("  server            start HTTP server")
  }
}
