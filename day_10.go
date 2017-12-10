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

func main() {
  smallR := makeRing(5)

  smallR.Do(func (x interface{}) {
    fmt.Println(x)
  })
}
