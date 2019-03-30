package main

import "fmt"

func main(){
  x := 15
  a := &x
  fmt.Println("address of x: ", a)
  *a = 123
  fmt.Println("x after change: ", x)
}
