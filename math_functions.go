package main
import(
  "fmt"
  "math"
  "math/rand"
)
func main(){
  //godoc <command> to get documentation
  fmt.Println("Square root of 17: ", math.Sqrt(17))
  fmt.Println("Random number from 1 to 100: ", rand.Intn(100))
  fmt.Printf("2 power 10 is: %f\n", math.Pow(2,10))
}
