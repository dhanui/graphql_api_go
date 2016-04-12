package handlers

import (
  "net/http"
  "encoding/json"

  "github.com/dhanui/graphql_api_go/schema"
)

func GraphqlHandler(w http.ResponseWriter, r *http.Request) {
  user, ok := basicAuth(r)
  if !ok {
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
  } else {
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
}
