package cmd
import (
  "net/http"
)

type Greet struct {

  ID    string
  Greet string
}

var greetings []*Greet = []*Greet{}

func New(addr string) *http.Server {
  initRoutes()

  return &http.Server{
    Addr: addr,
  }
} 
