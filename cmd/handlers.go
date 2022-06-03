package cmd

import (
  "fmt"
  "net/http"
  "encoding/json"
)

func index(w http.ResponseWriter, resp *http.Request) {

  if resp.Method != http.MethodGet {
    fmt.Fprintf(w, "Whoops, this method is not allowed")
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
  }
  fmt.Fprintf(w, "Whoops, this method is not allowed")

}

func getAllGreets(w http.ResponseWriter, resp *http.Request) {
  w.Header().Set("Content-Type", "application/json") 
  json.NewEncoder(w).Encode(greetings)
  fmt.Fprintf(w, "Greet")
}

func addGreet(w http.ResponseWriter, resp *http.Request) {

}
