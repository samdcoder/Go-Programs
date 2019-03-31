package main

import "fmt"

func sum(nums... int) int {
  var total int = 0
  for _ ,value := range nums {
    total += value
  }
  return total
}
func main(){
  //sum is a variadic function as it accepts different number of arguments 
  fmt.Println(sum(1,2,3))
  fmt.Println(sum(15,12))
}
