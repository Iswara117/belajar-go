package main

import "fmt"

func Divide(num1 int, num2 int) interface{} {

	defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
    }() 

	if num1 == 0 || num2 == 0 {
		panic("Tidak dapat membagi oleh nol")
	}

	// err := recover() 

	return num1 / num2
	
}


func main()  {
	var number1 int
	var number2 int

	fmt.Print("Masukkan angka pertama :")
	fmt.Scan(&number1)
	fmt.Print("Masukkan angka kedua :")
	fmt.Scan(&number2)

	fmt.Println("Hasil Pembagian : ", Divide(number1 , number2))
}