package main

import "fmt"

type OrangTua struct {
	Nama string
	Umur int
}

type Siswa struct {
	Orangtua OrangTua
	Nama string
	Umur int
	Kelas string
}


func main(){
	var siswa1 , siswa2 Siswa

	siswa1.Nama , siswa2.Nama = "Ali", "David"
	siswa1.Umur , siswa2.Umur = 15,14
	siswa1.Kelas , siswa2.Kelas = "9A", "8B"
	siswa1.Orangtua.Nama , siswa2.Orangtua.Nama = "Ali", "David"
	siswa1.Orangtua.Umur , siswa2.Orangtua.Umur = 15,14

	fmt.Print("Informasi Siswa 1:","\n")
	fmt.Printf("Nama: %s, Umur: %d, Kelas: %s", siswa1.Nama, siswa1.Umur, siswa1.Kelas)
	fmt.Print("\n")
	fmt.Printf("Orang Tua: %s, Umur: %d", siswa1.Orangtua.Nama, siswa1.Orangtua.Umur)
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Println("Informasi Siswa 2:")
	fmt.Printf("Nama: %s, Umur: %d, Kelas: %s", siswa2.Nama, siswa2.Umur, siswa2.Kelas)
	fmt.Print("\n")
	fmt.Printf("Orang Tua: %s, Umur: %d", siswa2.Orangtua.Nama, siswa2.Orangtua.Umur)

	siswa3 := Siswa{
		Nama:   "Fina",
		Umur: 16,
		Kelas: "10C",
		Orangtua: OrangTua{
			Nama: "Eva",
			Umur: 16,
		},

	}

	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Println("Informasi Siswa 3:")
	fmt.Printf("Nama: %s, Umur: %d, Kelas: %s", siswa3.Nama, siswa3.Umur, siswa3.Kelas)
	fmt.Print("\n")
	fmt.Printf("Orang Tua: %s, Umur: %d", siswa3.Orangtua.Nama, siswa3.Orangtua.Umur)
	// fmt.Printf("Nama: %s,",siswa1.Nama, "," , "Umur:", siswa1.Umur, "," , "Kelas:", siswa1.Kelas )
	fmt.Print("\n")

	var daftarSiswa = []Siswa{
		{Nama: "Eva", 
		Umur: 12, 
		Kelas:"6B", 
		Orangtua: OrangTua{
			Nama:"Rudi",
			Umur: 40,
		},
	},

		{Nama: "Faisal", 
		Umur: 11, 
		Kelas:"5A", 
		Orangtua: OrangTua{
			Nama: "Dewi",
			Umur: 38,
		}},
	
		{Nama: "Grace", 
		Umur: 10, 
		Kelas:"4C", 
		Orangtua: OrangTua{
			Nama: "Hendro",
			Umur: 42,
		}},
	}
	fmt.Print("\n")
	fmt.Print("Daftar Siswa:")
	for _, student := range daftarSiswa {
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Printf("Nama Siswa: %s, Umur: %d, Kelas: %s", student.Nama, student.Umur, student.Kelas)
		fmt.Print("\n")
		fmt.Printf("Orang Tua: %s, Umur: %d", student.Nama, student.Umur)
	}

}