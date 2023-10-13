package main

import "fmt"

func main() {
	var age int
	var gender string
	fmt.Print("masukkan umur :")
	fmt.Scan(&age)
	fmt.Print("masukkan gender :")
	fmt.Scan(&gender)

	if gender == "female" || age >= 21 {
		fmt.Print("Pelamar dapat dipekerjakan.")
	} else {
		fmt.Print("Pelamar tidak dapat dipekerjakan.")
	}
}