package main

import (
  "fmt"
  "net/http"
  "encoding/json"

  "./models"
  "./schema"
)

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
  result := schema.ExecuteQuery(r.URL.Query()["query"][0])

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(result)
}

func main() {
  err := models.InitDB()
  if (err != nil) {
    fmt.Printf("Error establishing connection to database: %v\n", err)
    return
  }

  http.HandleFunc("/graphql", graphqlHandler)
  http.ListenAndServe(":8080", nil)
}
