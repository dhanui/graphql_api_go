package main

import (
  "fmt"
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

func main() {
  err := models.InitDBConnection()
  if (err != nil) {
    fmt.Printf("Error initializing database connection: %s\n", err.Error())
    return
  }
  http.HandleFunc("/graphql", graphqlHandler)
  fmt.Printf("Starting HTTP server on port 8080...\n")
  err = http.ListenAndServe(":8080", nil)
  if err != nil {
    fmt.Printf("Error starting HTTP server: %s\n", err.Error())
  }
}
