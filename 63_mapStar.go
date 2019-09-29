// highlights using maps for storing pointers as integers.to be used with the 63/64 .go files
// for highlighting performance of the various methods
package main
import (
  "runtime"
)

func main() {
  var N = 40000000
  myMap := make(map[int]*int)
  for i := 0; i < N; i++ {
    value := int(i)
    myMap[value] = &value
  }
  runtime.GC()
  _ = myMap[0]
}
