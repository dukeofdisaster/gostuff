/* 
go doesn't have the while keyword() instead you can use a for loop with variable values.
looping over ranges will also becovered
*/
package main
import ( 
  "fmt"
)

func main() {

  // typical for loop
  for i := 0; i < 100 ; i++ {
    if i % 20 == 0 {
      continue
    }
    if i == 95 {
      break
    }

    fmt.Print(i, " ")
  }
  fmt.Println()

  // atypical
  i := 10
  for {
    if i < 0 {
      break
    }
    fmt.Print(i, " ")
    i--
  }
  fmt.Println()

  // loop with boolean
  i = 0
  anExpression := true
  for ok := true; ok; ok = anExpression {
    if i > 10 {
      anExpression = false
    }

    fmt.Print(i, " ")
    i++
  }
  fmt.Println()

  // range example
  anArray := [5]int{0, 1, -1, 2, -2}
  for i, value := range anArray {
    // note one small thing about go in these statements; adding a space after
    // the "index:" and "value:" strings is unnecessary because it's automatically 
    // inserted. 
    fmt.Println("index:", i, "value:", value)
  }
}

/* output
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 81 82 83 84 85 86 87 88 89 90 91 92 93 94
10 9 8 7 6 5 4 3 2 1 0
0 1 2 3 4 5 6 7 8 9 10 11
index: 0 value: 0
index: 1 value: 1
index: 2 value: -1
index: 3 value: 2
index: 4 value: -2

*/
