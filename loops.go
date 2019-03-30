package main

import "fmt"

func singleCondition(){
  var i = 0
  for i < 3 {
    fmt.Println(i)
    i = i + 1
  }
}

func infiniteLoop(){
  for{
    fmt.Println("infinite loop!")
  }
}

func normalLoop(){
  for i := 0; i <= 10; i++{
    fmt.Println(i)
  }
}

func main(){
  singleCondition()
  fmt.Println("Normal loop:")
  normalLoop()
  fmt.Println("Printing infinite loop: ")
  //infiniteLoop()
}
