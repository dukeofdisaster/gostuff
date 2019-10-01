// arrays are immutable and you'll probably avoid using them. instead, slices are what
// you'll want more often than not since they are dynamic
package main

import "fmt"

func main() {
  s1 := make([]int, 5)
  // note here that the rightmost of the sliced index, i.e :3 is exclusive,
  // thus the reslice will include two objects: s1[1], s1[2]
  reSlice := s1[1:3]
  fmt.Println(s1)
  fmt.Println(reSlice)

  reSlice[0] = -100
  reSlice[1] = 123456
  fmt.Println(s1)
  fmt.Println(reSlice)

}
/* output
[0 0 0 0 0]
[0 0]
[0 -100 123456 0 0]
[-100 123456]

*/
