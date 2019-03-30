package main

import "fmt"

func main(){
  const array_size int = 25
  var a [array_size]int
  for i := 0; i < array_size; i++ {
    a[i] = i * 10
  }
  fmt.Println("Printing array a")
  for i:= 0; i < array_size; i++ {
    fmt.Println(a[i])
  }
}
