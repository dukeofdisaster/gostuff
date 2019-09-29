// Highlights why we should be cautious with the usage of the garbage collector
// the way you store pointers has a great impact on performance of the collector, 
// esp. when dealing with large amounts of pointers; there's no visible / time difference here.
// The point is this technique can be reused and we'll compare it with other methods
package main

import (
  "runtime"
)

type mydata struct {
  i,j int
}

func main() {
  var N = 40000000
  // this is a slice of type mydata
  var structure []mydata
  for i := 0; i < N; i++ {
    value := int(i)
    structure = append(structure, mydata{value, value})
  }
  runtime.GC()
  // Prevents GC from garbage collecting the structure variable too early, as
  // it is not referenced or used outside of the for loop.
  _ = structure[0]
}
