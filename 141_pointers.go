// illustrates some pointer usage * == value,& == memory address
package main
import "fmt"

// allows you to update the variable passed without the need to return anything
func getPointer(n *int) {
  *n = *n * *n
}

func returnPointer(n int) *int {
  v := n * n
  return &v
}

func main() {
  i := -10
  j := 25
  fmt.Println("i start:", i)
  fmt.Println("j start:", j)

  // initializing to pointers by referencing
  // the memory address of the objects we want to point 
  // to
  pI := &i
  pJ := &j

  fmt.Println("pI memory address:", pI)
  fmt.Println("pJ memory address:", pJ)
  fmt.Println("pI value:", *pI)
  fmt.Println("pJ value:", *pJ)

  // decrement pI by one
  *pI = 123456
  fmt.Println("*pI is the dereferenced pointer, basically same as accessing i: ")
  *pI--
  fmt.Println("i:", i)

  getPointer(pJ)
  fmt.Println("j:",j)
  k := returnPointer(12)
  fmt.Println(*k)
  fmt.Println(k)
}

