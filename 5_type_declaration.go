package main

import "fmt"

func main() {
  type NoKTP string
  type Married bool

  var ktpSaya NoKTP = "123841239749"
  var aku Married = true

  fmt.Println(ktpSaya)
  fmt.Println(aku)
}
