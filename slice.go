package main

import "fmt"

func main(){
  numSlice := []int {1,2,3,4,5,6,7,8}
  numSlice2 := numSlice[0:3]
  fmt.Println(numSlice2)

  // make is used when you don't want defined set of values while starting
  // format: <name_of_slice> := make([]<data_type>, <initial_size_with_zeroes>, max_size)

  newSlice3 := make([]int, 10, 10)
  fmt.Println("newSlice: ", newSlice3)

  //copying from newSlice to newSlice3
  copy(newSlice3, numSlice)
  fmt.Println("newSlice3: ", newSlice3)
}
