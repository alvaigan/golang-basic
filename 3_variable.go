package main

import "fmt"

func main() {
  var name string

  name = "Alvaigan"
  fmt.Println("Hello", name)

  name = "Justin"
  fmt.Println("Hello", name)

  var friendName = "Baleo"
  fmt.Println("Hello", friendName)

  // jika datanya tidak diinisialisasi, maka bisa tanpa menggunakan "var"
  sekolah := "SD Tanjung Sari"
  sekolah = "Eh bukan di SD Tanjung Sari"
  fmt.Println("Kamu sekolah dimana? : ", sekolah)

  //multiple variable
  var (
    firstName = "Alvaigan"
    secondName = "Justin"
  )

  fmt.Println("Hello", firstName, secondName)

  const ahehe = 1000
  fmt.Println(ahehe)
}

