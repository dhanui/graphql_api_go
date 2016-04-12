package handlers

import (
  "net/http"
)

func cors(w *http.ResponseWriter) {
  ret := *w
  ret.Header().Set("Access-Control-Allow-Origin", "*")
  ret.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
  ret.Header().Set("Access-Control-Allow-Headers", "Authorization")
}
