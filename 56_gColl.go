// highlights garbage collection
package main

import (
  "fmt"
  "runtime"
  "time"
)

func printStats(mem runtime.MemStats) {
  runtime.ReadMemStats(&mem)
  fmt.Println("mem.Alloc: ", mem.Alloc)
  fmt.Println("mem.TotalAlloc: ", mem.TotalAlloc)
  fmt.Println("mem.HeapAlloc: ", mem.HeapAlloc)
  fmt.Println("mem.NumGC: ", mem.NumGC)
  fmt.Println("======================")
}

func main() {
  var mem runtime.MemStats
  printStats(mem)

  for i := 0; i < 10; i++{
    s := make([]byte, 500000000)
    if s == nil {
      fmt.Println("Operation failed...")
    }
  }
  printStats(mem)

  for i := 0; i < 10; i++{
    s := make([]byte, 100000000)
    if s == nil {
      fmt.Println("Operation failed")
    }
    time.Sleep(1*time.Second)
    printStats(mem)
  }
}

/* Sample output

mem.Alloc:  141704
mem.TotalAlloc:  141704
mem.HeapAlloc:  141704
mem.NumGC:  0
======================
mem.Alloc:  146368
mem.TotalAlloc:  5000217696
mem.HeapAlloc:  146368
mem.NumGC:  10
======================
mem.Alloc:  147160
mem.TotalAlloc:  5100228224
mem.HeapAlloc:  147160
mem.NumGC:  11
======================
mem.Alloc:  147424
mem.TotalAlloc:  5200238368
mem.HeapAlloc:  147424
mem.NumGC:  12
======================
mem.Alloc:  147424
mem.TotalAlloc:  5300248120
mem.HeapAlloc:  147424
mem.NumGC:  13
======================
mem.Alloc:  147424
mem.TotalAlloc:  5400257872
mem.HeapAlloc:  147424
mem.NumGC:  14
======================
mem.Alloc:  147424
mem.TotalAlloc:  5500267624
mem.HeapAlloc:  147424
mem.NumGC:  15
======================
mem.Alloc:  147424
mem.TotalAlloc:  5600277376
mem.HeapAlloc:  147424
mem.NumGC:  16
======================
mem.Alloc:  147424
mem.TotalAlloc:  5700287128
mem.HeapAlloc:  147424
mem.NumGC:  17
======================
mem.Alloc:  147424
mem.TotalAlloc:  5800296880
mem.HeapAlloc:  147424
mem.NumGC:  18
======================
mem.Alloc:  147424
mem.TotalAlloc:  5900306632
mem.HeapAlloc:  147424
mem.NumGC:  19
======================
mem.Alloc:  147424
mem.TotalAlloc:  6000316384
mem.HeapAlloc:  147424
mem.NumGC:  20
======================

*/

/* Output with GODEBUG=gctrace=1 go run 56_gColl.go

** USER COMMENT
Consider 4->4->0 MB ; 

First number from left is heap size before garbage collector (GC) runs.
Second number is the heap size when collector ends operation. 
Third number is size of the live heap
** END UESR COMMENT

gc 1 @0.017s 0%: 0.009+0.32+0.019 ms clock, 0.11+0.092/0.51/0.16+0.23 ms cpu, 4->4->0 MB, 5 MB goal, 12 P
gc 2 @0.024s 0%: 0.005+0.32+0.003 ms clock, 0.067+0.12/0.37/0.28+0.047 ms cpu, 4->4->0 MB, 5 MB goal, 12 P
gc 3 @0.032s 1%: 0.11+1.4+0.20 ms clock, 1.4+0.45/0.96/0+2.4 ms cpu, 4->5->1 MB, 5 MB goal, 12 P
gc 4 @0.039s 1%: 0.019+0.36+0.005 ms clock, 0.23+0.32/0.68/0.65+0.066 ms cpu, 4->4->1 MB, 5 MB goal, 12 P
gc 5 @0.040s 1%: 0.011+0.34+0.014 ms clock, 0.13+0.12/0.46/0.75+0.17 ms cpu, 4->4->1 MB, 5 MB goal, 12 P
gc 6 @0.045s 1%: 0.003+0.24+0.003 ms clock, 0.047+0/0.38/1.0+0.042 ms cpu, 4->4->1 MB, 5 MB goal, 12 P
gc 7 @0.047s 1%: 0.002+0.25+0.009 ms clock, 0.033+0.096/0.37/0.64+0.10 ms cpu, 4->4->1 MB, 5 MB goal, 12 P
# command-line-arguments
gc 1 @0.002s 4%: 0.042+0.98+0.007 ms clock, 0.51+0.044/1.6/0.67+0.086 ms cpu, 5->5->4 MB, 6 MB goal, 12 P
# command-line-arguments
gc 1 @0.000s 7%: 0.029+0.63+0.007 ms clock, 0.35+0.075/0.97/0.55+0.086 ms cpu, 5->6->6 MB, 6 MB goal, 12 P
gc 2 @0.006s 3%: 0.005+2.5+0.004 ms clock, 0.065+0.053/2.3/0.25+0.056 ms cpu, 8->9->9 MB, 12 MB goal, 12 P
gc 3 @0.010s 5%: 0.002+2.4+0.007 ms clock, 0.035+0.15/4.0/1.8+0.093 ms cpu, 15->16->15 MB, 18 MB goal, 12 P
gc 4 @0.034s 3%: 0.003+6.4+0.005 ms clock, 0.047+0.18/9.6/0.19+0.068 ms cpu, 27->27->25 MB, 31 MB goal, 12 P
mem.Alloc:  141896
mem.TotalAlloc:  141896
mem.HeapAlloc:  141896
mem.NumGC:  0
======================
gc 1 @0.005s 0%: 0.003+0.083+0.003 ms clock, 0.036+0.057/0.003/0.038+0.047 ms cpu, 477->477->0 MB, 478 MB goal, 12 P
gc 2 @0.187s 0%: 0.003+0.085+0.003 ms clock, 0.041+0.062/0.012/0.008+0.039 ms cpu, 476->476->0 MB, 477 MB goal, 12 P
gc 3 @0.202s 0%: 0.002+0.084+0.002 ms clock, 0.034+0.063/0.005/0.015+0.033 ms cpu, 476->476->0 MB, 477 MB goal, 12 P
gc 4 @0.220s 0%: 0.002+0.085+0.002 ms clock, 0.032+0.059/0.005/0.014+0.034 ms cpu, 476->476->0 MB, 477 MB goal, 12 P
gc 5 @0.238s 0%: 0.002+0.13+0.003 ms clock, 0.032+0.076/0.029/0.008+0.043 ms cpu, 476->476->0 MB, 477 MB goal, 12 P
gc 6 @0.254s 0%: 0.003+0.085+0.003 ms clock, 0.040+0.060/0.006/0.013+0.042 ms cpu, 476->476->0 MB, 477 MB goal, 12 P
gc 7 @0.272s 0%: 0.003+0.13+0.003 ms clock, 0.036+0.079/0.009/0.030+0.036 ms cpu, 476->476->0 MB, 477 MB goal, 12 P
gc 8 @0.289s 0%: 0.003+0.085+0.003 ms clock, 0.037+0.077/0/0.038+0.046 ms cpu, 476->476->0 MB, 477 MB goal, 12 P
gc 9 @0.308s 0%: 0.003+0.13+0.003 ms clock, 0.043+0.086/0.016/0.033+0.044 ms cpu, 476->477->0 MB, 477 MB goal, 12 P
gc 10 @0.325s 0%: 0.003+0.082+0.003 ms clock, 0.041+0/0.075/0.005+0.037 ms cpu, 476->476->0 MB, 477 MB goal, 12 P
mem.Alloc:  148480
mem.TotalAlloc:  5000219840
mem.HeapAlloc:  148480
mem.NumGC:  10
======================
gc 11 @0.328s 0%: 0.001+0.090+0.002 ms clock, 0.018+0.030/0.015/0.046+0.028 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
mem.Alloc:  149272
mem.TotalAlloc:  5100230368
mem.HeapAlloc:  149272
mem.NumGC:  11
======================
gc 12 @1.335s 0%: 0.002+0.13+0.002 ms clock, 0.026+0/0.11/0+0.034 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
mem.Alloc:  149536
mem.TotalAlloc:  5200240512
mem.HeapAlloc:  149536
mem.NumGC:  12
======================
gc 13 @2.345s 0%: 0.007+0.28+0.012 ms clock, 0.095+0/0.23/0.073+0.14 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
mem.Alloc:  149536
mem.TotalAlloc:  5300250264
mem.HeapAlloc:  149536
mem.NumGC:  13
======================
gc 14 @3.352s 0%: 0.002+0.13+0.002 ms clock, 0.026+0/0.073/0+0.034 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
mem.Alloc:  149536
mem.TotalAlloc:  5400260016
mem.HeapAlloc:  149536
mem.NumGC:  14
======================
gc 15 @4.359s 0%: 0.002+0.20+0.002 ms clock, 0.027+0/0.093/0+0.029 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
mem.Alloc:  149992
mem.TotalAlloc:  5500270224
mem.HeapAlloc:  149992
mem.NumGC:  15
======================
gc 16 @5.366s 0%: 0.002+0.13+0.002 ms clock, 0.024+0/0.073/0+0.029 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
mem.Alloc:  149992
mem.TotalAlloc:  5600279976
mem.HeapAlloc:  149992
mem.NumGC:  16
======================
gc 17 @6.373s 0%: 0.002+0.15+0.002 ms clock, 0.026+0/0.094/0+0.030 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
mem.Alloc:  149992
mem.TotalAlloc:  5700289728
mem.HeapAlloc:  149992
mem.NumGC:  17
======================
gc 18 @7.380s 0%: 0.002+0.15+0.002 ms clock, 0.026+0/0.070/0+0.030 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
mem.Alloc:  149992
mem.TotalAlloc:  5800299480
mem.HeapAlloc:  149992
mem.NumGC:  18
======================
gc 19 @8.387s 0%: 0.002+0.16+0.002 ms clock, 0.024+0/0.10/0+0.035 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
mem.Alloc:  149992
mem.TotalAlloc:  5900309232
mem.HeapAlloc:  149992
mem.NumGC:  19
======================
gc 20 @9.394s 0%: 0.002+0.15+0.002 ms clock, 0.024+0/0.094/0+0.027 ms cpu, 95->95->0 MB, 96 MB goal, 12 P
mem.Alloc:  149992
mem.TotalAlloc:  6000318984
mem.HeapAlloc:  149992
mem.NumGC:  20
======================

*/
