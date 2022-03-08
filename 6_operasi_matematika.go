package main

import "fmt"

func main() {
  var (
    a = 10
    b = 5
  )
  var c int

  c = a + b
  fmt.Println(c)


  c = a * b
  fmt.Println(c)

  
  a -= b
  fmt.Println(a)
}
