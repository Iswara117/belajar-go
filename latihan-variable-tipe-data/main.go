package main

import "fmt"

func main() {
	const a  = 10000
	const b  = 15000
	const c  = 7000

	const price_total int = a + b + c // jika ingin menggunakan const maka variable yang akan di jumlahkan juga harus const 

	fmt.Println(price_total)
}