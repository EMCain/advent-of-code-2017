package main

import (
  "container/ring"
  "fmt"
)

func makeRing(size int) *ring.Ring {
  r := ring.New(size)
  for i := 0; i < size; i++ {
    r.Value = i
    r = r.Next()
  }
  return r
}

func hash(r *ring.Ring, input []int) *ring.Ring {
  skip := 0
  newRing := ring.New(r.Len())
  for i := 0; i < len(input); i++ {
    sublist_size := input[i]
    sublist := make([]int, sublist_size)
    for j := 0; j < sublist_size; j++ {
      sublist[j] = r.Value.(int)
      r = r.Next()
    }
    fmt.Printf("%v", sublist)
    // TODO reverse list; should I apply this to r rather than creating newRing?
    for j:= 0; j < skip; j++ {
      r = r.Next()
    }
    skip++
  }
  return newRing
}

func main() {
  smallR := makeRing(5)

  smallR.Do(func (x interface{}) {
    fmt.Println(x)
  })
  hash(smallR, []int {3, 4, 1, 5})
}
