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
  r.Do(func(x interface{}) {
  })
  for i := 0; i < len(input); i++ {
    // create a sublist
    sublist_size := input[i]
    rCopy := r // this is only used for reading the value
    sublist := make([]int, sublist_size)
    for j := 0; j < sublist_size; j++ {
      sublist[j] = rCopy.Value.(int)
      rCopy = rCopy.Next()
    }
    // reverse order the items of sublist and put them in ring
    for j := 0; j < sublist_size; j++ {
      index := sublist_size - j - 1
      item := sublist[index]
      r.Value = item
      r = r.Next()
    }
    // advance "skip" times
    for j:= 0; j < skip; j++ {
      r = r.Next()
    }

    skip++
  }
  return r
}

func main() {
  smallR := makeRing(5)
  bigR := makeRing(256)
  hash(smallR, []int {3, 4, 1, 5})
  hash(bigR, []int {130,126,1,11,140,2,255,207,18,254,246,164,29,104,0,224})
  smallR.Do(func(x interface{}) {
    fmt.Printf("%d ", x)
  })
  fmt.Println("\n-------------------")
  bigR.Do(func(x interface{}) {
    fmt.Printf("%d ", x)
  })
}
