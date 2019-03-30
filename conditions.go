package main
import "fmt"

func switchTest(x int){
  switch x {
  case 1:
    fmt.Println("Passed 1")
  case 2:
    fmt.Println("Passed 2")
  case 3:
    fmt.Println("Passed 3")
  default:
    fmt.Println("default condition!")
  }
}

func main(){
  var x int = 100
  if x < 100 {
    fmt.Println("x is lesser than 100")
  }else if x > 100 {
    fmt.Println("x is greater than 100")
  }else {
    fmt.Println("x is equal to 100")
  }
  switchTest(3)
}
