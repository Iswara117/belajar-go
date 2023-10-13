package main

import "fmt"

func main(){
	contact := map[string]int{}

	contact["Alice"] = 1234567890
	contact["Bob"] = 9876543210

	fmt.Print("Semua Kontak", " " , contact)
	fmt.Print("\n Nomor telepon Alice:", " " ,contact["Alice"])

	contact["Charlie"] = 5555555555

	fmt.Print("\n\n","Setelah menambah kontak Charlie:", " " , contact)

	delete(contact, "Bob")

	fmt.Print("\n\n","Setelah hapus kontak Bob:", " " , contact)

	fmt.Print("\n Data Kontak")
	fmt.Print("\n")
	for key, value := range contact {
		fmt.Println(key,":",value)
	}
}