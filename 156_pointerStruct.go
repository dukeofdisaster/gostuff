// Highlights the use of functions that return pointers to structures
// - added p.158 use of the new keyword
package main
import "fmt"

type myStructure struct {
  Name string
  Last string
  Height int32
}

// pointer 
func createStruct(n, s string, h int32) *myStructure {
  if h > 300 {
    h = 0
  }
  return &myStructure{n, s, h}
}

// no pointer more appropro: NewStructure and NewStructurePointer
func retStructure(n, s string, h int32) myStructure {
  if h > 300 {
    h = 0
  }
  return myStructure{n,s,h}
}

func main() {
  s1 := createStruct("dukePtr", "ofdisasterPtr", 121)
  s2 := retStructure("duke", "ofdisaster", 123)

  // here the enclosing parenthesis around *s1 is required because otherwise
  // we would be trying to dereference a pointer to the Name variable, not the struct
  // giving us 
  /*
  user@box:~$ go run 156_pointerStruct.go
  # command-line-arguments
  ./156_pointerStruct.go:31:15: invalid indirect of s1.Name (type string)
  */
  fmt.Println((*s1).Name)
  fmt.Println(s2.Name)
  fmt.Println(s1)
  fmt.Println(s2)

  // new keyword usage. Important: new returns a pointer to the object
  pointer_new  := new(myStructure)
  // should all be zero / empty string
  fmt.Println(*pointer_new)
  fmt.Println(pointer_new)

  // this should be true since the struct hasn't been initialized
  if (*pointer_new).Name == "" {
    fmt.Println("As excpected (*pointer_new).Name is empty")
  }
  // default empty value of Height will be zero
  if (*pointer_new).Height == 0 {
    fmt.Println("As expected (*pointer_new).Height is 0")
  }

}

/* output
dukePointer
duke
&{dukePointer ofdisasterPointer 121}
{duke ofdisaster 123}
*/
