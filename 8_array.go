package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Alvaigan"
	names[1] = "Kasep"
	names[2] = "Ganteng"
  
  fmt.Println(names[0])
  fmt.Println(names[1])
  fmt.Println(names[2])

  var numbers = [5] int {
    1,
    2,
    3,
    4,
    5,
  }

  fmt.Println(numbers)
  fmt.Println(len(numbers))
  numbers[0] = 10
  fmt.Println(numbers)
}
