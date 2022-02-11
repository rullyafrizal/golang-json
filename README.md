# Belajar Golang JSON

## Encode JSON
- Go telah menyediakan function untuk mengkonversi data ke JSON, yaitu dengan menggunakan function `json.Marshal(interface{})`
- Parameter adalah interface kosong `interface{}`, sehingga bisa dimasukkan tipe data apapun

## JSON Object
- JSON Object di Go direpresentasikan dengan tipe data struct
- Tiap attribute di JSON Object merupakan atribute struct

## Decode JSON
- Untuk melakukan konversi dari JSON ke tipe data di Go (Decode), kita bisa menggunakan function `json.Unmarshal(byte[], interface{})`
- Dimana byte[] adalah data JSON-nya, sedangkan interface{} adalah tempat menyimpan hasil konversi, biasa berupa pointer

## JSON Array
- Di Go, JSON Array direpresentasikan dalam bentuk tipe data slice
- Konversi dari JSON atau ke JSON dilakukan otomatis oleh package json menggunakan tipe data slice