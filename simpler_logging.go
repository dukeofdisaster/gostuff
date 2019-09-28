// A different logging example than the book.. here we KISS
// since this will be a quicker way of logging anything we want
// - Pulled color ideas were from python projects where i used the same
package main

import (
  "log"
  "os"
)

func main() {
  // without os.O_RDWR no log will be written
  file, err := os.OpenFile("/var/log/go_logging.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
  if err != nil {
    log.Fatal(err)
  }
  // colors
  green := "\033[92m"
  end := "\033[0m"
  red := "\033[91m"

  // log gevels
  er := red+"ERROR\t"
  inf := green+"INFO\t"

  defer file.Close()
  log.SetOutput(file)
  log.Println(er+"Running from main in go_log.go"+end)
  log.Println(inf+"Running from main in go_log.go"+end)
}
