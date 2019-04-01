package main

import "fmt"

type person struct{
  name string
  age int
}

func main(){
  var p1 = person{"sameer", 24}
  fmt.Println(p1.name)
}
