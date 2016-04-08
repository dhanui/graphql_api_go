package main

import (
  "fmt"
  "os"
  "net/http"
  "encoding/json"

  "github.com/dhanui/graphql_api_go/models"
  "github.com/dhanui/graphql_api_go/schema"
)

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
  username, password, ok := r.BasicAuth()
  var user models.User
  if ok {
    user, ok = models.AuthenticateUser(username, password)
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

func main() {
  if len(os.Args) == 2 {
    switch os.Args[1] {
    case "migrate":
      errors := models.Migrate()
      if len(errors) > 0 {
        fmt.Println("Migration errors:")
        printErrors(errors)
      } else {
        fmt.Println("Migration successful")
      }
    case "rollback":
      errors := models.Rollback()
      if len(errors) > 0 {
        fmt.Println("Rollback errors:")
        printErrors(errors)
      } else {
        fmt.Println("Rollback successful")
      }
    case "server":
      err := models.InitDBConnection()
      if (err != nil) {
        fmt.Printf("Error initializing database connection: %s\n", err.Error())
        return
      }
      http.HandleFunc("/graphql", graphqlHandler)
      fmt.Printf("HTTP server listening on port 8080...\n")
      err = http.ListenAndServe(":8080", nil)
      if err != nil {
        fmt.Printf("Error starting HTTP server: %s\n", err.Error())
      }
    default:
      fmt.Printf("Unknown argument: %s\n", os.Args[1])
    }
  } else {
    fmt.Println("Usage: graphql_api_go [command]\n\nCommands:\n  migrate\tMigrate database\n  rollback\tRollback database\n  server\tStart HTTP server\n")
  }
}
