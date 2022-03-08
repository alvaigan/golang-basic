package main

import "fmt"

func main() {
  // tipe data number
  fmt.Println("Satu = ", 1)
  fmt.Println("Dua = ", 2)
  fmt.Println("Tiga koma lima = ", 3.5)

  // tipe data boolean
  fmt.Println("Benar = ", true)
  fmt.Println("Salah = ", false)

  // tipe data string
  fmt.Println("Hello Alvaigan")
  fmt.Println(len("Hello Alvaigan"))
  fmt.Println("Hello Alvaigan"[0]) // nilai yang ditampilkan adalah index ke 0 tapi representasi byte nya
}
