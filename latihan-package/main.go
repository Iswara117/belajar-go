package main

import (
	"fmt"
	"latihan-package/calculator"
)

func main(){
	num := calculator.NewNum1(5)
	// fmt.Println(num.multiply())
	fmt.Println("Hasil Penjumlahan:",num.Add())
}