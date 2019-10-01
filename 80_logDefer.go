// using defer can be useful in logging because it can allow you to 
// place logs from one function in closing/opening statement allowing 
// you to clearly trace function execution in log files
package main

import (
  "fmt"
  "log"
  "os"
)

var LOGFILE = "/var/log/go_deffered.log"

func one(aLog *log.Logger) {
  aLog.Println("-- Function one start --")
  defer aLog.Println("-- Function one end --")

  for i := 0; i < 10; i++ {
    aLog.Println(i)
  }
}

func two(aLog *log.Logger) {
  aLog.Println("-- Function two --")
  defer aLog.Println("-- Function two end --")

  for i := 10; i > 0; i-- {
    aLog.Println(i)
  }
}

func main() {
  f, err := os.OpenFile(LOGFILE, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println(err)
    return
 }
 defer f.Close()

 iLog := log.New(f, "logDefer", log.LstdFlags)
 iLog.Println("Hello there")
 iLog.Println("Another entry")
 one(iLog)
 two(iLog)

}

/* output
root@box:~/devstuff/gostuff/gostuff# cat /var/log/go_deffered.log
logDefer2019/09/29 16:31:13 Hello there
logDefer2019/09/29 16:31:13 Another entry
logDefer2019/09/29 16:31:13 -- Function one start --
logDefer2019/09/29 16:31:13 0
logDefer2019/09/29 16:31:13 1
logDefer2019/09/29 16:31:13 2
logDefer2019/09/29 16:31:13 3
logDefer2019/09/29 16:31:13 4
logDefer2019/09/29 16:31:13 5
logDefer2019/09/29 16:31:13 6
logDefer2019/09/29 16:31:13 7
logDefer2019/09/29 16:31:13 8
logDefer2019/09/29 16:31:13 9
logDefer2019/09/29 16:31:13 -- Function one end --
logDefer2019/09/29 16:31:13 -- Function two --
logDefer2019/09/29 16:31:13 10
logDefer2019/09/29 16:31:13 9
logDefer2019/09/29 16:31:13 8
logDefer2019/09/29 16:31:13 7
logDefer2019/09/29 16:31:13 6
logDefer2019/09/29 16:31:13 5
logDefer2019/09/29 16:31:13 4
logDefer2019/09/29 16:31:13 3
logDefer2019/09/29 16:31:13 2
logDefer2019/09/29 16:31:13 1
logDefer2019/09/29 16:31:13 -- Function two end --

*/
