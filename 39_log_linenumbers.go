package main

import (
  "fmt"
  "log"
  "os"
)

var LOGFILE = "/tmp/go_linenumbers.log"

// instead of implementing the same check over and over again
// we can write a small wrapper for it
func  dealwithit(some_error error) {
  if some_error != nil {
    fmt.Println(some_error)
 }
}

func main() {
  // Creates the file in append mode in stead of overwrite, creates file if it
  // does not exist, and creates it in Write / Read mode, with a umask of 0644
  f, err := os.OpenFile(LOGFILE, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
  dealwithit(err)

  defer f.Close()

  // creates logger from the file above, sets a prefix string to appear at start of each line, 
  // and sets the standard flags of the log file (date / time)
  thelog, err := log.New(f, "some text here", log.LstdFlags)
  dealwithit(err)

  // here we set the Lshortfile constant, which wll add the line number of the 
  // log message
  thelog.SetFlags(log.LstdFlags | log.Lshortfile)
  thelog.Println("Hello there!")
  thelog.Println("Another entry")
}
