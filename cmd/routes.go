package cmd

import (
  "net/http"
)


func initRoutes() {

  http.HandleFunc("/", index)
  http.HandleFunc("/greetings", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getAllGreets(w,r)
    case http.MethodPost:
        addGreet(w,r)
    default:
      w.WriteHeader(http.StatusMethodNotAllowed)
      return
    } 
  })
}
