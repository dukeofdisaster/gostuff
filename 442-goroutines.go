// highlighting go routines for understanding concurrency / multithreading / channels
package main

import (
  "fmt"
  "time"
)

func function() {
  for i:= 0; i < 10 ; i++ {
    fmt.Println(i)
  }
}

func main() {
  go function()

  go func() {
    for i := 10; i < 20; i++ {
      fmt.Print(i, " ")
    }
  }()

  time.Sleep(1*time.Second)
  fmt.Println()
}
