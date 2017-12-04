package main

import (
  "fmt"
  "math"
)

type Square struct {
  X, Y, Value int // the X and Y coordinates, with 1 at the origin (0, 0), and square's value
}

func (s Square) ManhattanDistance() int {
    return int(
      math.Abs(float64(s.X)) +
      math.Abs(float64(s.Y)))
}

func getNextDirection (old [2] int) [2] int {
  // there's probably a better way to do this, oh well
  switch old {
  case [2]int{1, 0}:
      return [2]int{0, 1}
    case [2]int{0, 1}:
      return [2]int{-1, 0}
    case [2]int{-1, 0}:
      return [2]int{0, -1}
    case [2]int{0, -1}:
      return [2]int{1, 0}
  }
  return  [2]int{0, 0}
}

/*
  directions:
    +x 1
    +y 1
    -x 2
    -y 2
    +x 3
    +y 3
    etc.

  values:
     1 | 0,  0
     2 | 1,  0
     3 | 1,  1
     4 | 0,  1
     5 | -1, 1
     6 | -1, 0
     7 | -1,-1

*/

func getSquares(num int) []Square {
  grid := make([]Square, num)
  direction := [2]int{0, -1}
  // the change in coordinates + or - 1 or 0
  var x, y int
  // x and y are the locations of the squares
  distance := 1
  // distance is how long it will keep going that direction
  distance_counter := 0
  // how long it has been going in that direction so far.


  for i, _ := range grid {
    // if value - 1 (or i) is a square number, distance will be increased and direction will change
    root := math.Sqrt(float64(i))
    if root == float64(int(root)) {
      // if that square is odd, direction will be ->, and if it's even direction will be <-
      // also the distance will go up  (or is this in the other case? )
      direction = getNextDirection(direction)
      distance_counter = 0
    } else if distance == distance_counter {
      distance++
      direction = getNextDirection(direction)
      distance_counter = 0
      // change direction

    }
    grid[i] = Square{x, y, i+1}

    // change x and y appropriately
    x += direction[0]
    y += direction[1]
    distance_counter++

  }

  return grid
}

func ManhattanDistanceForNumber(i int) int {
  grid := getSquares(i)
  return grid[i-1].ManhattanDistance()
}

func main() {
  fmt.Println("hello")
  fmt.Println(getSquares(10))

  fmt.Println(ManhattanDistanceForNumber(1))
  fmt.Println(ManhattanDistanceForNumber(12))
  fmt.Println(ManhattanDistanceForNumber(23))
  fmt.Println(ManhattanDistanceForNumber(1024))
  fmt.Println(ManhattanDistanceForNumber(368078))
}
