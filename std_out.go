package main

import (
  "io"
  "os"
)

// init function will run 1 time when the package is imported / ran
func init() {
  io.WriteString(os.Stdout, "Running program...\n")
}

func main() {
  //green := "\033[92m"
  end := "\033[0m"
  red := "\033[91m"
  myString := ""
  arguments := os.Args
  if len(arguments) == 1 {
    myString = "Please give me one argument!"
  } else {
      myString = arguments[1]
  }
  io.WriteString(os.Stdout, red)
  io.WriteString(os.Stdout, myString)
  io.WriteString(os.Stdout, "\n")
  io.WriteString(os.Stdout, end)
}
