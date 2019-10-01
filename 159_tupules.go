// highlights use of tuples... generally when handling errors returned from
// functions, you're accessing a tuple.
package main
import "fmt"
func retThree(x int) (int, int , int) {
  return 2 * x, x * x, -x
}

func main() {
  // these will have essentially the same output, minus the values obvs
  fmt.Println(retThree(10))
  n1, n2, n3 :=  retThree(20)
  fmt.Println("n1:",n1, "n2:", n2, "n3:", n3)

  fmt.Println("Swapping tuple values with the following:  n1, n2 = n2, n1 ")

  // using plain variables to initialize in tuples
  n1, n2 = n2, n1
  fmt.Println(n1, n2, n3)

  x1, x2, x3 := n1*2, n1*n1, -n1
  fmt.Println(x1, x2, x3)

}
