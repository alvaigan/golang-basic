package main

import "fmt"

func main() {
  var nama1 = "Alvaigan"
  var nama2 = "Kasep"

  var result bool = nama1 == nama2
  fmt.Println("Is nama1 and nama2 same value?")
  fmt.Println(result)
  

  var value1 = 100
  var value2 = 200
  fmt.Println(value1 == value2)
  fmt.Println(value1 != value2)
  fmt.Println(value1 <= value2)
  fmt.Println(value1 >= value2)
  fmt.Println(value1 < value2)
  fmt.Println(value1 > value2)
}
