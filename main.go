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
  if r.Method == "GET" {
    result := schema.ExecuteQuery(r.URL.Query()["query"][0])
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
  } else if r.Method == "POST" {
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
    fmt.Printf("Error establishing connection to database: %v\n", err)
    return
  }
  http.HandleFunc("/graphql", graphqlHandler)
  fmt.Printf("HTTP listening to port 8080\n")
  http.ListenAndServe(":8080", nil)
}
