// example of writing to Stderr
// this will let you succesfully direct stderr using fd 2
// go run std_err.go 2>/dev/null
package main

import (
  "io"
  "os"
)
func main() {
  myString := ""
  arguments := os.Args
  if len(arguments) == 1 {
    myString = "Please give me one argument!"
  } else {
    myString = arguments[1] 
  }
  io.WriteString(os.Stdout, "This is stdout\n")
  io.WriteString(os.Stderr, myString)
  io.WriteString(os.Stderr, "\n")
}

