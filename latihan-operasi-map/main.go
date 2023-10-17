package main

import "fmt"


func main(){

	fruits := map[string]int{
		"Apple": 10,
		"Banana": 15,
		"Orange": 8,
		"Grape": 20,
	}
	
	
	fmt.Println("Sebelum ditambah/hapus:")
	fmt.Println(fruits)
	fruits["Mango"] = 12
	fruits["Strawberry"] = 7
	fmt.Println("Setelah menambahkan Mango dan Strawberry:")
	fmt.Println(fruits)
	delete(fruits, "Orange")
	fmt.Println("Setelah menghapus buah Orange:")
	fmt.Println(fruits)
	fmt.Println("Jumlah apel yang tersedia adalah:", fruits["Apple"])


	fmt.Println("Daftar buah-buahan beserta jumlahnya:")
	for key, val := range fruits {
		fmt.Println(key, ":", val)
	}

}
