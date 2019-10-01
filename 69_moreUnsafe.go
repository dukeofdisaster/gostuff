// highlights unsafe package usage and usage of memory addresses
package main

import (
  "fmt"
  "unsafe"
)

func main () {
  array := [...]int{0,1,-2,-3,4}
  // Here, the variable "pointer" points to the memory address of array[0], 
  pointer := &array[0]
  fmt.Print(*pointer, " ")
  memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])

  for i:=0; i< len(array)-1; i++ {
    // pointer variable is converted to an unsafe.Pointer()
    pointer = (*int)(unsafe.Pointer(memoryAddress))

    // common dereference notation: *pointer allows us to get the value stored at the memory address the pointer points to
    fmt.Print(*pointer, " ")
    // pointer is converted to uintptr and stored in memory address
    // note that the value of unsafe.Sizeof(array[0]) is what allows us to move to the next memory address
    // and thus print the the value (integer) stored at that address
    memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
  }
  fmt.Println()
  // here we try to access an element of the array that does not exist
  pointer = (*int)(unsafe.Pointer(memoryAddress))
  fmt.Print("One more: ", *pointer, " ")
  memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
  fmt.Println()
}
/* output
root@box:~/devstuff/gostuff/gostuff# go run 69_moreUnsafe.go
0 1 -2 -3 4
One more: 824634322696
*/
