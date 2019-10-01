// defer keyword allows you to put statements for example like closing a file,
// the function that opens the file, that way they are easier to track and the file can be 
// closed when the surrounding function exits.
// defer operates in LIFO order; meaning if you have the defer func1() defer func2() defer func3(),
// in main on ascending lines, when main is about to exit, func3() will be called first
/* output
1 2 3
0 0 0
1 2 3
*/
package main
import (
  "fmt"
)
// essentialy has 3 defer calls starting with defer(3) defer(2) defer(1)
// thus when d1() is called, we get the items in assending order, 1 2 3 since it's LIFO
func d1() {
  for i := 3; i > 0; i-- {
    defer fmt.Print(i, " ")
  }
}

// here we have something strange, 0 0 0... this is because when the for loop ends, i == 0
/// but the deferred anonymous function is evaluated after the for loop ends because it has no parameters, 
// meaning it is called 3 times with terminating value of i, or 0; 
func d2() {
  for i := 3; i > 0; i-- {
    defer func() {
      fmt.Print(i, " ")
    }()
  }
  fmt.Println()
}

// here we get 1 2 3 again.
// since the anonymous function takes a paramter, each time it is defferred() it gets and uses the 
// current value of i, as a result, each execution of the anonymous function has a different value to process
// and we see the same normal pattern as we did in d1()
// This method is preferred because you intentionally pass the desired variable in the anonymous function in an
// easy to understand way.
func d3() {
  for i := 3; i > 0; i-- {
    defer func (n int) {
      fmt.Print(n, " ")
    }(i) // call the immediately defined anonymous func above with the value i
  }
}


func main() {
  d1()
  d2()
  fmt.Println()
  d3()
  fmt.Println()
}
