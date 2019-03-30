package main
import (
  "fmt"
  "reflect"
)
const PI = 3.14;

func adder(x , y float64) float64{
  return x+y
}

func multiple(a,b string)(string, string){
  //function with multiple return values
  return a,b
}

func main(){
  var num1, num2 float64 = 10,21
  fmt.Println(adder(num1,num2))
  fmt.Println("PI: ", PI)
  fmt.Println(multiple("Bruce", "Wayne"))

  //typecasting example
  var a int = 1234;
  var b float64 = float64(a)
  fmt.Println("b: ", b)
  fmt.Println("type of b: ", reflect.TypeOf(b))
}
