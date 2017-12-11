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
    fmt.Printf("%d ", x)
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
      fmt.Printf("\nIndex is %d, about to replace %d with item %d", index, r.Value.(int), item)
      r.Value = item
      r = r.Next()
    }
    // advance "skip" times
    for j:= 0; j < skip; j++ {
      r = r.Next()
    }
    fmt.Printf("\nSublist is %v, skip is %d, ring is...", sublist, skip)
    r.Do(func(x interface{}) {
      fmt.Printf("%d ", x)
    })

    skip++
  }
  return r
}

func main() {
  smallR := makeRing(5)

  smallR.Do(func (x interface{}) {
    fmt.Println(x)
  })
  hash(smallR, []int {3, 4, 1, 5})
}
