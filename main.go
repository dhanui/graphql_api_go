package main

import (
  "fmt"
  "net/http"
  "encoding/json"

  "./models"
  "./schema"
)

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
  username, password, ok := r.BasicAuth()
  if !ok || !models.AuthenticateUser(username, password) {
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
    return
  }
  if r.Method == "POST" {
    body := make([]byte, r.ContentLength)
    _, err := r.Body.Read(body)
    if err != nil {
      result := schema.ExecuteQuery(string(body[:]))
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
    fmt.Printf("Error establishing connection to database: %s\n", err.Error())
    return
  }
  http.HandleFunc("/graphql", graphqlHandler)
  err = http.ListenAndServe(":8080", nil)
  if err != nil {
    fmt.Printf("Error starting HTTP server: %s\n", err.Error())
  }
}
