## 1_hello_world.go
- file go bisa langsung dijalankan langsung dengan perintah: 
`go run 1_hello_world.go`
- file go bisa dicompile dulu (build dulu file yg nantinya bisa dicompile):
`go build 1_hello_world.go`
kemudian nantinya terdapat file 1_hello_world sebagai binary yang bisa dijalankan langsung

## 2_tipe_data.go
- tipe data number lebih detail:
1. int8
2. int16
3. int32
4. int64
5. uint8/16/32/64 => startnya dari 0
6. float32/64/complex64/complex128

- ada juga tipe data number yang alias
1. byte = uint8
2. rune = int32
3. int = (Minimal int32)
4. uint = (Minimal uint32)

- function untuk string:
1. len("string") => ngitung jumlah karakter
2. "string"[number] => ngambil posisi karakter

## 3_variable.go
- deklarasi bisa variable:
1. `var nama_variable string` (variable yang sudah dideklarasikan valuenya bisa diubah. tapi tipe datanya harus sama)
2. `var nama_variable = "data"` (tipe data variable akan mengikuti tipe data valuenya apa)
3. `nama_variable := "data"` (jika datanya tidak diinisialisasi, maka bisa juga tanpa menggunakan var)
- kalo ada variable yg tidak dipake bakal error
- tapi kalo ada const yang tidak dipake gak bakal error

## 4_konversi_tipe_data.go
- tinggal pake function int32(), int64(), string()

## 5_type_declaration.go
- `type` digunakan untuk melakukan alias suatu variable

## 6_operasi_matematika.co
- kurang lebih sama dengan bahasa pemrograman lain

## 7_operasi_perbandingan.go
- kurang lebih sama dengan bahasa pemrograman lain

## 8_array.go
- deklarasi array tanpa value:
`var names [3]string`

- deklarasi array dengan value:
`var numbers = [5] int {1,2,3,4,5}`

- `len()` untuk menghitung jumlah array
