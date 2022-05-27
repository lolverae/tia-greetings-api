package main

import (
  "net/http"
  "encoding/json"
  "regexp"
  "sync"
)

// greetings JSON structure
type greet struct {
  ID       string `json:"id"`
  Content  string `json:"content"`
}

type morningHandler struct {}
var getGreeting   = regexp.MustCompile(`^\/users[\/]*$`)
var createGreting = regexp.MustCompile(`^\/users[\/]*$`)

func (h *morningHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("content-type", "application/json")
  switch {
    case r.Method == http.MethodGet && listUserRe.MatchString(r.URL.Path):
        h.Get(w, r)
      return
    case r.Method == http.MethodPost && listUserRe.MatchString(r.URL.Path):
        h.Create(w, r)
      return
    default:
      notFound(w, r)
      return
    } 
}

func main(){
  mux := http.NewServceMux()
  mux.Handle("/greet/", &morningHandler{})
  http.ListenAndServe(":8080", mux)
}


