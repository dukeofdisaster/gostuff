// go can also be used to generate web-assembly code (wasm) that can run on 
// browsers. Wasm benefits: near native OS speed, faster than JS
// the code will be the same but the building will be different
// root@box$ GOOS=js GOARCH=wasm go build -o main.wasm 101_gowasm.go
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Created web assembly code in go!")
}
