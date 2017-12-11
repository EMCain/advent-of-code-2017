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

// part 2

func convertASCII(s string) []int {
  result := make([]int, len(s))
  // iterate over each character in the string
  // add it to the int slice
  for i := 0; i < len(s); i++ {
    r := rune(s[i])
    result[i] = int(r)
  }
  return append(result, []int{17, 31, 73, 47, 23}...)
}

func xorIntegers(input []int) int {
  result := input[0]
  for i := 1; i < len(input); i++ {
    result = result ^ input[i]
  }
  return result
}

// eventually I'll be passing a Ring into this.
func hexString(input []int) string {
  var s string
  for i := 0; i < len(input); i++ {
    s = fmt.Sprintf("%s%02x", s, input[i])
  }
  return s
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
  fmt.Printf("\nconverted: %v \n", convertASCII("1,2,3"))
  fmt.Printf("\n%v\n\n", xorIntegers([]int {65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}))

  // test the hex function
  fmt.Println("\n--hex test---")

  fmt.Println(hexString([]int{64, 7, 255}))
}
