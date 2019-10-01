// constant generator iota is used for declaring a sequence of related values 
// that use incrementing numbers without the need to explictly type each of them.
//
package main
import "fmt"

type Digit int
type Power2 int

const PI = 3.1415926

const (
  C1 = "c1c1c1"
  C2 = "c2c2c2"
  C3 = "c3c3c3"
)

func main() {
  // constant defined and used in an expression
  const s1 = 123
  var v1 float32 = s1 * 12
  fmt.Println(v1)
  fmt.Println(PI)

  // definition of a constant generator iota based on Digit (defined on line 7 above)
  // which is equivalent  to Zero =0, One = 1, Two = 2 ...
  const (
    Zero Digit = iota
    One
    Two
    Three
    Four
  )
  fmt.Println(One)
  fmt.Println(Two)
  const (
    // here iota has the value of 0
    p2_0 Power2 = 1 << iota
    // here iota has the value of 1
    _
    // here iota has the value of 2
    p2_2
    // here iota has the value of 3... etc
    _
    p2_4
    _
    p2_6
  )

  fmt.Println("2^0:", p2_0)
  fmt.Println("2^2:", p2_2)
  fmt.Println("2^4:", p2_4)
  fmt.Println("2^6:", p2_6)
}


