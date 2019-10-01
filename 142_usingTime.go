//highlights working with time; note RFC3339 == ISO 8601
package main
import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("Epoch time:", time.Now().Unix())
  t := time.Now()
  fmt.Println(t, t.Format(time.RFC3339))
  fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

  time.Sleep(time.Second)
  t1 := time.Now()
  fmt.Println("Time difference:", t1.Sub(t))

  formatT := t.Format("01 JANUARY 2006")
  fmt.Println(formatT)
  loc, _ := time.LoadLocation("Europe/Paris")
  londonTime := t.In(loc)
  fmt.Println("Paris:", londonTime)
}
/* output
Epoch time: 1569958896
2019-10-01 12:41:36.787572505 -0700 MST m=+0.000252886 2019-10-01T12:41:36-07:00
Tuesday 1 October 2019
Time difference: 1.000269055s
10 JANUARY 2019
Paris: 2019-10-01 21:41:36.787572505 +0200 CEST

*/
