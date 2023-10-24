package main

import "fmt"

func Sum(num1 int, num2 int) {
	
	defer fmt.Println("Program selesai")
	fmt.Println("Hasil penjumlahan:", num1 + num2)
	return
}

func main(){
	var number1 int
	var number2 int

	fmt.Print("Masukkan angka pertama :")
	fmt.Scan(&number1)
	fmt.Print("Masukkan angka kedua :")
	fmt.Scan(&number2)

	Sum(number1 , number2)

}