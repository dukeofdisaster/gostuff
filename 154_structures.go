// highlights use of structres and strucure literals
package main
import "fmt"

func main () {
  type XYZ struct {
    X int
    Y int
    Z int
  }

  var s1 XYZ
  fmt.Println(s1.Y, s1.Z)

  // one way of defining a structure literal
  p1 := XYZ{23, 12, -2}

  //another way in case you can't remmber the order
  p2 := XYZ{Z:12, Y: 13}
  fmt.Println(p1)
  fmt.Println(p2)
  
  pSlice := [4]XYZ{}
  pSlice[2] = p1
  pSlice[0] = p2
  fmt.Println(pSlice)

  p2 = XYZ {1, 2, 3}
  fmt.Println(pSlice)
  fmt.Println("Looping over pSlice\n======")

  // common for loop over a slice
  for i := range pSlice {
    fmt.Println(pSlice[i])
    fmt.Println(pSlice[i].Z)
  }

}
/* output
0 0
{23 12 -2}
{0 13 12}
[{0 13 12} {0 0 0} {23 12 -2} {0 0 0}]
[{0 13 12} {0 0 0} {23 12 -2} {0 0 0}]
Looping over pSlice
======
{0 13 12}
12
{0 0 0}
0
{23 12 -2}
-2
{0 0 0}
0
*/
