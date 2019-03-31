package main

import "fmt"

func main(){
  users := make(map[string] int)
  //key value pairs, key is string, value is int
  users["mark"] = 23
  users["john"] = 40
  users["steve"] = 32

  fmt.Println(users["steve"])
  for key, value := range users{
    fmt.Println(key, ": ", value)
    if key == "steve"{
      users[key] = -100
    }
  }
  fmt.Println("after looping: ", users)
  fmt.Println("size of map users: ", len(users))
  delete(users, "john")
  fmt.Println("After deleting: ", users)
}
