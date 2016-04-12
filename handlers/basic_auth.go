package handlers

import (
  "net/http"

  "github.com/dhanui/graphql_api_go/repository"
)

func basicAuth(r *http.Request) (user repository.User, ok bool) {
  username, password, ok := r.BasicAuth()
  if ok {
    user, ok = repository.AuthenticateUser(username, password)
  }
  return
}
