package main
import (
  "github.com/lolverae/tiaGreetingsApi/cmd"
)

func main() {

  srv := cmd.New(":9090")

  err := srv.ListenAndServe()

  if err != nil {
    panic(err)
  }
}
