package main

import "fmt"

func main(){
  //declare & initialize an array
  //<name> := [<size>] <data_type> {values1, value2}
  myArray := [5] int {10,20,30,40,50}

  for i, value := range myArray {
    fmt.Println(i, value)
  }

}
