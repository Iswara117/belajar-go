package main

import "fmt"


func main() {
	// Latihan
	/**
	Ubah kode dibawah ini menjadi lebih singkat
	*/

	// sekarang
	umurSaya := 20
	umurKakak := 30
	fmt.Printf("Umur saya sekarang adalah : %d\n", umurSaya)
	fmt.Printf("Umur kakak sekarang adalah : %d\n", umurKakak)

	// tahun depan
	umurSaya++
	umurKakak++

	fmt.Printf("Umur saya tahun depan adalah : %d\n", umurSaya)
	fmt.Printf("Umur kakak tahun depan adalah : %d\n", umurKakak)

	// 5 tahun lagi
	umurSaya += 5
	umurKakak += 5

	fmt.Printf("Umur saya 5 tahun lagi adalah : %d\n", umurSaya)
	fmt.Printf("Umur kakak 5 tahun lagi adalah : %d\n", umurKakak)

}